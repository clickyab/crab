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

		group := r.NewGroup("/campaign")

		/* Route {
			"Route": "/attributes/:id",
			"Method": "PUT",
			"Function": "Controller.attributes",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "attributesPayload",
			"Resource": "",
			"Scope": ""
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m0 = append(m0, middleware.PayloadUnMarshallerGenerator(attributesPayload{}))
		group.PUT("/attributes/:id", framework.Mix(c.attributes, m0...))
		// End route with key 0

		/* Route {
			"Route": "/budget/:id",
			"Method": "PUT",
			"Function": "Controller.budget",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "budgetPayload",
			"Resource": "",
			"Scope": ""
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m1 = append(m1, middleware.PayloadUnMarshallerGenerator(budgetPayload{}))
		group.PUT("/budget/:id", framework.Mix(c.budget, m1...))
		// End route with key 1

		/* Route {
			"Route": "/create",
			"Method": "POST",
			"Function": "Controller.createBase",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "createCampaignPayload",
			"Resource": "",
			"Scope": ""
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(createCampaignPayload{}))
		group.POST("/create", framework.Mix(c.createBase, m2...))
		// End route with key 2

		/* Route {
			"Route": "/base/:id",
			"Method": "PUT",
			"Function": "Controller.updateBase",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "campaignStatus",
			"Resource": "edit-campaign",
			"Scope": "self"
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit-campaign", "edit-campaign")
		m3 = append(m3, authz.AuthorizeGenerator("edit-campaign", "self"))

		// Make sure payload is the last middleware
		m3 = append(m3, middleware.PayloadUnMarshallerGenerator(campaignStatus{}))
		group.PUT("/base/:id", framework.Mix(c.updateBase, m3...))
		// End route with key 3

		/* Route {
			"Route": "/finalize/:id",
			"Method": "PUT",
			"Function": "Controller.finalize",
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

		group.PUT("/finalize/:id", framework.Mix(c.finalize, m4...))
		// End route with key 4

		/* Route {
			"Route": "/:id",
			"Method": "GET",
			"Function": "Controller.get",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get-campaign",
			"Scope": "self"
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get-campaign", "get-campaign")
		m5 = append(m5, authz.AuthorizeGenerator("get-campaign", "self"))

		group.GET("/:id", framework.Mix(c.get, m5...))
		// End route with key 5

		/* Route {
			"Route": "/:id/ad",
			"Method": "GET",
			"Function": "Controller.getCampaignAds",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_banner",
			"Scope": "self"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_banner", "get_banner")
		m6 = append(m6, authz.AuthorizeGenerator("get_banner", "self"))

		group.GET("/:id/ad", framework.Mix(c.getCampaignAds, m6...))
		// End route with key 6

		/* Route {
			"Route": "/wb/:id",
			"Method": "PUT",
			"Function": "Controller.updateWhiteBlackList",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "whiteBlackPayload",
			"Resource": "",
			"Scope": ""
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(whiteBlackPayload{}))
		group.PUT("/wb/:id", framework.Mix(c.updateWhiteBlackList, m7...))
		// End route with key 7

		/* Route {
			"Route": "/wblist/:id",
			"Method": "DELETE",
			"Function": "Controller.deleteWhiteBlackList",
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
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.DELETE("/wblist/:id", framework.Mix(c.deleteWhiteBlackList, m8...))
		// End route with key 8

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
