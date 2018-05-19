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
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("god", "god")
		m0 = append(m0, authz.AuthorizeGenerator("god", "global"))

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(createDomainPayload{}))
		group.POST("controllers-Controller-createDomainPost", "/create", framework.Mix(c.createDomainPost, m0...))
		// End route with key 0

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
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("god", "god")
		m1 = append(m1, authz.AuthorizeGenerator("god", "global"))

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(editDomainPayload{}))
		group.PUT("controllers-Controller-editDomainPut", "/edit/:id", framework.Mix(c.editDomainPut, m1...))
		// End route with key 1

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
