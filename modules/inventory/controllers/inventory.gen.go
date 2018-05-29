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
			"Route": "/base-publishers/statistics",
			"Method": "GET",
			"Function": "Controller.listBase_publisher_statistics",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "publisher_base_statistics",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("publisher_base_statistics", "publisher_base_statistics")
		m0 = append(m0, authz.AuthorizeGenerator("publisher_base_statistics", "self"))

		group.GET("controllers-Controller-listBase_publisher_statistics", "/base-publishers/statistics", framework.Mix(ctrl.listBase_publisher_statistics, m0...))
		// End route with key 0

		/* Route {
			"Route": "/base-publishers/statistics/definition",
			"Method": "GET",
			"Function": "Controller.defBase_publisher_statistics",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "publisher_base_statistics",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("publisher_base_statistics", "publisher_base_statistics")
		m1 = append(m1, authz.AuthorizeGenerator("publisher_base_statistics", "self"))

		group.GET("controllers-Controller-defBase_publisher_statistics", "/base-publishers/statistics/definition", framework.Mix(ctrl.defBase_publisher_statistics, m1...))
		// End route with key 1

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
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_inventory", "list_inventory")
		m2 = append(m2, authz.AuthorizeGenerator("list_inventory", "self"))

		group.GET("controllers-Controller-listInventory", "/inventory/list", framework.Mix(ctrl.listInventory, m2...))
		// End route with key 2

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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_inventory", "list_inventory")
		m3 = append(m3, authz.AuthorizeGenerator("list_inventory", "self"))

		group.GET("controllers-Controller-defInventory", "/inventory/list/definition", framework.Mix(ctrl.defInventory, m3...))
		// End route with key 3

		/* Route {
			"Route": "/publisher/list/single/:id",
			"Method": "GET",
			"Function": "Controller.listInvpublisher",
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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_inventory", "list_inventory")
		m4 = append(m4, authz.AuthorizeGenerator("list_inventory", "self"))

		group.GET("controllers-Controller-listInvpublisher", "/publisher/list/single/:id", framework.Mix(ctrl.listInvpublisher, m4...))
		// End route with key 4

		/* Route {
			"Route": "/publisher/list/single/:id/definition",
			"Method": "GET",
			"Function": "Controller.defInvpublisher",
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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_inventory", "list_inventory")
		m5 = append(m5, authz.AuthorizeGenerator("list_inventory", "self"))

		group.GET("controllers-Controller-defInvpublisher", "/publisher/list/single/:id/definition", framework.Mix(ctrl.defInvpublisher, m5...))
		// End route with key 5

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
			"Resource": "list_publisher",
			"Scope": "self"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_publisher", "list_publisher")
		m6 = append(m6, authz.AuthorizeGenerator("list_publisher", "self"))

		group.GET("controllers-Controller-listPublisher", "/publisher/list", framework.Mix(ctrl.listPublisher, m6...))
		// End route with key 6

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
			"Resource": "list_publisher",
			"Scope": "self"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_publisher", "list_publisher")
		m7 = append(m7, authz.AuthorizeGenerator("list_publisher", "self"))

		group.GET("controllers-Controller-defPublisher", "/publisher/list/definition", framework.Mix(ctrl.defPublisher, m7...))
		// End route with key 7

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
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m8 = append(m8, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m8 = append(m8, middleware.PayloadUnMarshallerGenerator(addInventoryPayload{}))
		group.PATCH("controllers-Controller-addPresetPatch", "/addpub/:id", framework.Mix(ctrl.addPresetPatch, m8...))
		// End route with key 8

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
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m9 = append(m9, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m9 = append(m9, middleware.PayloadUnMarshallerGenerator(changeLabelPayload{}))
		group.PUT("controllers-Controller-changeLabelPut", "/:id", framework.Mix(ctrl.changeLabelPut, m9...))
		// End route with key 9

		/* Route {
			"Route": "/inventory/:id",
			"Method": "PATCH",
			"Function": "Controller.inventoryChangeStatusPatch",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "changeStatusPayload",
			"Resource": "edit_inventory",
			"Scope": "self"
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m10 = append(m10, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m10 = append(m10, middleware.PayloadUnMarshallerGenerator(changeStatusPayload{}))
		group.PATCH("controllers-Controller-inventoryChangeStatusPatch", "/inventory/:id", framework.Mix(ctrl.inventoryChangeStatusPatch, m10...))
		// End route with key 10

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
			"Resource": "create_inventory",
			"Scope": "self"
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("create_inventory", "create_inventory")
		m11 = append(m11, authz.AuthorizeGenerator("create_inventory", "self"))

		// Make sure payload is the last middleware
		m11 = append(m11, middleware.PayloadUnMarshallerGenerator(createInventoryPayload{}))
		group.POST("controllers-Controller-createPresetPost", "/create", framework.Mix(ctrl.createPresetPost, m11...))
		// End route with key 11

		/* Route {
			"Route": "/duplicate",
			"Method": "POST",
			"Function": "Controller.duplicatePost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "duplicateInventoryPayload",
			"Resource": "duplicate_inventory",
			"Scope": "self"
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("duplicate_inventory", "duplicate_inventory")
		m12 = append(m12, authz.AuthorizeGenerator("duplicate_inventory", "self"))

		// Make sure payload is the last middleware
		m12 = append(m12, middleware.PayloadUnMarshallerGenerator(duplicateInventoryPayload{}))
		group.POST("controllers-Controller-duplicatePost", "/duplicate", framework.Mix(ctrl.duplicatePost, m12...))
		// End route with key 12

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
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_inventory", "edit_inventory")
		m13 = append(m13, authz.AuthorizeGenerator("edit_inventory", "self"))

		// Make sure payload is the last middleware
		m13 = append(m13, middleware.PayloadUnMarshallerGenerator(removeInventoryPayload{}))
		group.PATCH("controllers-Controller-removePresetPatch", "/removepub/:id", framework.Mix(ctrl.removePresetPatch, m13...))
		// End route with key 13

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
