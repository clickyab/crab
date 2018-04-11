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

type listInvpublisherResponse struct {
	Total   int64                             `json:"total"`
	Data    orm.SinglePublisherDataTableArray `json:"data"`
	Page    int                               `json:"page"`
	PerPage int                               `json:"per_page"`
	Hash    string                            `json:"hash"`
}

type listInvpublisherDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listInvpublisherDefinition permission.Columns
	Invpublishertmp            = []byte{}
)

// @Route {
// 		url = /publisher/list/single/:id
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = list_inventory:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "created_at"
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_status_ = string , filter the status field valid values are "accepted","pending","blocked"
//		_name_ = string , search the name field
//		_domain_ = string , search the domain field
//		_supplier_ = string , search the supplier field
//		200 = listInvpublisherResponse
// }
func (u *Controller) listInvpublisher(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)
	dateRange := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" && orm.PublisherType(e).IsValid() {
		filter["publishers.kind"] = e
	}

	if e := r.URL.Query().Get("status"); e != "" && orm.Status(e).IsValid() {
		filter["publishers.status"] = e
	}

	//add date filter
	if e := r.URL.Query().Get("from"); e != "" {
		dateRange["from-created_at"] = e
	}

	if e := r.URL.Query().Get("to"); e != "" {
		dateRange["to-created_at"] = e
	}

	search := make(map[string]string)

	if e := r.URL.Query().Get("q"); e != "" {
		search["publishers.name"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["publishers.domain"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["publishers.supplier"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "created_at") {
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
	dt, cnt, err := m.FillSinglePublisherDataTableArray(pc, filter, dateRange, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listInvpublisherResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Invpublishertmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /publisher/list/single/:id/definition
//		method = get
//		resource = list_inventory:self
//		200 = listInvpublisherDefResponse
// }
func (u *Controller) defInvpublisher(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Invpublishertmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listInvpublisherDefResponse{Checkable: true, SearchKey: "q", Multiselect: true, DateFilter: "created_at", Hash: hash, Columns: listInvpublisherDefinition},
	)
}

func init() {
	Invpublishertmp = []byte(` [
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
			"visible": false,
			"filter": false,
			"title": "DomainID",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "parent_ids",
			"name": "ParentIDs",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "ParentIDs",
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
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "ID",
			"type": "number",
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
			"data": "categories",
			"name": "Categories",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Categories",
			"type": "array",
			"filter_valid_map": null
		},
		{
			"data": "supplier",
			"name": "Supplier",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Supplier",
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
			"filter": true,
			"title": "Status",
			"type": "enum",
			"filter_valid_map": {
				"accepted": "ActiveStatus",
				"blocked": "BlockedStatus",
				"pending": "PendingStatus"
			}
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
			"data": "deleted_at",
			"name": "DeletedAt",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "DeletedAt",
			"type": "date",
			"filter_valid_map": null
		}
	] `)
	assert.Nil(json.Unmarshal(Invpublishertmp, &listInvpublisherDefinition))
}
