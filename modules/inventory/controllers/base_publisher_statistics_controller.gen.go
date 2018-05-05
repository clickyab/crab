// Code generated build with datatable DO NOT EDIT.

package controllers

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listBase_publisher_statisticsResponse struct {
	Total   int64                             `json:"total"`
	Data    orm.PublishersBaseStatisticsArray `json:"data"`
	Page    int                               `json:"page"`
	PerPage int                               `json:"per_page"`
	Hash    string                            `json:"hash"`
}

type listBase_publisher_statisticsDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	QueryParams string             `json:"queryparams"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listBase_publisher_statisticsDefinition permission.Columns
	Base_publisher_statisticstmp            = []byte{}
)

// @Route {
// 		url = /base-publishers/statistics
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = publisher_base_statistics:self
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_status_ = string , filter the status field valid values are "accepted","pending","blocked"
//		_name_ = string , search the name field
//		_domain_ = string , search the domain field
//		_supplier_ = string , search the supplier field
//		200 = listBase_publisher_statisticsResponse
// }
func (u *Controller) listBase_publisher_statistics(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" && orm.PublisherType(e).IsValid() {
		filter["pub.kind"] = e
	}

	if e := r.URL.Query().Get("status"); e != "" && orm.Status(e).IsValid() {
		filter["pub.status"] = e
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
		search["pub.name"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["pub.domain"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["pub.supplier"] = e
	}

	sort := ""
	order := "ASC"

	params := make(map[string]string)
	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}

	queryParams := make(map[string]string)

	pc := permission.NewInterfaceComplete(usr, usr.ID, "publisher_base_statistics", "self", domain.ID)
	dt, cnt, err := m.FillPublishersBaseStatistics(pc, filter, from, to, search, params, queryParams, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listBase_publisher_statisticsResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Base_publisher_statisticstmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /base-publishers/statistics/definition
//		method = get
//		resource = publisher_base_statistics:self
//		200 = listBase_publisher_statisticsDefResponse
// }
func (u *Controller) defBase_publisher_statistics(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Base_publisher_statisticstmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listBase_publisher_statisticsDefResponse{Checkable: false, SearchKey: "q", QueryParams: "", Multiselect: false, DateFilter: "created_at", Hash: hash, Columns: listBase_publisher_statisticsDefinition},
	)
}

func init() {
	Base_publisher_statisticstmp = []byte(` [
		{
			"data": "count",
			"name": "Count",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Count",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "avg_imp",
			"name": "AvgImp",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "AvgImp",
			"type": "",
			"filter_valid_map": null
		},
		{
			"data": "exchange_count",
			"name": "ExchangeCount",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "ExchangeCount",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "name",
			"name": "Name",
			"searchable": true,
			"sortable": false,
			"visible": false,
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
			"visible": false,
			"filter": false,
			"title": "Domain",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "supplier",
			"name": "Supplier",
			"searchable": true,
			"sortable": false,
			"visible": false,
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
			"visible": false,
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
			"visible": false,
			"filter": true,
			"title": "Status",
			"type": "enum",
			"filter_valid_map": {
				"accepted": "ActiveStatus",
				"blocked": "BlockedStatus",
				"pending": "PendingStatus"
			}
		}
	] `)
	assert.Nil(json.Unmarshal(Base_publisher_statisticstmp, &listBase_publisher_statisticsDefinition))
}
