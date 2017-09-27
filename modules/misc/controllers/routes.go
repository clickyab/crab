package misc

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"
	"regexp"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

// Controller is the controller for the misc package
// @Route {
//		group = /misc
// }
type Controller struct {
	controller.Base
}

var (
	swaggerRoute = config.RegisterBoolean("services.framework.swagger", true, "is any swagger code available?")
	panicToken   = config.RegisterString("services.framwork.panic", "caba73d2-189d-4011-9676-969774c2b158", "token for force panic")
)

func load(data map[string]interface{}) map[string]interface{} {
	f := regexp.MustCompile("^file://([a-f0-9]{40})$")
	for i := range data {
		if n, ok := data[i].(map[string]interface{}); ok {
			data[i] = load(n)
			continue
		}
		if _, ok := data[i].(string); !ok {
			continue
		}
		s := data[i].(string)
		m := f.FindStringSubmatch(s)
		if len(m) == 2 {
			d, err := Asset("swagger/" + m[1])
			assert.Nil(err)
			p := make(map[string]interface{})
			assert.Nil(json.Unmarshal(d, &p))
			data[i] = p
		}
	}

	return data
}

// Routes return the route registered with this
func (u *Controller) Routes(r *xmux.Mux, mountPoint string) {
	// This is a special route.
	if !swaggerRoute.Bool() {
		return
	}

	b, err := Asset("swagger/index.json")
	assert.Nil(err)

	var data = make(map[string]interface{})
	err = json.Unmarshal(b, &data)
	assert.Nil(err)
	data = load(data)

	r.GET("/misc/panic", xhandler.HandlerFuncC(
		func(_ context.Context, w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("token") == panicToken.String() {
				assert.Nil(0)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Noop ;)"))
		}))

	r.GET(filepath.Join(mountPoint, "/misc/swagger/index.json"),
		xhandler.HandlerFuncC(func(_ context.Context, w http.ResponseWriter, r *http.Request) {
			tmp := data
			tmp["host"] = r.Host
			framework.JSON(w, http.StatusOK, data)
		}))
}

func init() {
	router.Register(&Controller{})
}
