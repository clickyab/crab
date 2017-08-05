// Code generated build with router DO NOT EDIT.

package user

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

		group := r.NewGroup(mountPoint + "/user")

		/* Route {
			"Route": "/session/close",
			"Method": "GET",
			"Function": "Controller.closeSession",
			"RoutePkg": "user",
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

		group.GET("/session/close", xhandler.HandlerFuncC(framework.Mix(c.closeSession, m0...)))
		// End route with key 0

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
