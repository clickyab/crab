// Code generated build with router DO NOT EDIT.

package controllers

import (
	"sync"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

var once = sync.Once{}

// Routes return the route registered with this
func (c *Controller) Routes(r *xmux.Mux, mountPoint string) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{}

		group := r.NewGroup(mountPoint + "/upload")

		/* Route {
			"Route": "/:module",
			"Method": "POST",
			"Function": "Controller.Upload",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.POST("/:module", xhandler.HandlerFuncC(framework.Mix(c.Upload, m0...)))
		// End route with key 0

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
