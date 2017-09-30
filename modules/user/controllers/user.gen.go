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
)

var once = sync.Once{}

// Routes return the route registered with this
func (ctrl *Controller) Routes(r router.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup("/user")

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
		group.POST("/mail/check", framework.Mix(ctrl.checkMail, m0...))
		// End route with key 0

		/* Route {
			"Route": "/update",
			"Method": "PUT",
			"Function": "Controller.edit",
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
		group.PUT("/update", framework.Mix(ctrl.edit, m1...))
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
		group.POST("/login", framework.Mix(ctrl.login, m2...))
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

		group.GET("/logout", framework.Mix(ctrl.closeSession, m3...))
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

		group.GET("/logout/closeother", framework.Mix(ctrl.closeAllOtherSession, m4...))
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
			"Payload": "forgetPayload",
			"Resource": "",
			"Scope": ""
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(forgetPayload{}))
		group.POST("/password/forget", framework.Mix(ctrl.forgetPassword, m5...))
		// End route with key 5

		/* Route {
			"Route": "/password/verify/:token",
			"Method": "GET",
			"Function": "Controller.checkForgetHash",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/password/verify/:token", framework.Mix(ctrl.checkForgetHash, m6...))
		// End route with key 6

		/* Route {
			"Route": "/password/verify/",
			"Method": "POST",
			"Function": "Controller.checkForgetCode",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "forgetCodePayload",
			"Resource": "",
			"Scope": ""
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(forgetCodePayload{}))
		group.POST("/password/verify/", framework.Mix(ctrl.checkForgetCode, m7...))
		// End route with key 7

		/* Route {
			"Route": "/password/change/:token",
			"Method": "PUT",
			"Function": "Controller.changeForgetPassword",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "callBackPayload",
			"Resource": "",
			"Scope": ""
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m8 = append(m8, middleware.PayloadUnMarshallerGenerator(callBackPayload{}))
		group.PUT("/password/change/:token", framework.Mix(ctrl.changeForgetPassword, m8...))
		// End route with key 8

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
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m9 = append(m9, middleware.PayloadUnMarshallerGenerator(changePassword{}))
		group.PUT("/password/change", framework.Mix(ctrl.changePassword, m9...))
		// End route with key 9

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
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/ping", framework.Mix(ctrl.ping, m10...))
		// End route with key 10

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
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m11 = append(m11, middleware.PayloadUnMarshallerGenerator(registerPayload{}))
		group.POST("/register", framework.Mix(ctrl.register, m11...))
		// End route with key 11

		/* Route {
			"Route": "/email/verify/:token",
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
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/email/verify/:token", framework.Mix(ctrl.verifyEmail, m12...))
		// End route with key 12

		/* Route {
			"Route": "/email/verify",
			"Method": "POST",
			"Function": "Controller.verifyEmailCode",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "verifyEmailCodePayload",
			"Resource": "",
			"Scope": ""
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m13 = append(m13, middleware.PayloadUnMarshallerGenerator(verifyEmailCodePayload{}))
		group.POST("/email/verify", framework.Mix(ctrl.verifyEmailCode, m13...))
		// End route with key 13

		/* Route {
			"Route": "/email/verify/resend",
			"Method": "POST",
			"Function": "Controller.verifyResend",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "verifyResendPayload",
			"Resource": "",
			"Scope": ""
		} with key 14 */
		m14 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m14 = append(m14, middleware.PayloadUnMarshallerGenerator(verifyResendPayload{}))
		group.POST("/email/verify/resend", framework.Mix(ctrl.verifyResend, m14...))
		// End route with key 14

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
