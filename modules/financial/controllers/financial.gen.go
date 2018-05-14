// Code generated build with router DO NOT EDIT.

package controllers

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

		group := r.NewGroup("/financial")

		/* Route {
			"Route": "/payment/return/:bank/:hash",
			"Method": "POST",
			"Function": "Controller.backFromBank",
			"RoutePkg": "controllers",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{}...)

		group.POST("controllers-Controller-backFromBank", "/payment/return/:bank/:hash", framework.Mix(c.backFromBank, m0...))
		// End route with key 0

		/* Route {
			"Route": "/graph/spend",
			"Method": "GET",
			"Function": "Controller.graphBillinggraphreport",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "get_billing",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_billing", "get_billing")
		m1 = append(m1, authz.AuthorizeGenerator("get_billing", "self"))

		group.GET("controllers-Controller-graphBillinggraphreport", "/graph/spend", framework.Mix(c.graphBillinggraphreport, m1...))
		// End route with key 1

		/* Route {
			"Route": "/billing",
			"Method": "GET",
			"Function": "Controller.listBillingreport",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "get_billing",
			"Scope": "self"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_billing", "get_billing")
		m2 = append(m2, authz.AuthorizeGenerator("get_billing", "self"))

		group.GET("controllers-Controller-listBillingreport", "/billing", framework.Mix(c.listBillingreport, m2...))
		// End route with key 2

		/* Route {
			"Route": "/billing/definition",
			"Method": "GET",
			"Function": "Controller.defBillingreport",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "get_billing",
			"Scope": "self"
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_billing", "get_billing")
		m3 = append(m3, authz.AuthorizeGenerator("get_billing", "self"))

		group.GET("controllers-Controller-defBillingreport", "/billing/definition", framework.Mix(c.defBillingreport, m3...))
		// End route with key 3

		/* Route {
			"Route": "/payment/init",
			"Method": "POST",
			"Function": "Controller.getPaymentDataPost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "initPaymentPayload",
			"Resource": "make_payment",
			"Scope": "self"
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("make_payment", "make_payment")
		m4 = append(m4, authz.AuthorizeGenerator("make_payment", "self"))

		// Make sure payload is the last middleware
		m4 = append(m4, middleware.PayloadUnMarshallerGenerator(initPaymentPayload{}))
		group.POST("controllers-Controller-getPaymentDataPost", "/payment/init", framework.Mix(c.getPaymentDataPost, m4...))
		// End route with key 4

		/* Route {
			"Route": "/add",
			"Method": "POST",
			"Function": "Controller.registerSnapPost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "registerBankSnapPayload",
			"Resource": "create_bank_snap",
			"Scope": "self"
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("create_bank_snap", "create_bank_snap")
		m5 = append(m5, authz.AuthorizeGenerator("create_bank_snap", "self"))

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(registerBankSnapPayload{}))
		group.POST("controllers-Controller-registerSnapPost", "/add", framework.Mix(c.registerSnapPost, m5...))
		// End route with key 5

		/* Route {
			"Route": "/",
			"Method": "GET",
			"Function": "Controller.billingListGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "list_billing",
			"Scope": "self"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_billing", "list_billing")
		m6 = append(m6, authz.AuthorizeGenerator("list_billing", "self"))

		group.GET("controllers-Controller-billingListGet", "/", framework.Mix(c.billingListGet, m6...))
		// End route with key 6

		/* Route {
			"Route": "/gateways",
			"Method": "GET",
			"Function": "Controller.getGatewaysGet",
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
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-getGatewaysGet", "/gateways", framework.Mix(c.getGatewaysGet, m7...))
		// End route with key 7

		/* Route {
			"Route": "/payment/:id",
			"Method": "GET",
			"Function": "Controller.getPaymentTransactionGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "make_payment",
			"Scope": "self"
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("make_payment", "make_payment")
		m8 = append(m8, authz.AuthorizeGenerator("make_payment", "self"))

		group.GET("controllers-Controller-getPaymentTransactionGet", "/payment/:id", framework.Mix(c.getPaymentTransactionGet, m8...))
		// End route with key 8

		/* Route {
			"Route": "/manual-change-cash",
			"Method": "PUT",
			"Function": "Controller.manualChangeCashPut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "changeCashStatus",
			"Resource": "manual_change_cash",
			"Scope": "global"
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("manual_change_cash", "manual_change_cash")
		m9 = append(m9, authz.AuthorizeGenerator("manual_change_cash", "global"))

		// Make sure payload is the last middleware
		m9 = append(m9, middleware.PayloadUnMarshallerGenerator(changeCashStatus{}))
		group.PUT("controllers-Controller-manualChangeCashPut", "/manual-change-cash", framework.Mix(c.manualChangeCashPut, m9...))
		// End route with key 9

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
