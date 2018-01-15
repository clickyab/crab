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
		group.PUT("controllers-Controller-attributes", "/attributes/:id", framework.Mix(c.attributes, m0...))
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
		group.PUT("controllers-Controller-budget", "/budget/:id", framework.Mix(c.budget, m1...))
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
		group.POST("controllers-Controller-createBase", "/create", framework.Mix(c.createBase, m2...))
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
		group.PUT("controllers-Controller-updateBase", "/base/:id", framework.Mix(c.updateBase, m3...))
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

		group.PUT("controllers-Controller-finalize", "/finalize/:id", framework.Mix(c.finalize, m4...))
		// End route with key 4

		/* Route {
			"Route": "/get/:id",
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

		group.GET("controllers-Controller-get", "/get/:id", framework.Mix(c.get, m5...))
		// End route with key 5

		/* Route {
			"Route": "/list",
			"Method": "GET",
			"Function": "Controller.listCampaign",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "campaign_list",
			"Scope": "self"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_list", "campaign_list")
		m6 = append(m6, authz.AuthorizeGenerator("campaign_list", "self"))

		group.GET("controllers-Controller-listCampaign", "/list", framework.Mix(c.listCampaign, m6...))
		// End route with key 6

		/* Route {
			"Route": "/list/definition",
			"Method": "GET",
			"Function": "Controller.defCampaign",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "campaign_list",
			"Scope": "self"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_list", "campaign_list")
		m7 = append(m7, authz.AuthorizeGenerator("campaign_list", "self"))

		group.GET("controllers-Controller-defCampaign", "/list/definition", framework.Mix(c.defCampaign, m7...))
		// End route with key 7

		/* Route {
			"Route": "/:id/:stat",
			"Method": "PATCH",
			"Function": "Controller.archive",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "change_campaign",
			"Scope": "self"
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_campaign", "change_campaign")
		m8 = append(m8, authz.AuthorizeGenerator("change_campaign", "self"))

		group.PATCH("controllers-Controller-archive", "/:id/:stat", framework.Mix(c.archive, m8...))
		// End route with key 8

		/* Route {
			"Route": "/graph/all",
			"Method": "GET",
			"Function": "Controller.graphChart",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "ctrl",
			"Payload": "",
			"Resource": "campaign_graph",
			"Scope": "self"
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_graph", "campaign_graph")
		m9 = append(m9, authz.AuthorizeGenerator("campaign_graph", "self"))

		group.GET("controllers-Controller-graphChart", "/graph/all", framework.Mix(c.graphChart, m9...))
		// End route with key 9

		/* Route {
			"Route": "/:id",
			"Method": "PATCH",
			"Function": "Controller.copyCampaign",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "copyCampaignPayload",
			"Resource": "copy_campaign",
			"Scope": "self"
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("copy_campaign", "copy_campaign")
		m10 = append(m10, authz.AuthorizeGenerator("copy_campaign", "self"))

		// Make sure payload is the last middleware
		m10 = append(m10, middleware.PayloadUnMarshallerGenerator(copyCampaignPayload{}))
		group.PATCH("controllers-Controller-copyCampaign", "/:id", framework.Mix(c.copyCampaign, m10...))
		// End route with key 10

		/* Route {
			"Route": "/get/:id/ad",
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
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_banner", "get_banner")
		m11 = append(m11, authz.AuthorizeGenerator("get_banner", "self"))

		group.GET("controllers-Controller-getCampaignAds", "/get/:id/ad", framework.Mix(c.getCampaignAds, m11...))
		// End route with key 11

		/* Route {
			"Route": "/native/fetch",
			"Method": "POST",
			"Function": "Controller.getNativeData",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "getNativeDataPayload",
			"Resource": "",
			"Scope": ""
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m12 = append(m12, middleware.PayloadUnMarshallerGenerator(getNativeDataPayload{}))
		group.POST("controllers-Controller-getNativeData", "/native/fetch", framework.Mix(c.getNativeData, m12...))
		// End route with key 12

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
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m13 = append(m13, middleware.PayloadUnMarshallerGenerator(whiteBlackPayload{}))
		group.PUT("controllers-Controller-updateWhiteBlackList", "/wb/:id", framework.Mix(c.updateWhiteBlackList, m13...))
		// End route with key 13

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
		} with key 14 */
		m14 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.DELETE("controllers-Controller-deleteWhiteBlackList", "/wblist/:id", framework.Mix(c.deleteWhiteBlackList, m14...))
		// End route with key 14

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
