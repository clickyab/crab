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

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listDomains_data_tableResponse struct {
	Total   int64                  `json:"total"`
	Data    orm.DomainDetailsArray `json:"data"`
	Page    int                    `json:"page"`
	PerPage int                    `json:"per_page"`
	Hash    string                 `json:"hash"`
}

type listDomains_data_tableDefResponse struct {
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
	listDomains_data_tableDefinition permission.Columns
	Domains_data_tabletmp            = []byte{}
)

// @Route {
// 		url = /list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = list_domain:superGlobal
//		_sort_ = string, the sort and order like id:asc or id:desc available column "id","title","status","corporation_name","domain_base","owner_email","balance"
//		_status_ = string , filter the status field valid values are "enable","disable"
//		_title_ = string , search the title field
//		_domain_base_ = string , search the domain_base field
//		_owner_email_ = string , search the owner_email field
//		200 = listDomains_data_tableResponse
// }
func (u *Controller) listDomains_data_table(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("status"); e != "" && orm.DomainStatus(e).IsValid() {
		filter["status"] = e
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
		search["domain_base"] = e
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
	if !array.StringInArray(sort, "id", "title", "status", "corporation_name", "domain_base", "owner_email", "balance") {
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

	pc := permission.NewInterfaceComplete(usr, usr.ID, "list_domain", "superGlobal", domain.ID)
	dt, cnt, err := m.FillDomainDetails(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listDomains_data_tableResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Domains_data_tabletmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /list/definition
//		method = get
//		resource = list_domain:superGlobal
//		200 = listDomains_data_tableDefResponse
// }
func (u *Controller) defDomains_data_table(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Domains_data_tabletmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listDomains_data_tableDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, CheckLevel: false, PreventSelf: false, DateFilter: "", Hash: hash, Columns: listDomains_data_tableDefinition},
	)
}

func init() {
	Domains_data_tabletmp = []byte(` [
		{
			"data": "id",
			"name": "ID",
			"searchable": false,
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
			"data": "status",
			"name": "Status",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": true,
			"title": "Status",
			"type": "enum",
			"filter_valid_map": {
				"disable": "DisableDomainStatus",
				"enable": "EnableDomainStatus"
			}
		},
		{
			"data": "corporation_name",
			"name": "CorporationName",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "CorporationName",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "domain_base",
			"name": "DomainBase",
			"searchable": true,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "DomainBase",
			"type": "string",
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
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "balance",
			"name": "Balance",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Balance",
			"type": "number",
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
	assert.Nil(json.Unmarshal(Domains_data_tabletmp, &listDomains_data_tableDefinition))
}