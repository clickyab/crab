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
func (ctrl *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		groupMiddleware := []framework.Middleware{
			domain.Access,
		}

		group := r.NewGroup("/location")

		/* Route {
			"Route": "/countries",
			"Method": "GET",
			"Function": "Controller.countriesGet",
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

		group.GET("location-Controller-countriesGet", "/countries", framework.Mix(ctrl.countriesGet, m0...))
		// End route with key 0

		/* Route {
			"Route": "/provinces/:country_id",
			"Method": "GET",
			"Function": "Controller.provincesGet",
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

		group.GET("location-Controller-provincesGet", "/provinces/:country_id", framework.Mix(ctrl.provincesGet, m1...))
		// End route with key 1

		/* Route {
			"Route": "/cities/:province",
			"Method": "GET",
			"Function": "Controller.citiesGet",
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

		group.GET("location-Controller-citiesGet", "/cities/:province", framework.Mix(ctrl.citiesGet, m2...))
		// End route with key 2

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
