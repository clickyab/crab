// Code generated build with datatable DO NOT EDIT.

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listCampaignsResponse struct {
	Total   int64                    `json:"total"`
	Data    orm.CampaignDetailsArray `json:"data"`
	Page    int                      `json:"page"`
	PerPage int                      `json:"per_page"`
	Hash    string                   `json:"hash"`
}

type listCampaignsDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listCampaignsDefinition permission.Columns
	Campaignstmp            = []byte{}
)

// @Route {
// 		url = /list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = campaign_list:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "title","status","kind","total_imp","total_click","ectr","ecpc","ecpm","total_spend","max_bid","conversion","total_budget","today_spend","created_at","start_at","end_at","today_ctr","today_imp","today_click","creative","owner_email","conversion_rate","cpa","strategy","exchange"
//		_status_ = string , filter the status field valid values are "start","pause"
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_strategy_ = string , filter the strategy field valid values are "cpm","cpc","cpa"
//		_exchange_ = string , filter the exchange field valid values are "clickyab","all_except_clickyab","all"
//		_title_ = string , search the title field
//		_owner_email_ = string , search the owner_email field
//		200 = listCampaignsResponse
// }
func (u *Controller) listCampaigns(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("status"); e != "" && orm.Status(e).IsValid() {
		filter["status"] = e
	}

	if e := r.URL.Query().Get("kind"); e != "" && orm.CampaignKind(e).IsValid() {
		filter["kind"] = e
	}

	if e := r.URL.Query().Get("strategy"); e != "" && orm.Strategy(e).IsValid() {
		filter["strategy"] = e
	}

	if e := r.URL.Query().Get("exchange"); e != "" && orm.ExchangeType(e).IsValid() {
		filter["exchange"] = e
	}

	//add date filter
	var from, to string
	if e := r.URL.Query().Get("from"); e != "" {
		//validate param
		fromTime, err := time.Parse(time.RFC3339, e)
		if err != nil {
			u.JSON(w, http.StatusBadRequest, err)
			return
		}
		from = "" + "*" + fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	if e := r.URL.Query().Get("to"); e != "" {
		//validate param
		toTime, err := time.Parse(time.RFC3339, e)
		if err != nil {
			u.JSON(w, http.StatusBadRequest, err)
			return
		}
		to = "" + "*" + toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	search := make(map[string]string)

	if e := r.URL.Query().Get("q"); e != "" {
		search["title"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["u.email"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "title", "status", "kind", "total_imp", "total_click", "ectr", "ecpc", "ecpm", "total_spend", "max_bid", "conversion", "total_budget", "today_spend", "created_at", "start_at", "end_at", "today_ctr", "today_imp", "today_click", "creative", "owner_email", "conversion_rate", "cpa", "strategy", "exchange") {
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
	dt, cnt, err := m.FillCampaigns(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listCampaignsResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Campaignstmp)
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
//		200 = listCampaignsDefResponse
// }
func (u *Controller) defCampaigns(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Campaignstmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listCampaignsDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, DateFilter: "", Hash: hash, Columns: listCampaignsDefinition},
	)
}

func init() {
	Campaignstmp = []byte(` [
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
			"data": "title",
			"name": "Title",
			"searchable": true,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Title",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "status",
			"name": "Status",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": true,
			"title": "Status",
			"type": "enum",
			"filter_valid_map": {
				"pause": "PauseStatus",
				"start": "StartStatus"
			}
		},
		{
			"data": "kind",
			"name": "Kind",
			"searchable": false,
			"sortable": true,
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
			"data": "total_imp",
			"name": "TotalImp",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TotalImp",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "total_click",
			"name": "TotalClick",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TotalClick",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "ectr",
			"name": "ECTR",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "ECTR",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "ecpc",
			"name": "ECPC",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "ECPC",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "ecpm",
			"name": "ECPM",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "ECPM",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "total_spend",
			"name": "TotalSpend",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TotalSpend",
			"type": "number",
			"filter_valid_map": null
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
			"data": "conversion",
			"name": "Conversion",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Conversion",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "total_budget",
			"name": "TotalBudget",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TotalBudget",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "today_spend",
			"name": "TodaySpend",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TodaySpend",
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
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "EndAt",
			"type": "date",
			"filter_valid_map": null
		},
		{
			"data": "today_ctr",
			"name": "TodayCTR",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TodayCTR",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "today_imp",
			"name": "TodayImp",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TodayImp",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "today_click",
			"name": "TodayClick",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "TodayClick",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "creative",
			"name": "Creative",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Creative",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "owner_email",
			"name": "OwnerEmail",
			"searchable": true,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "OwnerEmail",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "conversion_rate",
			"name": "ConversionRate",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "ConversionRate",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "cpa",
			"name": "CPA",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "CPA",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "strategy",
			"name": "Strategy",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": true,
			"title": "Strategy",
			"type": "enum",
			"filter_valid_map": {
				"cpa": "CPA",
				"cpc": "CPC",
				"cpm": "CPM"
			}
		},
		{
			"data": "exchange",
			"name": "Exchange",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": true,
			"title": "Exchange",
			"type": "enum",
			"filter_valid_map": {
				"all": "All",
				"all_except_clickyab": "AllExceptClickyab",
				"clickyab": "Clickyab"
			}
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
	assert.Nil(json.Unmarshal(Campaignstmp, &listCampaignsDefinition))
}
