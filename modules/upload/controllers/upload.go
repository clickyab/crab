package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"sync"
	"time"

	"clickyab.com/crab/modules/upload/config"
	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
)

var routes = make(map[string]kind)
var lock = sync.RWMutex{}

type kind struct {
	maxSize int64
	mimes   []string
}

// Register add a route and settings for uploads
// name will be the route, maxsize is maximum allowed size for file upload file and the mimes is alloed mime types
func Register(name string, maxSize int64, mimes ...string) {
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
	m := xmux.Param(ctx, "module")
	u := authz.MustGetUser(ctx)
	lock.RLock()
	defer lock.RUnlock()
	s, ok := routes[m]
	if !ok {
		c.NotFoundResponse(w, fmt.Errorf("Not found"))
		return
	}
	r.ParseMultipartForm(s.maxSize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		c.BadResponse(w, fmt.Errorf(`file not found, the key for file should be "file"`))
		return
	}
	defer file.Close()
	ac, mime := validMIME(handler.Header, s.mimes)
	if !ac {
		c.BadResponse(w, fmt.Errorf("The file type is not valid"))
		return
	}

	ext := filepath.Ext(handler.Filename)

	now := time.Now()
	fp := fmt.Sprintf("%s/%s/", m, now.Format("2006/01/02"))
	os.MkdirAll(fp, os.ModeDir)

	fn := func() string {
		for {
			tmp := fmt.Sprintf("%s_%s.%s", u.ID, <-random.ID, ext)
			if _, err := os.Stat(fp + tmp); os.IsNotExist(err) {
				return tmp
			}
		}
	}()

	f, err := os.Create(fmt.Sprintf("%s/%s", config.Path, fn))
	assert.Nil(err)
	defer f.Close()

	size, er := io.Copy(f, file)
	assert.Nil(er)

	e := model.NewModelManager().CreateUpload(&model.Upload{
		Path:    fn,
		MIME:    mime,
		Size:    size,
		UserID:  u.ID,
		Section: m,
	})
	assert.Nil(e)

	c.JSON(w, http.StatusOK, fn)
}

func validMIME(a textproto.MIMEHeader, b []string) (bool, string) {
	ct := make([]string, 0)
	var ok bool
	if ct, ok = a["Content-Type"]; !ok {
		return false, ""
	}
	for _, ak := range ct {
		for _, bv := range b {
			if ak == bv {
				return true, ak
			}
		}
	}
	return false, ""
}