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
	"sync"
	"time"

	"bytes"

	"image"
	"image/gif"
	"image/jpeg"
	"image/png"

	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
)

var (
	routes = make(map[model.Mime]kind)
	lock   = sync.RWMutex{}
	// UPath default upload path
	UPath = config.RegisterString("crab.modules.upload.path", "/statics/uploads", "a path to the location that uploaded file should save")
	// Perm file perm
	Perm = config.RegisterInt("crab.modules.upload.Perm", 0777, "file will save with this permission")
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
// 		url = /:module
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
