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
			"Route": "/:banner_type/:id",
			"Method": "POST",
			"Function": "Controller.assignNormalBanner",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "assignBannerPayload",
			"Resource": "assign_banner",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("assign_banner", "assign_banner")
		m0 = append(m0, authz.AuthorizeGenerator("assign_banner", "self"))

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(assignBannerPayload{}))
		group.POST("controllers-Controller-assignNormalBanner", "/:banner_type/:id", framework.Mix(c.assignNormalBanner, m0...))
		// End route with key 0

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
