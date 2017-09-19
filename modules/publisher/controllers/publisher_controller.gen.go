// Code generated build with datatable DO NOT EDIT.

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/publisher/pub"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listPublisherResponse struct {
	Total   int64                       `json:"total"`
	Data    pub.PublisherDataTableArray `json:"data"`
	Page    int                         `json:"page"`
	PerPage int                         `json:"per_page"`
	Hash    string                      `json:"hash"`
}

type listPublisherDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listPublisherDefinition permission.Columns
	tmp                     = []byte{}
)

// @Route {
// 		url = /list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		resource = pub_list:self
//		_pub_type_ = string , filter the pub_type field valid values are "app","web"
//		_status_ = string , filter the status field valid values are "accepted","pending","blocked"
//		_domain_ = string , search the domain field
//		200 = listPublisherResponse
// }
func (u *Controller) listPublisher(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := pub.NewPubManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("pub_type"); e != "" && pub.PType(e).IsValid() {
		filter["publishers.pub_type"] = e
	}

	if e := r.URL.Query().Get("status"); e != "" && pub.Status(e).IsValid() {
		filter["publishers.status"] = e
	}

	search := make(map[string]string)

	if e := r.URL.Query().Get("domain"); e != "" {
		search["publishers.domain"] = e
	}

	sort := ""
	order := "ASC"

	params := make(map[string]string)
	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}

	pc := permission.NewInterfaceComplete(usr, usr.ID, "pub_list", "self", domain.ID)
	dt, cnt := m.FillPublisherDataTableArray(pc, filter, search, params, sort, order, p, c)
	res := listPublisherResponse{
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
//		resource = pub_list:self
//		200 = listPublisherDefResponse
// }
func (u *Controller) defPublisher(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(tmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listPublisherDefResponse{Checkable: true, Multiselect: true, Hash: hash, Columns: listPublisherDefinition},
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
			"data": "name",
			"name": "Name",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Name",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "supplier",
			"name": "Supplier",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Supplier",
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
			"data": "pub_type",
			"name": "PublisherType",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": true,
			"title": "PublisherType",
			"type": "enum",
			"filter_valid_map": {
				"app": "AppPubType",
				"web": "WebPubType"
			}
		},
		{
			"data": "status",
			"name": "PubStatus",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": true,
			"title": "PubStatus",
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
			"sortable": false,
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
		}
	] `)
	assert.Nil(json.Unmarshal(tmp, &listPublisherDefinition))
}
