package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"strings"

	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"mime"

	"clickyab.com/crab/modules/upload/errors"
	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/random"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
	"github.com/sirupsen/logrus"
)

var (
	routes = make(map[string]kind)
	lock   = sync.RWMutex{}
	// UPath default upload path
	UPath = config.RegisterString("crab.modules.upload.path", "/statics/uploads", "a path to the location that uploaded file should save")
	// Perm default perm
	Perm = config.RegisterInt("crab.modules.upload.Perm", 0777, "file will save with this permission")
	// MaxVideoDuration video max duration
	MaxVideoDuration = config.RegisterInt64("crab.modules.upload.video.duration", 90, "max video duration in seconds")
	videoSaveFormat  = config.RegisterString("crab.modules.upload.video.format", ".cy", "video format saved")
)

type kind struct {
	maxSize      int64
	mimes        []model.Mime
	maxChunksNum int
}

type uploadResponse struct {
	Src string `json:"src"`
}

// upload route for upload banner,native,avatar,video,...
// @Rest {
// 		url = /module/:module
//		protected = true
// 		method = post
// }
func (c *Controller) upload(ctx context.Context, r *http.Request) (*uploadResponse, error) {
	fileType := xmux.Param(ctx, "module")
	lock.RLock()
	defer lock.RUnlock()
	s, ok := routes[fileType]
	if !ok {
		return nil, errors.InvalidFileTypeError
	}
	currentUser := authz.MustGetUser(ctx)
	flowData, err := getChunkFlowData(r)
	if err != nil {
		return nil, errors.ChunkUploadError
	}
	if flowData.totalChunks > s.maxChunksNum {
		return nil, errors.InvalidError("overflow chunk num")
	}

	var tempDir = filepath.Join(UPath.String(), "temp")
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err = os.MkdirAll(tempDir, os.FileMode(Perm.Int()))
		assert.Nil(err)
	}
	chunkPathDir, file, err := chunkUpload(tempDir, flowData, r)
	if err != nil {
		_ = os.RemoveAll(chunkPathDir)
		return nil, errors.ChunkUploadError
	}
	if file == "" {
		return nil, nil
	}

	fileObj, err := os.Open(file)
	defer func() {
		err = fileObj.Close()
		assert.Nil(err)
	}()

	if err != nil {
		_ = os.RemoveAll(chunkPathDir)
		xlog.GetWithError(ctx, err).Debug("error while open file")
		return nil, errors.FileUploadError
	}

	fileInfo, err := fileObj.Stat()
	if err != nil {
		_ = os.RemoveAll(chunkPathDir)
		xlog.GetWithError(ctx, err).Debug("error while stat file")
		return nil, errors.FileOpenError
	}

	extension := strings.ToLower(filepath.Ext(fileObj.Name()))
	logrus.Debug(fileObj.Name())
	logrus.Debug(extension)

	fileMime := mime.TypeByExtension(extension)
	logrus.Debug(fileMime)
	logrus.Debug(s.mimes)
	if !checkForValidMimes(fileMime, s.mimes) {
		_ = os.RemoveAll(chunkPathDir)
		return nil, errors.WrongMimeType
	}
	var attr = &model.FileAttr{}

	now := time.Now()
	basePath, err := makeBaseDir(fileType, now)
	assert.Nil(err)

	var f = &os.File{}
	var fn string

	tArr := strings.Split(fileType, "-")
	assert.True(len(tArr) == 2)

	if tArr[1] == "video" { //validate video stuff like duration
		f, fn = generateFinalFile(currentUser.ID, basePath, videoSaveFormat.String())
		attr, err = videoUploadHandler(ctx, file, chunkPathDir, extension, f, attr, fileInfo, s.maxSize)
		if err != nil {
			_ = os.RemoveAll(chunkPathDir)
			_ = os.RemoveAll(f.Name())
			return nil, err
		}

	} else { //image selected
		f, fn = generateFinalFile(currentUser.ID, basePath, extension)
		attr, err = imageUploadHandler(fileMime, tArr[0], chunkPathDir, file, fileObj, f, fileInfo, s.maxSize)
		if err != nil {
			_ = os.RemoveAll(chunkPathDir)
			_ = os.RemoveAll(f.Name())
			return nil, err
		}
	}

	dbSavePath := filepath.Join(fileType, now.Format("2006/01/02"), fn)
	g := &model.Upload{
		ID:      dbSavePath,
		MIME:    fileMime,
		Size:    fileInfo.Size(),
		UserID:  currentUser.ID,
		Section: fileType,
		Attr:    *attr,
	}

	e := model.NewModelManager().CreateUpload(g)
	assert.Nil(e)
	return &uploadResponse{Src: g.ID}, nil
}

func getVideoDimensions(streams []map[string]interface{}) (int, int, error) {
	for i := range streams {
		if streams[i]["codec_type"] == "video" {
			w := streams[i]["width"].(int)
			h := streams[i]["height"].(int)
			return w, h, nil
		}
	}
	return 0, 0, errors.FileDimensionError
}

