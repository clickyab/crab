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
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listInventoryResponse struct {
	Total   int64                       `json:"total"`
	Data    orm.InventoryDataTableArray `json:"data"`
	Page    int                         `json:"page"`
	PerPage int                         `json:"per_page"`
	Hash    string                      `json:"hash"`
}

type listInventoryDefResponse struct {
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
	listInventoryDefinition permission.Columns
	Inventorytmp            = []byte{}
)

// @Route {
// 		url = /inventory/list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = list_inventory:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "id","created_at"
//		_status_ = string , filter the status field valid values are "enable","disable"
//		_label_ = string , search the label field
//		200 = listInventoryResponse
// }
func (u *Controller) listInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("status"); e != "" && orm.InventoryStatus(e).IsValid() {
		filter["i.status"] = e
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

	if e := r.URL.Query().Get("q"); e != "" {
		search["i.label"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "id", "created_at") {
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

	pc := permission.NewInterfaceComplete(usr, usr.ID, "list_inventory", "self", domain.ID)
	dt, cnt, err := m.FillInventoryDataTableArray(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listInventoryResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Inventorytmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /inventory/list/definition
//		method = get
//		resource = list_inventory:self
//		200 = listInventoryDefResponse
// }
func (u *Controller) defInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Inventorytmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listInventoryDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, CheckLevel: false, PreventSelf: false, DateFilter: "created_at", Hash: hash, Columns: listInventoryDefinition},
	)
}

func init() {
	Inventorytmp = []byte(` [
		{
			"data": "attached",
			"name": "AttachedCampaigns",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "AttachedCampaigns",
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
		},
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
			"data": "updated_at",
			"name": "UpdatedAt",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "UpdatedAt",
			"type": "date",
			"filter_valid_map": null
		},
		{
			"data": "user_id",
			"name": "UserID",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "UserID",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "domain_id",
			"name": "DomainID",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "DomainID",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "label",
			"name": "Label",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Label",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "status",
			"name": "Status",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": true,
			"title": "Status",
			"type": "enum",
			"filter_valid_map": {
				"disable": "DisableInventoryStatus",
				"enable": "EnableInventoryStatus"
			}
		}
	] `)
	assert.Nil(json.Unmarshal(Inventorytmp, &listInventoryDefinition))
}
