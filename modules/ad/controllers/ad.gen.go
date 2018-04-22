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

		group := r.NewGroup("/ad")

		/* Route {
			"Route": "/native",
			"Method": "POST",
			"Function": "Controller.addNativeCreativePost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "createNativePayload",
			"Resource": "add_creative",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("add_creative", "add_creative")
		m0 = append(m0, authz.AuthorizeGenerator("add_creative", "self"))

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(createNativePayload{}))
		group.POST("controllers-Controller-addNativeCreativePost", "/native", framework.Mix(c.addNativeCreativePost, m0...))
		// End route with key 0

		/* Route {
			"Route": "/native/:id",
			"Method": "PUT",
			"Function": "Controller.editNativeCreativePut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "editNativePayload",
			"Resource": "edit_creative",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_creative", "edit_creative")
		m1 = append(m1, authz.AuthorizeGenerator("edit_creative", "self"))

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(editNativePayload{}))
		group.PUT("controllers-Controller-editNativeCreativePut", "/native/:id", framework.Mix(c.editNativeCreativePut, m1...))
		// End route with key 1

		/* Route {
			"Route": "/creative/:id",
			"Method": "GET",
			"Function": "Controller.getCreativeGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_creative",
			"Scope": "self"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_creative", "get_creative")
		m2 = append(m2, authz.AuthorizeGenerator("get_creative", "self"))

		group.GET("controllers-Controller-getCreativeGet", "/creative/:id", framework.Mix(c.getCreativeGet, m2...))
		// End route with key 2

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
