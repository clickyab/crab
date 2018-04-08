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
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_list", "campaign_list")
		m0 = append(m0, authz.AuthorizeGenerator("campaign_list", "self"))

		group.GET("controllers-Controller-listCampaign", "/list", framework.Mix(c.listCampaign, m0...))
		// End route with key 0

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
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_list", "campaign_list")
		m1 = append(m1, authz.AuthorizeGenerator("campaign_list", "self"))

		group.GET("controllers-Controller-defCampaign", "/list/definition", framework.Mix(c.defCampaign, m1...))
		// End route with key 1

		/* Route {
			"Route": "/daily/:id",
			"Method": "GET",
			"Function": "Controller.listCampaigndaily",
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
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_list", "campaign_list")
		m2 = append(m2, authz.AuthorizeGenerator("campaign_list", "self"))

		group.GET("controllers-Controller-listCampaigndaily", "/daily/:id", framework.Mix(c.listCampaigndaily, m2...))
		// End route with key 2

		/* Route {
			"Route": "/daily/:id/definition",
			"Method": "GET",
			"Function": "Controller.defCampaigndaily",
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
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_list", "campaign_list")
		m3 = append(m3, authz.AuthorizeGenerator("campaign_list", "self"))

		group.GET("controllers-Controller-defCampaigndaily", "/daily/:id/definition", framework.Mix(c.defCampaigndaily, m3...))
		// End route with key 3

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
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_graph", "campaign_graph")
		m4 = append(m4, authz.AuthorizeGenerator("campaign_graph", "self"))

		group.GET("controllers-Controller-graphChart", "/graph/all", framework.Mix(c.graphChart, m4...))
		// End route with key 4

		/* Route {
			"Route": "/attributes/:id",
			"Method": "PUT",
			"Function": "Controller.attributesPut",
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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(attributesPayload{}))
		group.PUT("controllers-Controller-attributesPut", "/attributes/:id", framework.Mix(c.attributesPut, m5...))
		// End route with key 5

		/* Route {
			"Route": "/budget/:id",
			"Method": "PUT",
			"Function": "Controller.budgetPut",
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
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m6 = append(m6, middleware.PayloadUnMarshallerGenerator(budgetPayload{}))
		group.PUT("controllers-Controller-budgetPut", "/budget/:id", framework.Mix(c.budgetPut, m6...))
		// End route with key 6

		/* Route {
			"Route": "/create",
			"Method": "POST",
			"Function": "Controller.createBasePost",
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
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m7 = append(m7, middleware.PayloadUnMarshallerGenerator(createCampaignPayload{}))
		group.POST("controllers-Controller-createBasePost", "/create", framework.Mix(c.createBasePost, m7...))
		// End route with key 7

		/* Route {
			"Route": "/base/:id",
			"Method": "PUT",
			"Function": "Controller.updateBasePut",
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
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit-campaign", "edit-campaign")
		m8 = append(m8, authz.AuthorizeGenerator("edit-campaign", "self"))

		// Make sure payload is the last middleware
		m8 = append(m8, middleware.PayloadUnMarshallerGenerator(campaignStatus{}))
		group.PUT("controllers-Controller-updateBasePut", "/base/:id", framework.Mix(c.updateBasePut, m8...))
		// End route with key 8

		/* Route {
			"Route": "/finalize/:id",
			"Method": "PUT",
			"Function": "Controller.finalizePut",
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
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.PUT("controllers-Controller-finalizePut", "/finalize/:id", framework.Mix(c.finalizePut, m9...))
		// End route with key 9

		/* Route {
			"Route": "/get/:id",
			"Method": "GET",
			"Function": "Controller.getGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_campaign",
			"Scope": "self"
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_campaign", "get_campaign")
		m10 = append(m10, authz.AuthorizeGenerator("get_campaign", "self"))

		group.GET("controllers-Controller-getGet", "/get/:id", framework.Mix(c.getGet, m10...))
		// End route with key 10

		/* Route {
			"Route": "/:id/:stat",
			"Method": "PATCH",
			"Function": "Controller.archivePatch",
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
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_campaign", "change_campaign")
		m11 = append(m11, authz.AuthorizeGenerator("change_campaign", "self"))

		group.PATCH("controllers-Controller-archivePatch", "/:id/:stat", framework.Mix(c.archivePatch, m11...))
		// End route with key 11

		/* Route {
			"Route": "/:id",
			"Method": "PATCH",
			"Function": "Controller.copyCampaignPatch",
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
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("copy_campaign", "copy_campaign")
		m12 = append(m12, authz.AuthorizeGenerator("copy_campaign", "self"))

		// Make sure payload is the last middleware
		m12 = append(m12, middleware.PayloadUnMarshallerGenerator(copyCampaignPayload{}))
		group.PATCH("controllers-Controller-copyCampaignPatch", "/:id", framework.Mix(c.copyCampaignPatch, m12...))
		// End route with key 12

		/* Route {
			"Route": "/get/:id/ad",
			"Method": "GET",
			"Function": "Controller.getCampaignAdsGet",
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
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_banner", "get_banner")
		m13 = append(m13, authz.AuthorizeGenerator("get_banner", "self"))

		group.GET("controllers-Controller-getCampaignAdsGet", "/get/:id/ad", framework.Mix(c.getCampaignAdsGet, m13...))
		// End route with key 13

		/* Route {
			"Route": "/native/fetch",
			"Method": "POST",
			"Function": "Controller.getNativeDataPost",
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
		} with key 14 */
		m14 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m14 = append(m14, middleware.PayloadUnMarshallerGenerator(getNativeDataPayload{}))
		group.POST("controllers-Controller-getNativeDataPost", "/native/fetch", framework.Mix(c.getNativeDataPost, m14...))
		// End route with key 14

		/* Route {
			"Route": "/wb/:id",
			"Method": "PUT",
			"Function": "Controller.updateWhiteBlackListPut",
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
		} with key 15 */
		m15 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		// Make sure payload is the last middleware
		m15 = append(m15, middleware.PayloadUnMarshallerGenerator(whiteBlackPayload{}))
		group.PUT("controllers-Controller-updateWhiteBlackListPut", "/wb/:id", framework.Mix(c.updateWhiteBlackListPut, m15...))
		// End route with key 15

		/* Route {
			"Route": "/wblist/:id",
			"Method": "DELETE",
			"Function": "Controller.deleteWhiteBlackListDelete",
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
		} with key 16 */
		m16 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		group.DELETE("controllers-Controller-deleteWhiteBlackListDelete", "/wblist/:id", framework.Mix(c.deleteWhiteBlackListDelete, m16...))
		// End route with key 16

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
