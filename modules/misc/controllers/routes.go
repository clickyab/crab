package controllers

import (
	"context"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/framework/router"
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
	assetHandler := http.FileServer(rice.MustFindBox("../../../../swagger-ui").HTTPBox())
	framework.Any(r, "/swagger/*filename", func(_ context.Context, w http.ResponseWriter, r *http.Request) {
		http.StripPrefix("/swagger/", assetHandler).ServeHTTP(w, r)
	})
}

func init() {
	router.Register(&Controller{})
}
