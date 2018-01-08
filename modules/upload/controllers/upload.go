package controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"bytes"

	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"strings"

	"mime"

	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/safe"
	"github.com/rs/xmux"
)

var (
	routes = make(map[model.Mime]kind)
	lock   = sync.RWMutex{}
	// UPath default upload path
	UPath = config.RegisterString("crab.modules.upload.path", "/statics/uploads", "a path to the location that uploaded file should save")
	// ValidVideoMime valid mime video
	ValidVideoMime = config.RegisterString("crab.modules.upload.video.mime", "video/mp4", "comma separated valid video mime")
	// Perm default perm
	Perm = config.RegisterInt("crab.modules.upload.Perm", 0777, "file will save with this permission")
	// VideoMaxSize video max size
	VideoMaxSize = config.RegisterInt64("crab.modules.upload.video.size", 3145728, "max size of video upload")
	// MaxVideoDuration video max duration
	MaxVideoDuration   = config.RegisterInt64("crab.modules.upload.video.duration", 61, "max video duration in seconds")
	videoSaveFormat    = config.RegisterString("crab.modules.upload.video.format", ".cy", "video format saved")
	videoMaxChunkCount = config.RegisterInt("crab.modules.upload.video.max.chunk", 15, "video max chunk count")
)

type kind struct {
	maxSize int64
	mimes   []model.Mime
}

// Register add a route and settings for uploads
// name will be the route, maxsize is maximum allowed size for file upload file and the mimes is alloed mime types
func Register(name model.Mime, maxSize int64, mimes ...model.Mime) {
	assert.True(len(mimes) > 0)
	lock.Lock()
	defer lock.Unlock()

	_, ok := routes[name]
	assert.False(ok)
	routes[name] = kind{
		maxSize: maxSize,
		mimes:   mimes,
	}
}

// Controller is the controller for the user package
// @Route {
//		group = /upload
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

// Upload into the system
// @Route {
// 		url = /banner/:module
//		method = post
//		middleware = authz.Authenticate
// }
func (c Controller) Upload(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	fileType := xmux.Param(ctx, "module")
	m := model.Mime(fileType)
	u := authz.MustGetUser(ctx)
	lock.RLock()
	defer lock.RUnlock()
	s, ok := routes[m]
	if !ok {
		c.NotFoundResponse(w, errors.New("not found"))
		return
	}
	err := r.ParseMultipartForm(s.maxSize)
	if err != nil {
		c.BadResponse(w, fmt.Errorf("max upload size is : %d", s.maxSize))
		return
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		c.BadResponse(w, fmt.Errorf(`file not found, the key for file should be "file"`))
		return
	}
	defer file.Close()

	buff := &bytes.Buffer{}
	dimensionHandler := &bytes.Buffer{}
	multiHandler := io.MultiWriter(dimensionHandler, buff)
	io.Copy(multiHandler, file)
	ac, mime := validMIME(handler.Header, s.mimes)
	if !ac {
		c.BadResponse(w, fmt.Errorf("the file type is not valid"))
		return
	}
	var attr *model.FileAttr
	if mime == model.JPGMime || mime == model.PNGMime || mime == model.GifMime || mime == model.PJPGMime {
		attr, err = getDimension(mime, dimensionHandler, fileType)
		if err != nil {
			c.BadResponse(w, fmt.Errorf("cant get file dimensions"))
			return
		}
	}
	ext := filepath.Ext(handler.Filename)
	now := time.Now()
	fp := filepath.Join(UPath.String(), string(m), now.Format("2006/01/02"))
	err = os.MkdirAll(fp, os.FileMode(Perm.Int64()))
	assert.Nil(err)
	fn := func() string {
		for {
			tmp := fmt.Sprintf("%d_%s%s", u.ID, <-random.ID, ext)
			if _, err := os.Stat(fp + tmp); os.IsNotExist(err) {
				return tmp
			}
		}
	}()
	f, err := os.OpenFile(filepath.Join(fp, fn), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(Perm.Int64()))
	assert.Nil(err)
	defer f.Close()

	finalPath := filepath.Join(string(m), now.Format("2006/01/02"), fn)
	size, er := io.Copy(f, buff)
	assert.Nil(er)

	g := &model.Upload{
		ID:      finalPath,
		MIME:    string(mime),
		Size:    size,
		UserID:  u.ID,
		Section: string(m),
		Attr:    *attr,
	}
	e := model.NewModelManager().CreateUpload(g)
	assert.Nil(e)
	c.JSON(w, http.StatusOK, struct {
		Src string `json:"src"`
	}{
		Src: g.ID,
	})
}

func validMIME(a textproto.MIMEHeader, b []model.Mime) (bool, model.Mime) {
	var ct []string
	var ok bool
	if ct, ok = a["Content-Type"]; !ok {
		return false, ""
	}
	for _, ak := range ct {
		for _, bv := range b {
			if ak == string(bv) {
				return true, model.Mime(ak)
			}
		}
	}
	return false, ""
}

