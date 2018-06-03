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
			"Resource": "daily_campaign",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("daily_campaign", "daily_campaign")
		m0 = append(m0, authz.AuthorizeGenerator("daily_campaign", "self"))

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
			"Resource": "daily_campaign",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("daily_campaign", "daily_campaign")
		m1 = append(m1, authz.AuthorizeGenerator("daily_campaign", "self"))

		group.GET("controllers-Controller-defCampaigndaily", "/daily/:id/definition", framework.Mix(c.defCampaigndaily, m1...))
		// End route with key 1

		/* Route {
			"Route": "/log/:id",
			"Method": "GET",
			"Function": "Controller.listCampaignlog",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "log_campaign",
			"Scope": "self"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("log_campaign", "log_campaign")
		m2 = append(m2, authz.AuthorizeGenerator("log_campaign", "self"))

		group.GET("controllers-Controller-listCampaignlog", "/log/:id", framework.Mix(c.listCampaignlog, m2...))
		// End route with key 2

		/* Route {
			"Route": "/log/:id/definition",
			"Method": "GET",
			"Function": "Controller.defCampaignlog",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "log_campaign",
			"Scope": "self"
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("log_campaign", "log_campaign")
		m3 = append(m3, authz.AuthorizeGenerator("log_campaign", "self"))

		group.GET("controllers-Controller-defCampaignlog", "/log/:id/definition", framework.Mix(c.defCampaignlog, m3...))
		// End route with key 3

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
			"Resource": "list_campaign",
			"Scope": "self"
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_campaign", "list_campaign")
		m4 = append(m4, authz.AuthorizeGenerator("list_campaign", "self"))

		group.GET("controllers-Controller-listCampaigns", "/list", framework.Mix(c.listCampaigns, m4...))
		// End route with key 4

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
			"Resource": "list_campaign",
			"Scope": "self"
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_campaign", "list_campaign")
		m5 = append(m5, authz.AuthorizeGenerator("list_campaign", "self"))

		group.GET("controllers-Controller-defCampaigns", "/list/definition", framework.Mix(c.defCampaigns, m5...))
		// End route with key 5

		/* Route {
			"Route": "/status-list",
			"Method": "GET",
			"Function": "Controller.listCampaigns_creative",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "list_campaign",
			"Scope": "superGlobal"
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_campaign", "list_campaign")
		m6 = append(m6, authz.AuthorizeGenerator("list_campaign", "superGlobal"))

		group.GET("controllers-Controller-listCampaigns_creative", "/status-list", framework.Mix(c.listCampaigns_creative, m6...))
		// End route with key 6

		/* Route {
			"Route": "/status-list/definition",
			"Method": "GET",
			"Function": "Controller.defCampaigns_creative",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "list_campaign",
			"Scope": "superGlobal"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("list_campaign", "list_campaign")
		m7 = append(m7, authz.AuthorizeGenerator("list_campaign", "superGlobal"))

		group.GET("controllers-Controller-defCampaigns_creative", "/status-list/definition", framework.Mix(c.defCampaigns_creative, m7...))
		// End route with key 7

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
			"Resource": "graph_campaign",
			"Scope": "self"
		} with key 8 */
		m8 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("graph_campaign", "graph_campaign")
		m8 = append(m8, authz.AuthorizeGenerator("graph_campaign", "self"))

		group.GET("controllers-Controller-graphChartall", "/graph/all", framework.Mix(c.graphChartall, m8...))
		// End route with key 8

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
			"Resource": "graph_daily_campaign",
			"Scope": "self"
		} with key 9 */
		m9 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("graph_daily_campaign", "graph_daily_campaign")
		m9 = append(m9, authz.AuthorizeGenerator("graph_daily_campaign", "self"))

		group.GET("controllers-Controller-graphChartdaily", "/graph/daily/:id", framework.Mix(c.graphChartdaily, m9...))
		// End route with key 9

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
		} with key 10 */
		m10 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_publisher", "campaign_publisher")
		m10 = append(m10, authz.AuthorizeGenerator("campaign_publisher", "self"))

		group.GET("controllers-Controller-listPublisherdetails", "/publisher-details/:id", framework.Mix(c.listPublisherdetails, m10...))
		// End route with key 10

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
		} with key 11 */
		m11 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_publisher", "campaign_publisher")
		m11 = append(m11, authz.AuthorizeGenerator("campaign_publisher", "self"))

		group.GET("controllers-Controller-defPublisherdetails", "/publisher-details/:id/definition", framework.Mix(c.defPublisherdetails, m11...))
		// End route with key 11

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
		} with key 12 */
		m12 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("archive_campaign", "archive_campaign")
		m12 = append(m12, authz.AuthorizeGenerator("archive_campaign", "self"))

		group.PATCH("controllers-Controller-archivePatch", "/archive/:id", framework.Mix(c.archivePatch, m12...))
		// End route with key 12

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
		} with key 13 */
		m13 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m13 = append(m13, authz.AuthorizeGenerator("edit_campaign", "self"))

		// Make sure payload is the last middleware
		m13 = append(m13, middleware.PayloadUnMarshallerGenerator(assignInventoryPayload{}))
		group.PUT("controllers-Controller-assignInventoryPut", "/inventory/:id", framework.Mix(c.assignInventoryPut, m13...))
		// End route with key 13

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
		} with key 14 */
		m14 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_attributes", "edit_attributes")
		m14 = append(m14, authz.AuthorizeGenerator("edit_attributes", "self"))

		// Make sure payload is the last middleware
		m14 = append(m14, middleware.PayloadUnMarshallerGenerator(attributesPayload{}))
		group.PUT("controllers-Controller-attributesPut", "/attributes/:id", framework.Mix(c.attributesPut, m14...))
		// End route with key 14

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
		} with key 15 */
		m15 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m15 = append(m15, authz.AuthorizeGenerator("edit_campaign", "self"))

		// Make sure payload is the last middleware
		m15 = append(m15, middleware.PayloadUnMarshallerGenerator(campaignBase{}))
		group.PUT("controllers-Controller-updateBasePut", "/base/:id", framework.Mix(c.updateBasePut, m15...))
		// End route with key 15

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
		} with key 16 */
		m16 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_budget", "edit_budget")
		m16 = append(m16, authz.AuthorizeGenerator("edit_budget", "self"))

		// Make sure payload is the last middleware
		m16 = append(m16, middleware.PayloadUnMarshallerGenerator(budgetPayload{}))
		group.PUT("controllers-Controller-budgetPut", "/budget/:id", framework.Mix(c.budgetPut, m16...))
		// End route with key 16

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
		} with key 17 */
		m17 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_campaign_status", "change_campaign_status")
		m17 = append(m17, authz.AuthorizeGenerator("change_campaign_status", "self"))

		// Make sure payload is the last middleware
		m17 = append(m17, middleware.PayloadUnMarshallerGenerator(changeCampaignStatus{}))
		group.PATCH("controllers-Controller-changeStatusPatch", "/status/:id", framework.Mix(c.changeStatusPatch, m17...))
		// End route with key 17

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
			"Payload": "",
			"Resource": "copy_campaign",
			"Scope": "self"
		} with key 18 */
		m18 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("copy_campaign", "copy_campaign")
		m18 = append(m18, authz.AuthorizeGenerator("copy_campaign", "self"))

		group.PATCH("controllers-Controller-copyCampaignPatch", "/copy/:id", framework.Mix(c.copyCampaignPatch, m18...))
		// End route with key 18

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
			"Resource": "create_campaign",
			"Scope": "self"
		} with key 19 */
		m19 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("create_campaign", "create_campaign")
		m19 = append(m19, authz.AuthorizeGenerator("create_campaign", "self"))

		// Make sure payload is the last middleware
		m19 = append(m19, middleware.PayloadUnMarshallerGenerator(createCampaignPayload{}))
		group.POST("controllers-Controller-createBasePost", "/create", framework.Mix(c.createBasePost, m19...))
		// End route with key 19

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
		} with key 20 */
		m20 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_campaign", "edit_campaign")
		m20 = append(m20, authz.AuthorizeGenerator("edit_campaign", "self"))

		group.PUT("controllers-Controller-finalizePut", "/finalize/:id", framework.Mix(c.finalizePut, m20...))
		// End route with key 20

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
		} with key 21 */
		m21 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_campaign", "get_campaign")
		m21 = append(m21, authz.AuthorizeGenerator("get_campaign", "self"))

		group.GET("controllers-Controller-getGet", "/get/:id", framework.Mix(c.getGet, m21...))
		// End route with key 21

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
		} with key 22 */
		m22 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_creative", "get_creative")
		m22 = append(m22, authz.AuthorizeGenerator("get_creative", "self"))

		group.GET("controllers-Controller-getCreativeByCampaignGet", "/creative/:id", framework.Mix(c.getCreativeByCampaignGet, m22...))
		// End route with key 22

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
		} with key 23 */
		m23 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_campaign", "get_campaign")
		m23 = append(m23, authz.AuthorizeGenerator("get_campaign", "self"))

		group.GET("controllers-Controller-getCampaignProgressGet", "/progress/:id", framework.Mix(c.getCampaignProgressGet, m23...))
		// End route with key 23

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
