// Code generated build with router DO NOT EDIT.

package user

import (
	"sync"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

var once = sync.Once{}

// Routes return the route registered with this
func (u *Controller) Routes(r *xmux.Mux, mountPoint string) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup(mountPoint + "/user")

		/* Route {
			"Route": "/mail/check",
			"Method": "POST",
			"Function": "Controller.checkMail",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "checkMailPayload",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(checkMailPayload{}))
		group.POST("/mail/check", xhandler.HandlerFuncC(framework.Mix(u.checkMail, m0...)))
		// End route with key 0

		/* Route {
			"Route": "/register",
			"Method": "POST",
			"Function": "Controller.Register",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"domain.Access"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "registerPayload",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			domain.Access,
		}...)

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(registerPayload{}))
		group.POST("/register", xhandler.HandlerFuncC(framework.Mix(u.Register, m1...)))
		// End route with key 1

		initializer.DoInitialize(u)
	})
}

func init() {
	router.Register(&Controller{})
}
