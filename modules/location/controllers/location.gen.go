// Code generated build with router DO NOT EDIT.

package location

import (
	"sync"

	"clickyab.com/crab/modules/domain/middleware/domain"
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

		group := r.NewGroup(mountPoint + "/location")

		/* Route {
			"Route": "/countries",
			"Method": "GET",
			"Function": "Controller.Countries",
			"RoutePkg": "location",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/countries", xhandler.HandlerFuncC(framework.Mix(ctrl.Countries, m0...)))
		// End route with key 0

		/* Route {
			"Route": "/provinces/:country_id",
			"Method": "GET",
			"Function": "Controller.Provinces",
			"RoutePkg": "location",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/provinces/:country_id", xhandler.HandlerFuncC(framework.Mix(ctrl.Provinces, m1...)))
		// End route with key 1

		/* Route {
			"Route": "/cities/:provinces_id",
			"Method": "GET",
			"Function": "Controller.Cities",
			"RoutePkg": "location",
			"RouteMiddleware": null,
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "",
			"Scope": ""
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{}...)

		group.GET("/cities/:provinces_id", xhandler.HandlerFuncC(framework.Mix(ctrl.Cities, m2...)))
		// End route with key 2

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
