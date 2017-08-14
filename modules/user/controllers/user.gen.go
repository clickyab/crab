// Code generated build with router DO NOT EDIT.

package user

import (
	"sync"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
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
			"Route": "/corporation",
			"Method": "PUT",
			"Function": "Controller.editCorporation",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "corporation",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(corporation{}))
		group.PUT("/corporation", xhandler.HandlerFuncC(framework.Mix(u.editCorporation, m1...)))
		// End route with key 1

		/* Route {
			"Route": "/personal",
			"Method": "PUT",
			"Function": "Controller.EditPersonal",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "personalPayload",
			"Resource": "",
			"Scope": ""
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(personalPayload{}))
		group.PUT("/personal", xhandler.HandlerFuncC(framework.Mix(u.EditPersonal, m2...)))
		// End route with key 2

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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m3 = append(m3, middleware.PayloadUnMarshallerGenerator(loginPayload{}))
		group.POST("/login", xhandler.HandlerFuncC(framework.Mix(u.login, m3...)))
		// End route with key 3

		/* Route {
			"Route": "/logout",
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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/logout", xhandler.HandlerFuncC(framework.Mix(u.closeSession, m4...)))
		// End route with key 4

		/* Route {
			"Route": "/logout/closeother",
			"Method": "GET",
			"Function": "Controller.closeAllOtherSession",
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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/logout/closeother", xhandler.HandlerFuncC(framework.Mix(u.closeAllOtherSession, m5...)))
		// End route with key 5

		/* Route {
			"Route": "/password/forget",
			"Method": "POST",
			"Function": "Controller.forgetPassword",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "forget",
			"Resource": "",
			"Scope": ""
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m6 = append(m6, middleware.PayloadUnMarshallerGenerator(forget{}))
		group.POST("/password/forget", xhandler.HandlerFuncC(framework.Mix(u.forgetPassword, m6...)))
		// End route with key 6

		/* Route {
			"Route": "/password/callback",
			"Method": "PUT",
			"Function": "Controller.forgetCallBack",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "callBack",
			"Resource": "",
			"Scope": ""
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(callBack{}))
		group.PUT("/password/callback", xhandler.HandlerFuncC(framework.Mix(u.forgetCallBack, m7...)))
		// End route with key 7

		/* Route {
			"Route": "/password/change",
			"Method": "PUT",
			"Function": "Controller.changePassword",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "changePassword",
			"Resource": "",
			"Scope": ""
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m8 = append(m8, middleware.PayloadUnMarshallerGenerator(changePassword{}))
		group.PUT("/password/change", xhandler.HandlerFuncC(framework.Mix(u.changePassword, m8...)))
		// End route with key 8

		/* Route {
			"Route": "/ping",
			"Method": "GET",
			"Function": "Controller.ping",
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
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/ping", xhandler.HandlerFuncC(framework.Mix(u.ping, m9...)))
		// End route with key 9

		/* Route {
			"Route": "/register",
			"Method": "POST",
			"Function": "Controller.Register",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "registerPayload",
			"Resource": "",
			"Scope": ""
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m10 = append(m10, middleware.PayloadUnMarshallerGenerator(registerPayload{}))
		group.POST("/register", xhandler.HandlerFuncC(framework.Mix(u.Register, m10...)))
		// End route with key 10

		initializer.DoInitialize(u)
	})
}

func init() {
	router.Register(&Controller{})
}
