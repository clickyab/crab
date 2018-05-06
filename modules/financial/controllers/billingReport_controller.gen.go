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
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listBillingreportResponse struct {
	Total   int64                     `json:"total"`
	Data    orm.BillingDataTableArray `json:"data"`
	Page    int                       `json:"page"`
	PerPage int                       `json:"per_page"`
	Hash    string                    `json:"hash"`
}

type listBillingreportDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	QueryParams string             `json:"queryparams"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listBillingreportDefinition permission.Columns
	Billingreporttmp            = []byte{}
)

// @Route {
// 		url = /billing
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = get_billing:self
//		_sort_ = string, the sort and order like id:asc or id:desc available column "amount","balance","created_at"
//		_pay_model_ = string , filter the pay_model field valid values are "online_payment","bank_snap","manual_cash_change"
//		_user_id_ = string , search the user_id field
//		_first_name_ = string , search the first_name field
//		_last_name_ = string , search the last_name field
//		_email_ = string , search the email field
//		200 = listBillingreportResponse
// }
func (u *Controller) listBillingreport(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := orm.NewOrmManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("pay_model"); e != "" && orm.PayModels(e).IsValid() {
		filter["pay_model"] = e
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
		search["user_id"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["first_name"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["last_name"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["email"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "amount", "balance", "created_at") {
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

	queryParams := make(map[string]string)

	pc := permission.NewInterfaceComplete(usr, usr.ID, "get_billing", "self", domain.ID)
	dt, cnt, err := m.FillBilling(pc, filter, from, to, search, params, queryParams, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listBillingreportResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Billingreporttmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /billing/definition
//		method = get
//		resource = get_billing:self
//		200 = listBillingreportDefResponse
// }
func (u *Controller) defBillingreport(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Billingreporttmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listBillingreportDefResponse{Checkable: false, SearchKey: "q", QueryParams: "", Multiselect: false, DateFilter: "created_at", Hash: hash, Columns: listBillingreportDefinition},
	)
}

func init() {
	Billingreporttmp = []byte(` [
		{
			"data": "id",
			"name": "ID",
			"searchable": false,
			"sortable": false,
			"visible": false,
			"filter": false,
			"title": "ID",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "user_id",
			"name": "UserID",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "UserID",
			"type": "number",
			"filter_valid_map": null
		},
		{
			"data": "pay_model",
			"name": "PayModel",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": true,
			"title": "PayModel",
			"type": "enum",
			"filter_valid_map": {
				"bank_snap": "BankSnapModel",
				"manual_cash_change": "ManualCashChangeModel",
				"online_payment": "OnlinePaymentModel"
			}
		},
		{
			"data": "first_name",
			"name": "FirstName",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "FirstName",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "last_name",
			"name": "LastName",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "LastName",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "email",
			"name": "Email",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Email",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "amount",
			"name": "Amount",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Amount",
			"type": "number",
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
	assert.Nil(json.Unmarshal(Billingreporttmp, &listBillingreportDefinition))
}
