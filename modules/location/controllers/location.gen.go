// Code generated build with router DO NOT EDIT.

package location

import (
	"sync"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
)

var once = sync.Once{}

// Routes return the route registered with this
func (ctrl *Controller) Routes(r router.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup("/location")

		/* Route {
			"Route": "/countries",
			"Method": "GET",
			"Function": "Controller.countries",
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

		group.GET("/countries", framework.Mix(ctrl.countries, m0...))
		// End route with key 0

		/* Route {
			"Route": "/provinces/:country_id",
			"Method": "GET",
			"Function": "Controller.provinces",
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

		group.GET("/provinces/:country_id", framework.Mix(ctrl.provinces, m1...))
		// End route with key 1

		/* Route {
			"Route": "/cities/:provinces_id",
			"Method": "GET",
			"Function": "Controller.cities",
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

		group.GET("/cities/:provinces_id", framework.Mix(ctrl.cities, m2...))
		// End route with key 2

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
