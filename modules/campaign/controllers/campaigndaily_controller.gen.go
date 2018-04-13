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

type listCampaigndailyResponse struct {
	Total   int64                           `json:"total"`
	Data    orm.CampaignDailyDataTableArray `json:"data"`
	Page    int                             `json:"page"`
	PerPage int                             `json:"per_page"`
	Hash    string                          `json:"hash"`
}

type listCampaigndailyDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listCampaigndailyDefinition permission.Columns
	Campaigndailytmp            = []byte{}
)

// @Route {
// 		url = /daily/:id
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = campaign_list:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "created_at","imp","click","conv","spent","ctr"
//		200 = listCampaigndailyResponse
// }
func (u *Controller) listCampaigndaily(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
		from = "" + ":" + fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	if e := r.URL.Query().Get("to"); e != "" {
		//validate param
		toTime, err := time.Parse(time.RFC3339, e)
		if err != nil {
			u.JSON(w, http.StatusBadRequest, err)
			return
		}
		to = "" + ":" + toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	search := make(map[string]string)

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "created_at", "imp", "click", "conv", "spent", "ctr") {
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
	dt, cnt, err := m.FillCampaignDailyDataTableArray(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listCampaigndailyResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Campaigndailytmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /daily/:id/definition
//		method = get
//		resource = campaign_list:self
//		200 = listCampaigndailyDefResponse
// }
func (u *Controller) defCampaigndaily(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Campaigndailytmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listCampaigndailyDefResponse{Checkable: false, SearchKey: "", Multiselect: false, DateFilter: "", Hash: hash, Columns: listCampaigndailyDefinition},
	)
}

func init() {
	Campaigndailytmp = []byte(` [
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
			"data": "imp",
			"name": "Imp",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Imp",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "click",
			"name": "Click",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Click",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "conv",
			"name": "Conv",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Conv",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "cpm",
			"name": "Cpm",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Cpm",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "cpc",
			"name": "Cpc",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Cpc",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "spent",
			"name": "Spent",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Spent",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "cpa",
			"name": "Cpa",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Cpa",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "ctr",
			"name": "Ctr",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Ctr",
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
	assert.Nil(json.Unmarshal(Campaigndailytmp, &listCampaigndailyDefinition))
}
