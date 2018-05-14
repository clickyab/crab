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

		group := r.NewGroup("/ad")

		/* Route {
			"Route": "/campaign/:id",
			"Method": "GET",
			"Function": "Controller.listCreativecampaignreport",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "campaign_creative",
			"Scope": "self"
		} with key 0 */
		m0 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_creative", "campaign_creative")
		m0 = append(m0, authz.AuthorizeGenerator("campaign_creative", "self"))

		group.GET("controllers-Controller-listCreativecampaignreport", "/campaign/:id", framework.Mix(c.listCreativecampaignreport, m0...))
		// End route with key 0

		/* Route {
			"Route": "/campaign/:id/definition",
			"Method": "GET",
			"Function": "Controller.defCreativecampaignreport",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "u",
			"Payload": "",
			"Resource": "campaign_creative",
			"Scope": "self"
		} with key 1 */
		m1 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("campaign_creative", "campaign_creative")
		m1 = append(m1, authz.AuthorizeGenerator("campaign_creative", "self"))

		group.GET("controllers-Controller-defCreativecampaignreport", "/campaign/:id/definition", framework.Mix(c.defCreativecampaignreport, m1...))
		// End route with key 1

		/* Route {
			"Route": "/native",
			"Method": "POST",
			"Function": "Controller.addNativeCreativePost",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "createNativePayload",
			"Resource": "add_creative",
			"Scope": "self"
		} with key 2 */
		m2 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("add_creative", "add_creative")
		m2 = append(m2, authz.AuthorizeGenerator("add_creative", "self"))

		// Make sure payload is the last middleware
		m2 = append(m2, middleware.PayloadUnMarshallerGenerator(createNativePayload{}))
		group.POST("controllers-Controller-addNativeCreativePost", "/native", framework.Mix(c.addNativeCreativePost, m2...))
		// End route with key 2

		/* Route {
			"Route": "/campaign-creative-status/:id",
			"Method": "PATCH",
			"Function": "Controller.changeCampaignCreativeStatusPatch",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "changeStatus",
			"Resource": "change_creative_status",
			"Scope": "global"
		} with key 3 */
		m3 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_creative_status", "change_creative_status")
		m3 = append(m3, authz.AuthorizeGenerator("change_creative_status", "global"))

		// Make sure payload is the last middleware
		m3 = append(m3, middleware.PayloadUnMarshallerGenerator(changeStatus{}))
		group.PATCH("controllers-Controller-changeCampaignCreativeStatusPatch", "/campaign-creative-status/:id", framework.Mix(c.changeCampaignCreativeStatusPatch, m3...))
		// End route with key 3

		/* Route {
			"Route": "/change-creatives-status/:id",
			"Method": "PUT",
			"Function": "Controller.changeCreativesStatusPut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "changeStatusPayload",
			"Resource": "change_creatives_status",
			"Scope": "global"
		} with key 4 */
		m4 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("change_creatives_status", "change_creatives_status")
		m4 = append(m4, authz.AuthorizeGenerator("change_creatives_status", "global"))

		// Make sure payload is the last middleware
		m4 = append(m4, middleware.PayloadUnMarshallerGenerator(changeStatusPayload{}))
		group.PUT("controllers-Controller-changeCreativesStatusPut", "/change-creatives-status/:id", framework.Mix(c.changeCreativesStatusPut, m4...))
		// End route with key 4

		/* Route {
			"Route": "/native/:id",
			"Method": "PUT",
			"Function": "Controller.editNativeCreativePut",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "editNativePayload",
			"Resource": "edit_creative",
			"Scope": "self"
		} with key 5 */
		m5 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("edit_creative", "edit_creative")
		m5 = append(m5, authz.AuthorizeGenerator("edit_creative", "self"))

		// Make sure payload is the last middleware
		m5 = append(m5, middleware.PayloadUnMarshallerGenerator(editNativePayload{}))
		group.PUT("controllers-Controller-editNativeCreativePut", "/native/:id", framework.Mix(c.editNativeCreativePut, m5...))
		// End route with key 5

		/* Route {
			"Route": "/creative/:id",
			"Method": "GET",
			"Function": "Controller.getCreativeGet",
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
		} with key 6 */
		m6 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_creative", "get_creative")
		m6 = append(m6, authz.AuthorizeGenerator("get_creative", "self"))

		group.GET("controllers-Controller-getCreativeGet", "/creative/:id", framework.Mix(c.getCreativeGet, m6...))
		// End route with key 6

		/* Route {
			"Route": "/creative-reject-reasons",
			"Method": "GET",
			"Function": "Controller.getCreativeRejectReasonsGet",
			"RoutePkg": "controllers",
			"RouteMiddleware": [
				"authz.Authenticate"
			],
			"RouteFuncMiddleware": "",
			"RecType": "Controller",
			"RecName": "c",
			"Payload": "",
			"Resource": "get_creative_reject_reason",
			"Scope": "self"
		} with key 7 */
		m7 := append(groupMiddleware, []framework.Middleware{
			authz.Authenticate,
		}...)

		permission.Register("get_creative_reject_reason", "get_creative_reject_reason")
		m7 = append(m7, authz.AuthorizeGenerator("get_creative_reject_reason", "self"))

		group.GET("controllers-Controller-getCreativeRejectReasonsGet", "/creative-reject-reasons", framework.Mix(c.getCreativeRejectReasonsGet, m7...))
		// End route with key 7

		initializer.DoInitialize(c)
	})
}

func init() {
	router.Register(&Controller{})
}
