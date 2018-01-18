// Code generated build with router DO NOT EDIT.

package controllers

import (
	"sync"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
)

var once = sync.Once{}

// Routes return the route registered with this
func (c *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup("/asset")

		/* Route {
			"Route": "/browser",
			"Method": "GET",
			"Function": "Controller.browserGet",
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
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-browserGet", "/browser", framework.Mix(c.browserGet, m0...))
		// End route with key 0

		/* Route {
			"Route": "/category",
			"Method": "GET",
			"Function": "Controller.categoryGet",
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
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-categoryGet", "/category", framework.Mix(c.categoryGet, m1...))
		// End route with key 1

		/* Route {
			"Route": "/isp/:kind",
			"Method": "GET",
			"Function": "Controller.ispGet",
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
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-ispGet", "/isp/:kind", framework.Mix(c.ispGet, m2...))
		// End route with key 2

		/* Route {
			"Route": "/manufacturers",
			"Method": "GET",
			"Function": "Controller.manufacturerGet",
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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-manufacturerGet", "/manufacturers", framework.Mix(c.manufacturerGet, m3...))
		// End route with key 3

		/* Route {
			"Route": "/os",
			"Method": "GET",
			"Function": "Controller.osGet",
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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-osGet", "/os", framework.Mix(c.osGet, m4...))
		// End route with key 4

		/* Route {
			"Route": "/platform",
			"Method": "GET",
			"Function": "Controller.platformGet",
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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-platformGet", "/platform", framework.Mix(c.platformGet, m5...))
		// End route with key 5

		/* Route {
			"Route": "/region",
			"Method": "GET",
			"Function": "Controller.regionGet",
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
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.GET("controllers-Controller-regionGet", "/region", framework.Mix(c.regionGet, m6...))
		// End route with key 6

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
