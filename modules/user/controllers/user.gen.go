// Code generated build with router DO NOT EDIT.

package user

import (
	"sync"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

var once = sync.Once{}

// Routes return the route registered with this
func (u *Controller) Routes(r *xmux.Mux, mountPoint string) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{}

		group := r.NewGroup(mountPoint + "/user")

		/* Route {
			"Route": "/dummy",
			"Method": "GET",
			"Function": "Controller.dummy",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/dummy", xhandler.HandlerFuncC(framework.Mix(u.dummy, m0...)))
		// End route with key 0

		initializer.DoInitialize(u)
	})
}

func init() {
	router.Register(&Controller{})
}
