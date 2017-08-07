package controllers

import (
	"context"
	"net/http"

	"github.com/clickyab/services/assert"

	"fmt"
	"net/textproto"
	"sync"

	"strings"

	"os"

	"io"

	"time"

	"clickyab.com/crab/modules/upload/config"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
)

var routes = make(map[string]kind)
var lock = sync.Mutex{}

type kind struct {
	maxSize int64
	accept  []string
}

func Register(name string, maxSize int64, accept ...string) {
	lock.Lock()
	defer lock.Unlock()

	_, ok := routes[name]
	assert.False(ok)
	routes[name] = kind{
		maxSize: maxSize,
		accept:  accept,
	}
}

// Controller is the controller for the user package
// @Route {
//		group = /upload
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
	s, ok := routes[m]
	if !ok {
		c.NotFoundResponse(w, fmt.Errorf("Not found"))
	}
	r.ParseMultipartForm(s.maxSize)
	file, handler, err := r.FormFile("file")
	if err != nil {
		c.BadResponse(w, fmt.Errorf(`file not found, the key for file should be "file"`))
	}
	ac, _ := validMIME(handler.Header, s.accept)
	if !ac {
		c.BadResponse(w, fmt.Errorf("The file type is not valid"))
	}
	defer file.Close()
	ext := ""
	if tx := strings.Split(handler.Filename, ","); len(tx) > 1 {
		ext = tx[len(tx)-1]
	}
	now := time.Now()
	fn := ""
	for {
		fn := fmt.Sprintf("%s/%s/%s.%s", m, now.Format("2006/01/02"), <-random.ID, ext)
		if _, err := os.Stat(fn); os.IsNotExist(err) {
			break
		}
	}
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", ucfg.Path, fn), os.O_WRONLY|os.O_CREATE, 0666)
	assert.Nil(err)
	defer f.Close()

	_, er := io.Copy(f, file)
	assert.Nil(er)
	// TODO save into database
	c.JSON(w, http.StatusOK, fn)
}

func validMIME(a textproto.MIMEHeader, b []string) (bool, string) {
	for ak := range a {
		for _, bv := range b {
			if ak == bv {
				return true, ak
			}
		}
	}
	return false, ""
}