// Upload video into the system
// @Route {
// 		url = /video
//		method = post
//		middleware = authz.Authenticate
// }
func (c Controller) videoUpload(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	currentUser := authz.MustGetUser(ctx)
	flowData, err := getChunkFlowData(r)
	if err != nil {
		c.BadResponse(w, nil)
		return
	}
	if flowData.totalChunks > videoMaxChunkCount.Int() {
		c.BadResponse(w, errors.New("file size not valid"))
		return
	}
	var tempDir = filepath.Join(os.TempDir(), "uploads")
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err = os.MkdirAll(tempDir, os.FileMode(Perm.Int()))
		assert.Nil(err)
	}
	chunkPathDir, file, err := chunkUpload(tempDir, flowData, r)
	if err != nil {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("failed to upload chunks"))
		return
	}
	if file == "" {
		c.OKResponse(w, errors.New("failed storing chunks"))
		return
	}
	// open uploaded file in tmp folder
	fileObj, err := os.Open(file)
	defer func() {
		err = fileObj.Close()
		assert.Nil(err)
	}()

	if err != nil {
		c.BadResponse(w, nil)
		return
	}
	fileInfo, err := fileObj.Stat()
	if err != nil {
		c.BadResponse(w, errors.New("cant open uploaded file"))
		return
	}
	extension := strings.ToLower(filepath.Ext(fileObj.Name()))
	//check if file extension is valid
	mimeType := mime.TypeByExtension(extension)
	validMimeArr := strings.Split(ValidVideoMime.String(), ",")
	isValidMime := func() bool {
		for i := range validMimeArr {
			if validMimeArr[i] == mimeType {
				return true
			}
		}
		return false
	}()
	if !isValidMime {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("video mime type not valid"))
		return
	}
	size := fileInfo.Size()
	//check size
	if size > VideoMaxSize.Int64() {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("video is too large to be uploaded"))
		return
	}

	info, err := getVideoInfo(file)
	if err != nil {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("uploaded file is not readable"))
		return
	}

	if _, ok := info["format"]; !ok {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("file format is not readable"))
		return
	}
	format := info["format"].(map[string]interface{})
	if _, ok := format["duration"]; !ok {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("cant get duration from file"))
		return
	}
	duration, err := strconv.ParseFloat(format["duration"].(string), 64)
	if err != nil {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, errors.New("error parsing duration from video"))
		return
	}
	if int64(duration) > MaxVideoDuration.Int64() {
		os.RemoveAll(chunkPathDir)
		c.BadResponse(w, fmt.Errorf("maximum duration is %d seconds", MaxVideoDuration.Int64()))
		return
	}
	convertedPath := strings.TrimRight(file, extension) + videoSaveFormat.String()
	now := time.Now()
	basePath := filepath.Join(UPath.String(), "video", now.Format("2006/01/02"))
	err = os.MkdirAll(basePath, os.FileMode(Perm.Int64()))
	assert.Nil(err)
	fn := generateFileName(currentUser.ID, basePath, videoSaveFormat.String())
	f, err := os.OpenFile(filepath.Join(basePath, fn), os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(Perm.Int64()))
	assert.Nil(err)
	defer f.Close()

	//converting the video
	safe.GoRoutine(ctx, func() {
		doConvert(file, convertedPath, chunkPathDir, f.Name())
	})

	dbSavePath := filepath.Join("video", now.Format("2006/01/02"), fn)
	g := &model.Upload{
		ID:      dbSavePath,
		MIME:    mimeType,
		Size:    size,
		UserID:  currentUser.ID,
		Section: "video",
		Attr: model.FileAttr{
			Video: &model.VideoAttr{
				Duration: int(duration),
			},
		},
	}
	e := model.NewModelManager().CreateUpload(g)
	assert.Nil(e)
	c.JSON(w, http.StatusOK, struct {
		Src string `json:"src"`
	}{
		Src: g.ID,
	})

}

func doConvert(file, convertedPath, chunkPathDir, f string) error {
	err := convertVideo(file, convertedPath)
	if err != nil {
		os.RemoveAll(chunkPathDir)
		return errors.New("cant convert video")
	}
	err = os.Rename(convertedPath, f)
	assert.Nil(err)
	//remove chunk dir
	os.RemoveAll(chunkPathDir)
	return nil
}

func getDimension(mime model.Mime, dimensionHandler *bytes.Buffer, bannerType string) (*model.FileAttr, error) {
	a := model.FileAttr{}
	var imgConf image.Config
	var err error
	switch mime {
	case model.JPGMime:
		imgConf, err = jpeg.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.New("cant get file dimensions")
		}

	case model.PJPGMime:
		imgConf, err = jpeg.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.New("cant get file dimensions")
		}
	case model.GifMime:
		imgConf, err = gif.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.New("cant get file dimensions")
		}
	case model.PNGMime:
		imgConf, err = png.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.New("cant get file dimensions")
		}
	}

	switch bannerType {
	case "banner":
		a = model.FileAttr{
			Banner: &model.BannerAttr{
				Width:  imgConf.Width,
				Height: imgConf.Height,
			},
		}
	case "native":
		a = model.FileAttr{
			Native: &model.NativeAttr{
				Width:  imgConf.Width,
				Height: imgConf.Height,
			},
		}
	case "avatar":
		a = model.FileAttr{
			Avatar: &model.AvatarAttr{
				Width:  imgConf.Width,
				Height: imgConf.Height,
			},
		}
	}

	return &a, nil
}

func generateFileName(userID int64, basePath, fileName string) string {
	for {
		tmp := fmt.Sprintf("%d_%s%s", userID, <-random.ID, fileName)
		if _, err := os.Stat(basePath + tmp); os.IsNotExist(err) {
			return tmp
		}
	}
}
