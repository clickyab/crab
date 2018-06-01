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
	Total   int64                  `json:"total"`
	Data    orm.CampaignDailyArray `json:"data"`
	Page    int                    `json:"page"`
	PerPage int                    `json:"per_page"`
	Hash    string                 `json:"hash"`
}

type listCampaigndailyDefResponse struct {
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
	listCampaigndailyDefinition permission.Columns
	Campaigndailytmp            = []byte{}
)

// @Route {
// 		url = /daily/:id
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = daily_campaign:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "date","impression","click","ectr","ecpc","ecpm","spend","conversion","conversion_rate","cpa"
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
		from = "daily_id" + "*" + fromTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	if e := r.URL.Query().Get("to"); e != "" {
		//validate param
		toTime, err := time.Parse(time.RFC3339, e)
		if err != nil {
			u.JSON(w, http.StatusBadRequest, err)
			return
		}
		to = "daily_id" + "*" + toTime.Truncate(time.Hour*24).Format("2006-01-02 00:00:00")
	}

	search := make(map[string]string)

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "date", "impression", "click", "ectr", "ecpc", "ecpm", "spend", "conversion", "conversion_rate", "cpa") {
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

	pc := permission.NewInterfaceComplete(usr, usr.ID, "daily_campaign", "self", domain.ID)
	dt, cnt, err := m.FillDaily(pc, filter, from, to, search, params, sort, order, p, c)
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
//		resource = daily_campaign:self
//		200 = listCampaigndailyDefResponse
// }
func (u *Controller) defCampaigndaily(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Campaigndailytmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listCampaigndailyDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, CheckLevel: false, PreventSelf: false, DateFilter: "daily_id", Hash: hash, Columns: listCampaigndailyDefinition},
	)
}

func init() {
	Campaigndailytmp = []byte(` [
		{
			"data": "date",
			"name": "Date",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Date",
			"type": "date",
			"filter_valid_map": null
		},
		{
			"data": "impression",
			"name": "Impression",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Impression",
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
			"data": "spend",
			"name": "Spend",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Spend",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "conversion",
			"name": "Conversion",
			"searchable": false,
			"sortable": true,
			"visible": false,
			"filter": false,
			"title": "Conversion",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "conversion_rate",
			"name": "ConversionRate",
			"searchable": false,
			"sortable": true,
			"visible": false,
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
			"visible": false,
			"filter": false,
			"title": "CPA",
			"type": "number",
			"filter_valid_map": null
		}
	] `)
	assert.Nil(json.Unmarshal(Campaigndailytmp, &listCampaigndailyDefinition))
}
