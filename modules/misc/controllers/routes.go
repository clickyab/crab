package misc

import (
	"context"
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/framework/router"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

type Controller struct {
	controller.Base
}

var swaggerRoute = config.RegisterBoolean("services.framework.swagger", true, "is any swagger code available?")

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
