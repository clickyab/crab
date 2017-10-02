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

		group := r.NewGroup(mountPoint + "/inventory")

		/* Route {
			"Route": "/presets",
			"Method": "GET",
			"Function": "Controller.whiteBlackLists",
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
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/presets", xhandler.HandlerFuncC(framework.Mix(u.whiteBlackLists, m0...)))
		// End route with key 0

		/* Route {
			"Route": "/preset/:id",
			"Method": "GET",
			"Function": "Controller.whiteBlackList",
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
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("/preset/:id", xhandler.HandlerFuncC(framework.Mix(u.whiteBlackList, m1...)))
		// End route with key 1

		/* Route {
			"Route": "/preset",
			"Method": "POST",
			"Function": "Controller.addPreset",
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
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(whiteBlackList{}))
		group.POST("/preset", xhandler.HandlerFuncC(framework.Mix(u.addPreset, m2...)))
		// End route with key 2

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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("inventory_list", "inventory_list")
		m3 = append(m3, authz.AuthorizeGenerator("inventory_list", "self"))

		group.GET("/list", xhandler.HandlerFuncC(framework.Mix(u.listInventory, m3...)))
		// End route with key 3

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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("inventory_list", "inventory_list")
		m4 = append(m4, authz.AuthorizeGenerator("inventory_list", "self"))

		group.GET("/list/definition", xhandler.HandlerFuncC(framework.Mix(u.defInventory, m4...)))
		// End route with key 4

		initializer.DoInitialize(u)
	})
}

func init() {
	router.Register(&Controller{})
}
