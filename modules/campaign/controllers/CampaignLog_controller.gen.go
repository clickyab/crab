// Code generated build with datatable DO NOT EDIT.

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listCampaignlogResponse struct {
	Total   int64                `json:"total"`
	Data    orm.CampaignLogArray `json:"data"`
	Page    int                  `json:"page"`
	PerPage int                  `json:"per_page"`
	Hash    string               `json:"hash"`
}

type listCampaignlogDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	CheckLevel  bool               `json:"checklevel"`
	PreventSelf bool               `json:"preventself"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listCampaignlogDefinition permission.Columns
	Campaignlogtmp            = []byte{}
)

// @Route {
// 		url = /log/:id
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = log_campaign:self
//		200 = listCampaignlogResponse
// }
func (u *Controller) listCampaignlog(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	//add date filter
	var from, to string
	if e := r.URL.Query().Get("from"); e != "" {
		//validate param
		fromTime, err := time.Parse(time.RFC3339, e)
		if err != nil {
			u.JSON(w, http.StatusBadRequest, err)
			return
		}
		from = "created_at" + "*" + fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	if e := r.URL.Query().Get("to"); e != "" {
		//validate param
		toTime, err := time.Parse(time.RFC3339, e)
		if err != nil {
			u.JSON(w, http.StatusBadRequest, err)
			return
		}
		to = "created_at" + "*" + toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	search := make(map[string]string)

	sort := ""
	order := "ASC"

	params := make(map[string]string)
	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}

	pc := permission.NewInterfaceComplete(usr, usr.ID, "log_campaign", "self", domain.ID)
	dt, cnt, err := m.FillCampaignLog(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listCampaignlogResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Campaignlogtmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /log/:id/definition
//		method = get
//		resource = log_campaign:self
//		200 = listCampaignlogDefResponse
// }
func (u *Controller) defCampaignlog(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Campaignlogtmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listCampaignlogDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, CheckLevel: false, PreventSelf: false, DateFilter: "created_at", Hash: hash, Columns: listCampaignlogDefinition},
	)
}

func init() {
	Campaignlogtmp = []byte(` [
		{
			"data": "created_at",
			"name": "CreatedAt",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "CreatedAt",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "action",
			"name": "Action",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Action",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "impersonator_email",
			"name": "ImpersonatorEmail",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "ImpersonatorEmail",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "manipulator_email",
			"name": "ManipulatorEmail",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "ManipulatorEmail",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "owner_email",
			"name": "OwnerEmail",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "OwnerEmail",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "data",
			"name": "Data",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "Data",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "campaign_name",
			"name": "CampaignName",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "CampaignName",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "kind",
			"name": "Kind",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Kind",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "strategy",
			"name": "Strategy",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Strategy",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "start_at",
			"name": "StartAt",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "StartAt",
			"type": "",
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
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "total_budget",
			"name": "TotalBudget",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "TotalBudget",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "daily_budget",
			"name": "DailyBudget",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "DailyBudget",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "max_bid",
			"name": "MaxBid",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "MaxBid",
			"type": "",
			"filter_valid_map": null
		}
	] `)
	assert.Nil(json.Unmarshal(Campaignlogtmp, &listCampaignlogDefinition))
}
