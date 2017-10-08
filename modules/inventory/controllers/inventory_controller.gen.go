// Code generated build with datatable DO NOT EDIT.

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	Columns     permission.Columns `json:"columns"`
}

var (
	listInventoryDefinition permission.Columns
	tmp                     = []byte{}
)

// @Route {
// 		url = /list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		resource = inventory_list:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "created_at","publisher"
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_name_ = string , search the name field
//		_domain_ = string , search the domain field
//		200 = listInventoryResponse
// }
func (u *Controller) listInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" && orm.PublisherType(e).IsValid() {
		filter["inventories.kind"] = e
	}

	search := make(map[string]string)

	if e := r.URL.Query().Get("name"); e != "" {
		search["inventories.name"] = e
	}

	if e := r.URL.Query().Get("domain"); e != "" {
		search["inventories.domain"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "created_at", "publisher") {
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

	pc := permission.NewInterfaceComplete(usr, usr.ID, "inventory_list", "self", domain.ID)
	dt, cnt := m.FillInventoryDataTableArray(pc, filter, search, params, sort, order, p, c)
	res := listInventoryResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(tmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /list/definition
//		method = get
//		resource = inventory_list:self
//		200 = listInventoryDefResponse
// }
func (u *Controller) defInventory(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(tmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listInventoryDefResponse{Checkable: true, Multiselect: true, Hash: hash, Columns: listInventoryDefinition},
	)
}

func init() {
	tmp = []byte(` [
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
			"data": "name",
			"name": "Name",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Name",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "domain",
			"name": "Domain",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Domain",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "cat",
			"name": "Cat",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Cat",
			"type": "array",
			"filter_valid_map": null
		},
		{
			"data": "publisher",
			"name": "Publisher",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Publisher",
			"type": "string",
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
				"app": "PublisherTypeAPP",
				"web": "PublisherTypeWeb"
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
			"type": "enum",
			"filter_valid_map": null
		}
	] `)
	assert.Nil(json.Unmarshal(tmp, &listInventoryDefinition))
}
