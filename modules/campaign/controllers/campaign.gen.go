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
			"Resource": "campaign_daily",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_daily", "campaign_daily")
		m0 = append(m0, authz.AuthorizeGenerator("campaign_daily", "self"))

		group.GET("controllers-Controller-listCampaigndaily", "/daily/:id", framework.Mix(c.listCampaigndaily, m0...))
		// End route with key 0

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
			"Resource": "campaign_daily",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_daily", "campaign_daily")
		m1 = append(m1, authz.AuthorizeGenerator("campaign_daily", "self"))

		group.GET("controllers-Controller-defCampaigndaily", "/daily/:id/definition", framework.Mix(c.defCampaigndaily, m1...))
		// End route with key 1

		/* Route {
			"Route": "/list",
			"Method": "GET",
			"Function": "Controller.listCampaigns",
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

		group.GET("controllers-Controller-listCampaigns", "/list", framework.Mix(c.listCampaigns, m2...))
		// End route with key 2

		/* Route {
			"Route": "/list/definition",
			"Method": "GET",
			"Function": "Controller.defCampaigns",
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

		group.GET("controllers-Controller-defCampaigns", "/list/definition", framework.Mix(c.defCampaigns, m3...))
		// End route with key 3

		/* Route {
			"Route": "/graph/all",
			"Method": "GET",
			"Function": "Controller.graphChartall",
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

		group.GET("controllers-Controller-graphChartall", "/graph/all", framework.Mix(c.graphChartall, m4...))
		// End route with key 4

		/* Route {
			"Route": "/graph/daily/:id",
			"Method": "GET",
			"Function": "Controller.graphChartdaily",
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
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_graph", "campaign_graph")
		m5 = append(m5, authz.AuthorizeGenerator("campaign_graph", "self"))

		group.GET("controllers-Controller-graphChartdaily", "/graph/daily/:id", framework.Mix(c.graphChartdaily, m5...))
		// End route with key 5

		/* Route {
			"Route": "/publisher-details/:id",
			"Method": "GET",
			"Function": "Controller.listPublisherdetails",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "campaign_publisher",
			"Scope": "self"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_publisher", "campaign_publisher")
		m6 = append(m6, authz.AuthorizeGenerator("campaign_publisher", "self"))

		group.GET("controllers-Controller-listPublisherdetails", "/publisher-details/:id", framework.Mix(c.listPublisherdetails, m6...))
		// End route with key 6

		/* Route {
			"Route": "/publisher-details/:id/definition",
			"Method": "GET",
			"Function": "Controller.defPublisherdetails",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "campaign_publisher",
			"Scope": "self"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_publisher", "campaign_publisher")
		m7 = append(m7, authz.AuthorizeGenerator("campaign_publisher", "self"))

		group.GET("controllers-Controller-defPublisherdetails", "/publisher-details/:id/definition", framework.Mix(c.defPublisherdetails, m7...))
		// End route with key 7

		/* Route {
			"Route": "/archive/:id",
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
			"Resource": "archive_campaign",
			"Scope": "self"
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("archive_campaign", "archive_campaign")
		m8 = append(m8, authz.AuthorizeGenerator("archive_campaign", "self"))

		group.PATCH("controllers-Controller-archivePatch", "/archive/:id", framework.Mix(c.archivePatch, m8...))
		// End route with key 8

		/* Route {
			"Route": "/inventory/:id",
			"Method": "PUT",
			"Function": "Controller.assignInventoryPut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "assignInventoryPayload",
			"Resource": "edit_campaign",
			"Scope": "self"
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m9 = append(m9, authz.AuthorizeGenerator("edit_campaign", "self"))

		// Make sure payload is the last middleware
		m9 = append(m9, middleware.PayloadUnMarshallerGenerator(assignInventoryPayload{}))
		group.PUT("controllers-Controller-assignInventoryPut", "/inventory/:id", framework.Mix(c.assignInventoryPut, m9...))
		// End route with key 9

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
			"Resource": "edit_attributes",
			"Scope": "self"
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_attributes", "edit_attributes")
		m10 = append(m10, authz.AuthorizeGenerator("edit_attributes", "self"))

		// Make sure payload is the last middleware
		m10 = append(m10, middleware.PayloadUnMarshallerGenerator(attributesPayload{}))
		group.PUT("controllers-Controller-attributesPut", "/attributes/:id", framework.Mix(c.attributesPut, m10...))
		// End route with key 10

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
			"Payload": "campaignBase",
			"Resource": "edit_campaign",
			"Scope": "self"
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m11 = append(m11, authz.AuthorizeGenerator("edit_campaign", "self"))

		// Make sure payload is the last middleware
		m11 = append(m11, middleware.PayloadUnMarshallerGenerator(campaignBase{}))
		group.PUT("controllers-Controller-updateBasePut", "/base/:id", framework.Mix(c.updateBasePut, m11...))
		// End route with key 11

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
			"Resource": "edit_budget",
			"Scope": "self"
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_budget", "edit_budget")
		m12 = append(m12, authz.AuthorizeGenerator("edit_budget", "self"))

		// Make sure payload is the last middleware
		m12 = append(m12, middleware.PayloadUnMarshallerGenerator(budgetPayload{}))
		group.PUT("controllers-Controller-budgetPut", "/budget/:id", framework.Mix(c.budgetPut, m12...))
		// End route with key 12

		/* Route {
			"Route": "/status/:id",
			"Method": "PATCH",
			"Function": "Controller.changeStatusPatch",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "changeCampaignStatus",
			"Resource": "change_campaign_status",
			"Scope": "self"
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_campaign_status", "change_campaign_status")
		m13 = append(m13, authz.AuthorizeGenerator("change_campaign_status", "self"))

		// Make sure payload is the last middleware
		m13 = append(m13, middleware.PayloadUnMarshallerGenerator(changeCampaignStatus{}))
		group.PATCH("controllers-Controller-changeStatusPatch", "/status/:id", framework.Mix(c.changeStatusPatch, m13...))
		// End route with key 13

		/* Route {
			"Route": "/copy/:id",
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
		} with key 14 */
		m14 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("copy_campaign", "copy_campaign")
		m14 = append(m14, authz.AuthorizeGenerator("copy_campaign", "self"))

		// Make sure payload is the last middleware
		m14 = append(m14, middleware.PayloadUnMarshallerGenerator(copyCampaignPayload{}))
		group.PATCH("controllers-Controller-copyCampaignPatch", "/copy/:id", framework.Mix(c.copyCampaignPatch, m14...))
		// End route with key 14

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
			"Resource": "edit_campaign",
			"Scope": "self"
		} with key 15 */
		m15 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m15 = append(m15, authz.AuthorizeGenerator("edit_campaign", "self"))

		// Make sure payload is the last middleware
		m15 = append(m15, middleware.PayloadUnMarshallerGenerator(createCampaignPayload{}))
		group.POST("controllers-Controller-createBasePost", "/create", framework.Mix(c.createBasePost, m15...))
		// End route with key 15

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
			"Resource": "edit_campaign",
			"Scope": "self"
		} with key 16 */
		m16 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m16 = append(m16, authz.AuthorizeGenerator("edit_campaign", "self"))

		group.PUT("controllers-Controller-finalizePut", "/finalize/:id", framework.Mix(c.finalizePut, m16...))
		// End route with key 16

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
		} with key 17 */
		m17 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_campaign", "get_campaign")
		m17 = append(m17, authz.AuthorizeGenerator("get_campaign", "self"))

		group.GET("controllers-Controller-getGet", "/get/:id", framework.Mix(c.getGet, m17...))
		// End route with key 17

		/* Route {
			"Route": "/creative/:id",
			"Method": "GET",
			"Function": "Controller.getCreativeByCampaignGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_creative",
			"Scope": "self"
		} with key 18 */
		m18 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_creative", "get_creative")
		m18 = append(m18, authz.AuthorizeGenerator("get_creative", "self"))

		group.GET("controllers-Controller-getCreativeByCampaignGet", "/creative/:id", framework.Mix(c.getCreativeByCampaignGet, m18...))
		// End route with key 18

		/* Route {
			"Route": "/progress/:id",
			"Method": "GET",
			"Function": "Controller.getCampaignProgressGet",
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
		} with key 19 */
		m19 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_campaign", "get_campaign")
		m19 = append(m19, authz.AuthorizeGenerator("get_campaign", "self"))

		group.GET("controllers-Controller-getCampaignProgressGet", "/progress/:id", framework.Mix(c.getCampaignProgressGet, m19...))
		// End route with key 19

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