func videoUploadHandler(ctx context.Context, file, chunkPathDir, rawExtension string, f *os.File, attr *model.FileAttr, fileInfo os.FileInfo, maxSize int64) (*model.FileAttr, error) {
	info, err := getVideoInfo(file)
	if err != nil {
		return attr, errors.FileNotReadableError
	}
	if _, ok := info["format"]; !ok {
		return attr, errors.FileFormatNotReadableError
	}
	duration, err := getDuration(info["format"].(map[string]interface{}))
	if err != nil {
		return attr, errors.FileDurationError
	}

	width, height, err := getVideoDimensions(info["streams"].([]map[string]interface{}))
	if err != nil {
		return attr, err
	}

	err = validateVideo(info["format"].(map[string]interface{}), fileInfo, maxSize, duration)
	if err != nil {
		return attr, err
	}

	convertedPath := strings.TrimRight(file, rawExtension) + videoSaveFormat.String()

	//converting the video
	safe.GoRoutine(ctx, func() {
		err := doConvert(file, convertedPath, chunkPathDir, f.Name())
		_ = os.RemoveAll(chunkPathDir)
		assert.Nil(err)
	})
	attr.Video = &model.VideoAttr{
		Duration: int(duration),
		Width:    width,
		Height:   height,
	}
	return attr, nil
}

func generateFinalFile(userID int64, basePath, extension string) (*os.File, string) {
	fn := generateFileName(userID, basePath, extension)
	finalFileTarget := filepath.Join(basePath, fn)
	f, err := os.OpenFile(finalFileTarget, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.FileMode(Perm.Int64()))
	assert.Nil(err)
	defer func() { assert.Nil(f.Close()) }()
	return f, fn
}

func imageUploadHandler(fileMime, fileType, chunkPathDir, file string, fileObj, f *os.File, fileInfo os.FileInfo, maxSize int64) (*model.FileAttr, error) {
	var attr = &model.FileAttr{}
	var err error
	if fileMime == string(model.JPGMime) || fileMime == string(model.PNGMime) || fileMime == string(model.GifMime) || fileMime == string(model.PJPGMime) {
		attr, err = getDimension(model.Mime(fileMime), fileObj, fileType)
		if err != nil {
			return attr, err
		}
	}
	err = validateImage(fileInfo, maxSize)
	if err != nil {
		return attr, errors.FileDimensionError
	}
	assert.Nil(os.Rename(file, f.Name()))
	_ = os.RemoveAll(chunkPathDir)
	return attr, nil
}

func validateImage(fileInfo os.FileInfo, maxSize int64) error {
	//validate size
	size := fileInfo.Size()
	if size > maxSize {
		return errors.LargeFileUploadError
	}
	return nil
}

func checkForValidMimes(mime string, mimes []model.Mime) bool {
	for i := range mimes {
		if string(mimes[i]) == mime {
			return true
		}
	}
	return false
}

func makeBaseDir(fileType string, now time.Time) (string, error) {
	basePath := filepath.Join(UPath.String(), fileType, now.Format("2006/01/02"))
	err := os.MkdirAll(basePath, os.FileMode(Perm.Int64()))
	return basePath, err
}

// Register add a route and settings for uploads
// name will be the route, maxsize is maximum allowed size for file upload file and the mimes is alloed mime types
func Register(name string, maxSize int64, maxChunk int, mimes ...model.Mime) {
	assert.True(len(mimes) > 0)
	lock.Lock()
	defer lock.Unlock()

	_, ok := routes[name]
	assert.False(ok)
	routes[name] = kind{
		maxSize:      maxSize,
		mimes:        mimes,
		maxChunksNum: maxChunk,
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

func validateVideo(format map[string]interface{}, fileInfo os.FileInfo, maxSize, duration int64) error {
	// check format
	if _, ok := format["format_name"]; !ok {
		return errors.InvalidFileTypeError
	}
	formatName := format["format_name"].(string)

	formatsArr := strings.Split(formatName, ",")
	if !array.StringInArray("mp4", formatsArr...) {
		return errors.InvalidFileTypeError
	}

	//validate size
	size := fileInfo.Size()
	if size > maxSize {
		return errors.LargeFileUploadError
	}

	if int64(duration) > MaxVideoDuration.Int64() {
		return errors.FileDurationLimitError
	}
	return nil
}

func getDuration(format map[string]interface{}) (int64, error) {
	if _, ok := format["duration"]; !ok {
		return 0, errors.FileDurationError
	}
	duration, err := strconv.ParseFloat(format["duration"].(string), 64)
	if err != nil {
		return 0, errors.FileDurationError
	}
	return int64(duration), nil
}

func doConvert(file, convertedPath, chunkPathDir, f string) error {
	err := convertVideo(file, convertedPath)
	if err != nil {
		_ = os.RemoveAll(chunkPathDir)
		return errors.FileConvertError
	}
	return os.Rename(convertedPath, f)
}

func getDimension(mime model.Mime, dimensionHandler io.Reader, bannerType string) (*model.FileAttr, error) {
	a := model.FileAttr{}
	var imgConf image.Config
	var err error
	switch mime {
	case model.JPGMime:
		imgConf, err = jpeg.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.FileDimensionError
		}

	case model.PJPGMime:
		imgConf, err = jpeg.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.FileDimensionError
		}
	case model.GifMime:
		imgConf, err = gif.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.FileDimensionError
		}
	case model.PNGMime:
		imgConf, err = png.DecodeConfig(dimensionHandler)
		if err != nil {
			return nil, errors.FileDimensionError
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
	case "user":
		a = model.FileAttr{
			Avatar: &model.AvatarAttr{
				Width:  imgConf.Width,
				Height: imgConf.Height,
			},
		}
	case "domain":
		a = model.FileAttr{
			Avatar: &model.AvatarAttr{
				Width:  imgConf.Width,
				Height: imgConf.Height,
			},
		}
	}

	return &a, nil
}

func generateFileName(userID int64, basePath, ext string) string {
	for {
		tmp := fmt.Sprintf("%d_%s%s", userID, <-random.ID, ext)
		if _, err := os.Stat(basePath + tmp); os.IsNotExist(err) {
			return tmp
		}
	}
}
