// Code generated build with datatable DO NOT EDIT.

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listCampaignResponse struct {
	Total   int64                      `json:"total"`
	Data    orm.CampaignDataTableArray `json:"data"`
	Page    int                        `json:"page"`
	PerPage int                        `json:"per_page"`
	Hash    string                     `json:"hash"`
}

type listCampaignDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listCampaignDefinition permission.Columns
	Campaigntmp            = []byte{}
)

// @Route {
// 		url = /list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = campaign_list:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "created_at","start_at","max_bid"
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_cost_type_ = string , filter the cost_type field valid values are "cpm","cpc","cpa"
//		_title_ = string , search the title field
//		_owner_email_ = string , search the owner_email field
//		200 = listCampaignResponse
// }
func (u *Controller) listCampaign(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)
	dateRange := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" && orm.CampaignKind(e).IsValid() {
		filter["cp.kind"] = e
	}

	if e := r.URL.Query().Get("cost_type"); e != "" && orm.Strategy(e).IsValid() {
		filter["cp.cost_type"] = e
	}

	//add date filter
	if e := r.URL.Query().Get("from"); e != "" {
		dateRange["from-cp.created_at"] = e
	}

	if e := r.URL.Query().Get("to"); e != "" {
		dateRange["to-cp.created_at"] = e
	}

	search := make(map[string]string)

	if e := r.URL.Query().Get(""); e != "" {
		search["cp.title"] = e
	}

	if e := r.URL.Query().Get(""); e != "" {
		search["owner.email"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "created_at", "start_at", "max_bid") {
		sort = ""
	}
	order := strings.ToUpper(parts[1])
	if !array.StringInArray(order, "ASC", "DESC") {
		order = "ASC"
	}

	params := make(map[string]string)
	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}

	pc := permission.NewInterfaceComplete(usr, usr.ID, "campaign_list", "self", domain.ID)
	dt, cnt, err := m.FillCampaignDataTableArray(pc, filter, dateRange, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listCampaignResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Campaigntmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /list/definition
//		method = get
//		resource = campaign_list:self
//		200 = listCampaignDefResponse
// }
func (u *Controller) defCampaign(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Campaigntmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listCampaignDefResponse{Checkable: false, SearchKey: "", Multiselect: false, DateFilter: "cp.created_at", Hash: hash, Columns: listCampaignDefinition},
	)
}

func init() {
	Campaigntmp = []byte(` [
		{
			"data": "id",
			"name": "ID",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "ID",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "created_at",
			"name": "CreatedAt",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "CreatedAt",
			"type": "date",
			"filter_valid_map": null
		},
		{
			"data": "active",
			"name": "Active",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Active",
			"type": "bool",
			"filter_valid_map": null
		},
		{
			"data": "kind",
			"name": "Kind",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": true,
			"title": "Kind",
			"type": "enum",
			"filter_valid_map": {
				"app": "AppCampaign",
				"web": "WebCampaign"
			}
		},
		{
			"data": "status",
			"name": "Status",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Status",
			"type": "bool",
			"filter_valid_map": null
		},
		{
			"data": "start_at",
			"name": "StartAt",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "StartAt",
			"type": "date",
			"filter_valid_map": null
		},
		{
			"data": "end_at",
			"name": "EndAt",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "EndAt",
			"type": "date",
			"filter_valid_map": null
		},
		{
			"data": "title",
			"name": "Title",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Title",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "budget",
			"name": "Budget",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Budget",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "daily_limit",
			"name": "DailyLimit",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "DailyLimit",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "cost_type",
			"name": "CostType",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": true,
			"title": "CostType",
			"type": "enum",
			"filter_valid_map": {
				"cpa": "CPA",
				"cpc": "CPC",
				"cpm": "CPM"
			}
		},
		{
			"data": "max_bid",
			"name": "MaxBid",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "MaxBid",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "avg_cpc",
			"name": "AvgCPC",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "AvgCPC",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "avg_cpm",
			"name": "AvgCPM",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "AvgCPM",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "ctr",
			"name": "Ctr",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Ctr",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_imp",
			"name": "TotalImp",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TotalImp",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_click",
			"name": "TotalClick",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TotalClick",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_conv",
			"name": "TotalConv",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TotalConv",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_cpc",
			"name": "TotalCpc",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TotalCpc",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_cpm",
			"name": "TotalCpm",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TotalCpm",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_spent",
			"name": "TotalSpent",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TotalSpent",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "today_imp",
			"name": "TodayImp",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TodayImp",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "today_click",
			"name": "TodayClick",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TodayClick",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "today_ctr",
			"name": "TodayCtr",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "TodayCtr",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "parent_email",
			"name": "ParentEmail",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "ParentEmail",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "owner_email",
			"name": "OwnerEmail",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "OwnerEmail",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "owner_id",
			"name": "OwnerID",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "OwnerID",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "domain_id",
			"name": "DomainID",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "DomainID",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "_actions",
			"name": "Actions",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "Actions",
			"type": "",
			"filter_valid_map": null
		}
	] `)
	assert.Nil(json.Unmarshal(Campaigntmp, &listCampaignDefinition))
}
