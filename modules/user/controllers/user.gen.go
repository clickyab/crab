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
func (ctrl *Controller) Routes(r *xmux.Mux, mountPoint string) {
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
		group.POST("/mail/check", xhandler.HandlerFuncC(framework.Mix(ctrl.checkMail, m0...)))
		// End route with key 0

		/* Route {
			"Route": "/personal",
			"Method": "PUT",
			"Function": "Controller.Edit",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "userPayload",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(userPayload{}))
		group.PUT("/personal", xhandler.HandlerFuncC(framework.Mix(ctrl.Edit, m1...)))
		// End route with key 1

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
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(loginPayload{}))
		group.POST("/login", xhandler.HandlerFuncC(framework.Mix(ctrl.login, m2...)))
		// End route with key 2

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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/logout", xhandler.HandlerFuncC(framework.Mix(ctrl.closeSession, m3...)))
		// End route with key 3

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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/logout/closeother", xhandler.HandlerFuncC(framework.Mix(ctrl.closeAllOtherSession, m4...)))
		// End route with key 4

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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(forget{}))
		group.POST("/password/forget", xhandler.HandlerFuncC(framework.Mix(ctrl.forgetPassword, m5...)))
		// End route with key 5

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
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m6 = append(m6, middleware.PayloadUnMarshallerGenerator(callBack{}))
		group.PUT("/password/callback", xhandler.HandlerFuncC(framework.Mix(ctrl.forgetCallBack, m6...)))
		// End route with key 6

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
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(changePassword{}))
		group.PUT("/password/change", xhandler.HandlerFuncC(framework.Mix(ctrl.changePassword, m7...)))
		// End route with key 7

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
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/ping", xhandler.HandlerFuncC(framework.Mix(ctrl.ping, m8...)))
		// End route with key 8

		/* Route {
			"Route": "/register",
			"Method": "POST",
			"Function": "Controller.register",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "registerPayload",
			"Resource": "",
			"Scope": ""
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m9 = append(m9, middleware.PayloadUnMarshallerGenerator(registerPayload{}))
		group.POST("/register", xhandler.HandlerFuncC(framework.Mix(ctrl.register, m9...)))
		// End route with key 9

		/* Route {
			"Route": "/verify/:hash/:key",
			"Method": "GET",
			"Function": "Controller.verifyEmail",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/verify/:hash/:key", xhandler.HandlerFuncC(framework.Mix(ctrl.verifyEmail, m10...)))
		// End route with key 10

		/* Route {
			"Route": "/verify/resend",
			"Method": "POST",
			"Function": "Controller.verifyResend",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{}...)

		group.POST("/verify/resend", xhandler.HandlerFuncC(framework.Mix(ctrl.verifyResend, m11...)))
		// End route with key 11

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
