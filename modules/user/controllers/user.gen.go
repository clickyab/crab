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
	"github.com/clickyab/services/permission"
)

var once = sync.Once{}

// Routes return the route registered with this
func (c *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup("/user")

		/* Route {
			"Route": "/list",
			"Method": "GET",
			"Function": "Controller.listUsers_list",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "user_list",
			"Scope": "global"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("user_list", "user_list")
		m0 = append(m0, authz.AuthorizeGenerator("user_list", "global"))

		group.GET("user-Controller-listUsers_list", "/list", framework.Mix(c.listUsers_list, m0...))
		// End route with key 0

		/* Route {
			"Route": "/list/definition",
			"Method": "GET",
			"Function": "Controller.defUsers_list",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "user_list",
			"Scope": "global"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("user_list", "user_list")
		m1 = append(m1, authz.AuthorizeGenerator("user_list", "global"))

		group.GET("user-Controller-defUsers_list", "/list/definition", framework.Mix(c.defUsers_list, m1...))
		// End route with key 1

		/* Route {
			"Route": "/add-to/whitelabel",
			"Method": "POST",
			"Function": "Controller.registerToWhitelabelPost",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "addUserToWhitelabelPayload",
			"Resource": "add_to_whitelabel_user",
			"Scope": "global"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("add_to_whitelabel_user", "add_to_whitelabel_user")
		m2 = append(m2, authz.AuthorizeGenerator("add_to_whitelabel_user", "global"))

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(addUserToWhitelabelPayload{}))
		group.POST("user-Controller-registerToWhitelabelPost", "/add-to/whitelabel", framework.Mix(c.registerToWhitelabelPost, m2...))
		// End route with key 2

		/* Route {
			"Route": "/avatar",
			"Method": "PUT",
			"Function": "Controller.avatarPut",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "avatarPayload",
			"Resource": "",
			"Scope": ""
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m3 = append(m3, middleware.PayloadUnMarshallerGenerator(avatarPayload{}))
		group.PUT("user-Controller-avatarPut", "/avatar", framework.Mix(c.avatarPut, m3...))
		// End route with key 3

		/* Route {
			"Route": "/password/change/:token",
			"Method": "PUT",
			"Function": "Controller.changeForgetPasswordPut",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "callBackPayload",
			"Resource": "",
			"Scope": ""
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m4 = append(m4, middleware.PayloadUnMarshallerGenerator(callBackPayload{}))
		group.PUT("user-Controller-changeForgetPasswordPut", "/password/change/:token", framework.Mix(c.changeForgetPasswordPut, m4...))
		// End route with key 4

		/* Route {
			"Route": "/password/change",
			"Method": "PUT",
			"Function": "Controller.changePasswordPut",
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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(changePassword{}))
		group.PUT("user-Controller-changePasswordPut", "/password/change", framework.Mix(c.changePasswordPut, m5...))
		// End route with key 5

		/* Route {
			"Route": "/mail/check",
			"Method": "POST",
			"Function": "Controller.checkMailPost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "checkMailPayload",
			"Resource": "",
			"Scope": ""
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m6 = append(m6, middleware.PayloadUnMarshallerGenerator(checkMailPayload{}))
		group.POST("user-Controller-checkMailPost", "/mail/check", framework.Mix(c.checkMailPost, m6...))
		// End route with key 6

		/* Route {
			"Route": "/update",
			"Method": "PUT",
			"Function": "Controller.editPut",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "userPayload",
			"Resource": "",
			"Scope": ""
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(userPayload{}))
		group.PUT("user-Controller-editPut", "/update", framework.Mix(c.editPut, m7...))
		// End route with key 7

		/* Route {
			"Route": "/password/verify/",
			"Method": "POST",
			"Function": "Controller.checkForgetCodePost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "forgetCodePayload",
			"Resource": "",
			"Scope": ""
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m8 = append(m8, middleware.PayloadUnMarshallerGenerator(forgetCodePayload{}))
		group.POST("user-Controller-checkForgetCodePost", "/password/verify/", framework.Mix(c.checkForgetCodePost, m8...))
		// End route with key 8

		/* Route {
			"Route": "/password/verify/:token",
			"Method": "GET",
			"Function": "Controller.checkForgetHashGet",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("user-Controller-checkForgetHashGet", "/password/verify/:token", framework.Mix(c.checkForgetHashGet, m9...))
		// End route with key 9

		/* Route {
			"Route": "/password/forget",
			"Method": "POST",
			"Function": "Controller.forgetPasswordPost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "forgetPayload",
			"Resource": "",
			"Scope": ""
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m10 = append(m10, middleware.PayloadUnMarshallerGenerator(forgetPayload{}))
		group.POST("user-Controller-forgetPasswordPost", "/password/forget", framework.Mix(c.forgetPasswordPost, m10...))
		// End route with key 10

		/* Route {
			"Route": "/login",
			"Method": "POST",
			"Function": "Controller.loginPost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "loginPayload",
			"Resource": "",
			"Scope": ""
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m11 = append(m11, middleware.PayloadUnMarshallerGenerator(loginPayload{}))
		group.POST("user-Controller-loginPost", "/login", framework.Mix(c.loginPost, m11...))
		// End route with key 11

		/* Route {
			"Route": "/logout",
			"Method": "GET",
			"Function": "Controller.closeSessionGet",
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
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("user-Controller-closeSessionGet", "/logout", framework.Mix(c.closeSessionGet, m12...))
		// End route with key 12

		/* Route {
			"Route": "/logout/closeother",
			"Method": "GET",
			"Function": "Controller.closeAllOtherSessionGet",
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
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("user-Controller-closeAllOtherSessionGet", "/logout/closeother", framework.Mix(c.closeAllOtherSessionGet, m13...))
		// End route with key 13

		/* Route {
			"Route": "/ping",
			"Method": "GET",
			"Function": "Controller.pingGet",
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
		} with key 14 */
		m14 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("user-Controller-pingGet", "/ping", framework.Mix(c.pingGet, m14...))
		// End route with key 14

		/* Route {
			"Route": "/register",
			"Method": "POST",
			"Function": "Controller.registerPost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "registerPayload",
			"Resource": "",
			"Scope": ""
		} with key 15 */
		m15 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m15 = append(m15, middleware.PayloadUnMarshallerGenerator(registerPayload{}))
		group.POST("user-Controller-registerPost", "/register", framework.Mix(c.registerPost, m15...))
		// End route with key 15

		/* Route {
			"Route": "/search/mail",
			"Method": "POST",
			"Function": "Controller.searchByMailPost",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "searchUserPayload",
			"Resource": "",
			"Scope": ""
		} with key 16 */
		m16 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m16 = append(m16, middleware.PayloadUnMarshallerGenerator(searchUserPayload{}))
		group.POST("user-Controller-searchByMailPost", "/search/mail", framework.Mix(c.searchByMailPost, m16...))
		// End route with key 16

		/* Route {
			"Route": "/store",
			"Method": "POST",
			"Function": "Controller.storePost",
			"RoutePkg": "user",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "storePayload",
			"Resource": "",
			"Scope": ""
		} with key 17 */
		m17 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m17 = append(m17, middleware.PayloadUnMarshallerGenerator(storePayload{}))
		group.POST("user-Controller-storePost", "/store", framework.Mix(c.storePost, m17...))
		// End route with key 17

		/* Route {
			"Route": "/email/verify/:token",
			"Method": "GET",
			"Function": "Controller.verifyEmailGet",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 18 */
		m18 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("user-Controller-verifyEmailGet", "/email/verify/:token", framework.Mix(c.verifyEmailGet, m18...))
		// End route with key 18

		/* Route {
			"Route": "/email/verify",
			"Method": "POST",
			"Function": "Controller.verifyEmailCodePost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "verifyEmailCodePayload",
			"Resource": "",
			"Scope": ""
		} with key 19 */
		m19 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m19 = append(m19, middleware.PayloadUnMarshallerGenerator(verifyEmailCodePayload{}))
		group.POST("user-Controller-verifyEmailCodePost", "/email/verify", framework.Mix(c.verifyEmailCodePost, m19...))
		// End route with key 19

		/* Route {
			"Route": "/email/verify/resend",
			"Method": "POST",
			"Function": "Controller.verifyResendPost",
			"RoutePkg": "user",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "verifyResendPayload",
			"Resource": "",
			"Scope": ""
		} with key 20 */
		m20 := append(groupMiddleware, []framework.Middleware{}...)

		// Make sure payload is the last middleware
		m20 = append(m20, middleware.PayloadUnMarshallerGenerator(verifyResendPayload{}))
		group.POST("user-Controller-verifyResendPost", "/email/verify/resend", framework.Mix(c.verifyResendPost, m20...))
		// End route with key 20

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
