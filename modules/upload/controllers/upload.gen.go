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

		group := r.NewGroup("/upload")

		/* Route {
			"Route": "/banner/:module",
			"Method": "POST",
			"Function": "Controller.Upload",
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

		group.POST("controllers-Controller-Upload", "/banner/:module", framework.Mix(c.Upload, m0...))
		// End route with key 0

		/* Route {
			"Route": "/video",
			"Method": "POST",
			"Function": "Controller.videoUpload",
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

		group.POST("controllers-Controller-videoUpload", "/video", framework.Mix(c.videoUpload, m1...))
		// End route with key 1

		/* Route {
			"Route": "/video/:id",
			"Method": "GET",
			"Function": "Controller.getVideoReady",
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

		group.GET("controllers-Controller-getVideoReady", "/video/:id", framework.Mix(c.getVideoReady, m2...))
		// End route with key 2

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
