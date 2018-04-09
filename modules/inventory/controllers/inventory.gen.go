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
			"Route": "/inventory/list",
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
			"Resource": "list_inventory",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_inventory", "list_inventory")
		m0 = append(m0, authz.AuthorizeGenerator("list_inventory", "self"))

		group.GET("controllers-Controller-listInventory", "/inventory/list", framework.Mix(ctrl.listInventory, m0...))
		// End route with key 0

		/* Route {
			"Route": "/inventory/list/definition",
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
			"Resource": "list_inventory",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_inventory", "list_inventory")
		m1 = append(m1, authz.AuthorizeGenerator("list_inventory", "self"))

		group.GET("controllers-Controller-defInventory", "/inventory/list/definition", framework.Mix(ctrl.defInventory, m1...))
		// End route with key 1

		/* Route {
			"Route": "/publisher/list",
			"Method": "GET",
			"Function": "Controller.listPublisher",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "publisher_list",
			"Scope": "self"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("publisher_list", "publisher_list")
		m2 = append(m2, authz.AuthorizeGenerator("publisher_list", "self"))

		group.GET("controllers-Controller-listPublisher", "/publisher/list", framework.Mix(ctrl.listPublisher, m2...))
		// End route with key 2

		/* Route {
			"Route": "/publisher/list/definition",
			"Method": "GET",
			"Function": "Controller.defPublisher",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "publisher_list",
			"Scope": "self"
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("publisher_list", "publisher_list")
		m3 = append(m3, authz.AuthorizeGenerator("publisher_list", "self"))

		group.GET("controllers-Controller-defPublisher", "/publisher/list/definition", framework.Mix(ctrl.defPublisher, m3...))
		// End route with key 3

		/* Route {
			"Route": "/addpub/:id",
			"Method": "PATCH",
			"Function": "Controller.addPresetPatch",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "addInventoryPayload",
			"Resource": "edit_inventory",
			"Scope": "self"
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m4 = append(m4, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m4 = append(m4, middleware.PayloadUnMarshallerGenerator(addInventoryPayload{}))
		group.PATCH("controllers-Controller-addPresetPatch", "/addpub/:id", framework.Mix(ctrl.addPresetPatch, m4...))
		// End route with key 4

		/* Route {
			"Route": "/removepub/:id",
			"Method": "PATCH",
			"Function": "Controller.removePresetPatch",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "removeInventoryPayload",
			"Resource": "edit_inventory",
			"Scope": "self"
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m5 = append(m5, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(removeInventoryPayload{}))
		group.PATCH("controllers-Controller-removePresetPatch", "/removepub/:id", framework.Mix(ctrl.removePresetPatch, m5...))
		// End route with key 5

		/* Route {
			"Route": "/:id",
			"Method": "PUT",
			"Function": "Controller.changeLabelPut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "changeLabelPayload",
			"Resource": "edit_inventory",
			"Scope": "self"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m6 = append(m6, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m6 = append(m6, middleware.PayloadUnMarshallerGenerator(changeLabelPayload{}))
		group.PUT("controllers-Controller-changeLabelPut", "/:id", framework.Mix(ctrl.changeLabelPut, m6...))
		// End route with key 6

		/* Route {
			"Route": "/create",
			"Method": "POST",
			"Function": "Controller.createPresetPost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "createInventoryPayload",
			"Resource": "add_inventory",
			"Scope": "self"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("add_inventory", "add_inventory")
		m7 = append(m7, authz.AuthorizeGenerator("add_inventory", "self"))

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(createInventoryPayload{}))
		group.POST("controllers-Controller-createPresetPost", "/create", framework.Mix(ctrl.createPresetPost, m7...))
		// End route with key 7

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
