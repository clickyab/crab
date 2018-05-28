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
			"Resource": "god",
			"Scope": "global"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("god", "god")
		m0 = append(m0, authz.AuthorizeGenerator("god", "global"))

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(changeDomainStatusPayload{}))
		group.PUT("controllers-Controller-changeDomainStatusPut", "/change-domain-status/:id", framework.Mix(c.changeDomainStatusPut, m0...))
		// End route with key 0

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
			"Resource": "god",
			"Scope": "global"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("god", "god")
		m1 = append(m1, authz.AuthorizeGenerator("god", "global"))

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(createDomainPayload{}))
		group.POST("controllers-Controller-createDomainPost", "/create", framework.Mix(c.createDomainPost, m1...))
		// End route with key 1

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
			"Resource": "god",
			"Scope": "global"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("god", "god")
		m2 = append(m2, authz.AuthorizeGenerator("god", "global"))

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(editDomainPayload{}))
		group.PUT("controllers-Controller-editDomainPut", "/edit/:id", framework.Mix(c.editDomainPut, m2...))
		// End route with key 2

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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_detail_domain", "get_detail_domain")
		m3 = append(m3, authz.AuthorizeGenerator("get_detail_domain", "global"))

		group.GET("controllers-Controller-getDomainDetailGet", "/get/:id", framework.Mix(c.getDomainDetailGet, m3...))
		// End route with key 3

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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("controllers-Controller-getDomainConfigGet", "/config/:name", framework.Mix(c.getDomainConfigGet, m4...))
		// End route with key 4

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
