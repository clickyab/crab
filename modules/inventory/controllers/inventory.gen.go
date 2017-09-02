// Code generated build with router DO NOT EDIT.

package controllers

import (
	"sync"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
	"github.com/rs/xhandler"
	"github.com/rs/xmux"
)

var once = sync.Once{}

// Routes return the route registered with this
func (ctrl *Controller) Routes(r *xmux.Mux, mountPoint string) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup(mountPoint + "/inventory")

		/* Route {
			"Route": "/:user",
			"Method": "GET",
			"Function": "Controller.Inventories",
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

		group.GET("/:user", xhandler.HandlerFuncC(framework.Mix(ctrl.Inventories, m0...)))
		// End route with key 0

		/* Route {
			"Route": "/:id",
			"Method": "GET",
			"Function": "Controller.Inventory",
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

		group.GET("/:id", xhandler.HandlerFuncC(framework.Mix(ctrl.Inventory, m1...)))
		// End route with key 1

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
