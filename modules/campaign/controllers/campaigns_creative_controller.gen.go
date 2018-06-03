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

type listCampaigns_creativeResponse struct {
	Total   int64                           `json:"total"`
	Data    orm.CampaignCreativeStatusArray `json:"data"`
	Page    int                             `json:"page"`
	PerPage int                             `json:"per_page"`
	Hash    string                          `json:"hash"`
}

type listCampaigns_creativeDefResponse struct {
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
	listCampaigns_creativeDefinition permission.Columns
	Campaigns_creativetmp            = []byte{}
)

// @Route {
// 		url = /status-list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = list_campaign:superGlobal
//		_sort_ = string, the sort and order like id:asc or id:desc available column "id","title","kind","creative_count","owner_email","owner_mobile","created_at"
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_id_ = string , search the id field
//		_title_ = string , search the title field
//		200 = listCampaigns_creativeResponse
// }
func (u *Controller) listCampaigns_creative(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" && orm.CampaignKind(e).IsValid() {
		filter["kind"] = e
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
		search["c.id"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["title"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "id", "title", "kind", "creative_count", "owner_email", "owner_mobile", "created_at") {
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

	pc := permission.NewInterfaceComplete(usr, usr.ID, "list_campaign", "superGlobal", domain.ID)
	dt, cnt, err := m.FillCampaignCreativeStatus(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listCampaigns_creativeResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Campaigns_creativetmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /status-list/definition
//		method = get
//		resource = list_campaign:superGlobal
//		200 = listCampaigns_creativeDefResponse
// }
func (u *Controller) defCampaigns_creative(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Campaigns_creativetmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listCampaigns_creativeDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, CheckLevel: false, PreventSelf: true, DateFilter: "", Hash: hash, Columns: listCampaigns_creativeDefinition},
	)
}

func init() {
	Campaigns_creativetmp = []byte(` [
		{
			"data": "id",
			"name": "ID",
			"searchable": true,
			"sortable": true,
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
			"data": "creative_count",
			"name": "CreativeCount",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "CreativeCount",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "owner_email",
			"name": "OwnerEmail",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "OwnerEmail",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "owner_mobile",
			"name": "OwnerMobile",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "OwnerMobile",
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
	assert.Nil(json.Unmarshal(Campaigns_creativetmp, &listCampaigns_creativeDefinition))
}
