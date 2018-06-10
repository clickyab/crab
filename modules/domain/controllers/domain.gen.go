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

		group := r.NewGroup("/domain")

		/* Route {
			"Route": "/list",
			"Method": "GET",
			"Function": "Controller.listDomains_data_table",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "list_domain",
			"Scope": "superGlobal"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_domain", "list_domain")
		m0 = append(m0, authz.AuthorizeGenerator("list_domain", "superGlobal"))

		group.GET("controllers-Controller-listDomains_data_table", "/list", framework.Mix(c.listDomains_data_table, m0...))
		// End route with key 0

		/* Route {
			"Route": "/list/definition",
			"Method": "GET",
			"Function": "Controller.defDomains_data_table",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "list_domain",
			"Scope": "superGlobal"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_domain", "list_domain")
		m1 = append(m1, authz.AuthorizeGenerator("list_domain", "superGlobal"))

		group.GET("controllers-Controller-defDomains_data_table", "/list/definition", framework.Mix(c.defDomains_data_table, m1...))
		// End route with key 1

		/* Route {
			"Route": "/change-domain-status/:id",
			"Method": "PUT",
			"Function": "Controller.changeDomainStatusPut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "changeDomainStatusPayload",
			"Resource": "change_domain_status",
			"Scope": "superGlobal"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_domain_status", "change_domain_status")
		m2 = append(m2, authz.AuthorizeGenerator("change_domain_status", "superGlobal"))

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(changeDomainStatusPayload{}))
		group.PUT("controllers-Controller-changeDomainStatusPut", "/change-domain-status/:id", framework.Mix(c.changeDomainStatusPut, m2...))
		// End route with key 2

		/* Route {
			"Route": "/create",
			"Method": "POST",
			"Function": "Controller.createDomainPost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "createDomainPayload",
			"Resource": "create_domain",
			"Scope": "superGlobal"
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("create_domain", "create_domain")
		m3 = append(m3, authz.AuthorizeGenerator("create_domain", "superGlobal"))

		// Make sure payload is the last middleware
		m3 = append(m3, middleware.PayloadUnMarshallerGenerator(createDomainPayload{}))
		group.POST("controllers-Controller-createDomainPost", "/create", framework.Mix(c.createDomainPost, m3...))
		// End route with key 3

		/* Route {
			"Route": "/edit/:id",
			"Method": "PUT",
			"Function": "Controller.editDomainPut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "editDomainPayload",
			"Resource": "edit_domain",
			"Scope": "superGlobal"
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_domain", "edit_domain")
		m4 = append(m4, authz.AuthorizeGenerator("edit_domain", "superGlobal"))

		// Make sure payload is the last middleware
		m4 = append(m4, middleware.PayloadUnMarshallerGenerator(editDomainPayload{}))
		group.PUT("controllers-Controller-editDomainPut", "/edit/:id", framework.Mix(c.editDomainPut, m4...))
		// End route with key 4

		/* Route {
			"Route": "/get/:id",
			"Method": "GET",
			"Function": "Controller.getDomainDetailGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_detail_domain",
			"Scope": "global"
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_detail_domain", "get_detail_domain")
		m5 = append(m5, authz.AuthorizeGenerator("get_detail_domain", "global"))

		group.GET("controllers-Controller-getDomainDetailGet", "/get/:id", framework.Mix(c.getDomainDetailGet, m5...))
		// End route with key 5

		/* Route {
			"Route": "/config/:name",
			"Method": "GET",
			"Function": "Controller.getDomainConfigGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("controllers-Controller-getDomainConfigGet", "/config/:name", framework.Mix(c.getDomainConfigGet, m6...))
		// End route with key 6

		/* Route {
			"Route": "/super-global-config",
			"Method": "GET",
			"Function": "Controller.getGlobalConfigGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_global_config",
			"Scope": "superGlobal"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_global_config", "get_global_config")
		m7 = append(m7, authz.AuthorizeGenerator("get_global_config", "superGlobal"))

		group.GET("controllers-Controller-getGlobalConfigGet", "/super-global-config", framework.Mix(c.getGlobalConfigGet, m7...))
		// End route with key 7

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
