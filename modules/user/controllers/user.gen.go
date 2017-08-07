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

		group.GET("/dummy", xhandler.HandlerFuncC(framework.Mix(c.dummy, m0...)))
		// End route with key 0

		/* Route {
			"Route": "/login",
			"Method": "POST",
			"Function": "Controller.login",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "loginPayload",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(loginPayload{}))
		group.POST("/login", xhandler.HandlerFuncC(framework.Mix(c.login, m1...)))
		// End route with key 1

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
