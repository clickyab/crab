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
func (ctrl *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup("/inventory")

		/* Route {
			"Route": "/list",
			"Method": "GET",
			"Function": "Controller.listInventory",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "inventory_list",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("inventory_list", "inventory_list")
		m0 = append(m0, authz.AuthorizeGenerator("inventory_list", "self"))

		group.GET("controllers-Controller-listInventory", "/list", framework.Mix(ctrl.listInventory, m0...))
		// End route with key 0

		/* Route {
			"Route": "/list/definition",
			"Method": "GET",
			"Function": "Controller.defInventory",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "inventory_list",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("inventory_list", "inventory_list")
		m1 = append(m1, authz.AuthorizeGenerator("inventory_list", "self"))

		group.GET("controllers-Controller-defInventory", "/list/definition", framework.Mix(ctrl.defInventory, m1...))
		// End route with key 1

		/* Route {
			"Route": "/presets",
			"Method": "GET",
			"Function": "Controller.whiteBlackListsGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-whiteBlackListsGet", "/presets", framework.Mix(ctrl.whiteBlackListsGet, m2...))
		// End route with key 2

		/* Route {
			"Route": "/preset/:id",
			"Method": "GET",
			"Function": "Controller.whiteBlackListGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-whiteBlackListGet", "/preset/:id", framework.Mix(ctrl.whiteBlackListGet, m3...))
		// End route with key 3

		/* Route {
			"Route": "/preset",
			"Method": "POST",
			"Function": "Controller.addPresetPost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "whiteBlackList",
			"Resource": "",
			"Scope": ""
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m4 = append(m4, middleware.PayloadUnMarshallerGenerator(whiteBlackList{}))
		group.POST("controllers-Controller-addPresetPost", "/preset", framework.Mix(ctrl.addPresetPost, m4...))
		// End route with key 4

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
