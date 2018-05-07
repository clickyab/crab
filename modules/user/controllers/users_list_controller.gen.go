// Code generated build with datatable DO NOT EDIT.

package user

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/array"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

type listUsers_listResponse struct {
	Total   int64             `json:"total"`
	Data    aaa.UserListArray `json:"data"`
	Page    int               `json:"page"`
	PerPage int               `json:"per_page"`
	Hash    string            `json:"hash"`
}

type listUsers_listDefResponse struct {
	Hash        string             `json:"hash"`
	Checkable   bool               `json:"checkable"`
	Multiselect bool               `json:"multiselect"`
	DateFilter  string             `json:"datefilter"`
	SearchKey   string             `json:"searchkey"`
	Columns     permission.Columns `json:"columns"`
}

var (
	listUsers_listDefinition permission.Columns
	Users_listtmp            = []byte{}
)

// @Route {
// 		url = /list
//		method = get
//		_c_ = int , count per page
//		_p_ = int , page number
//		_q_ = string , parameter for search
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		resource = user_list:global
//		_sort_ = string, the sort and order like id:asc or id:desc available column "balance","created_at"
//		_status_ = string , filter the status field valid values are "registered","blocked","active"
//		_full_name_ = string , search the full_name field
//		_email_ = string , search the email field
//		_cellphone_ = string , search the cellphone field
//		_land_line_ = string , search the land_line field
//		_ssn_ = string , search the ssn field
//		200 = listUsers_listResponse
// }
func (u *Controller) listUsers_list(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := aaa.NewAaaManager()
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	p, c := framework.GetPageAndCount(r, false)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("status"); e != "" && aaa.UserValidStatus(e).IsValid() {
		filter["users.status"] = e
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
		search["users.full_name"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["users.email"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["users.cellphone"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["users.land_line"] = e
	}

	if e := r.URL.Query().Get("q"); e != "" {
		search["users.ssn"] = e
	}

	s := r.URL.Query().Get("sort")
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		parts = append(parts, "asc")
	}
	sort := parts[0]
	if !array.StringInArray(sort, "balance", "created_at") {
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

	pc := permission.NewInterfaceComplete(usr, usr.ID, "user_list", "global", domain.ID)
	dt, cnt, err := m.FillUsers(pc, filter, from, to, search, params, sort, order, p, c)
	if err != nil {
		u.JSON(w, http.StatusBadRequest, err)
		return
	}
	res := listUsers_listResponse{
		Total:   cnt,
		Data:    dt.Filter(usr),
		Page:    p,
		PerPage: c,
	}

	h := sha1.New()
	_, _ = h.Write(Users_listtmp)
	res.Hash = fmt.Sprintf("%x", h.Sum(nil))

	u.OKResponse(
		w,
		res,
	)
}

// @Route {
// 		url = /list/definition
//		method = get
//		resource = user_list:global
//		200 = listUsers_listDefResponse
// }
func (u *Controller) defUsers_list(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := sha1.New()
	_, _ = h.Write(Users_listtmp)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	u.OKResponse(
		w,
		listUsers_listDefResponse{Checkable: false, SearchKey: "q", Multiselect: false, DateFilter: "created_at", Hash: hash, Columns: listUsers_listDefinition},
	)
}

func init() {
	Users_listtmp = []byte(` [
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
			"data": "full_name",
			"name": "FullName",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "FullName",
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
				"active": "ActiveUserStatus",
				"blocked": "BlockedUserStatus",
				"registered": "RegisteredUserStatus"
			}
		},
		{
			"data": "balance",
			"name": "Balance",
			"searchable": false,
			"sortable": true,
			"visible": true,
			"filter": false,
			"title": "Balance",
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
			"data": "cellphone",
			"name": "CellPhone",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "CellPhone",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "land_line",
			"name": "LandLine",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "LandLine",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "account_type",
			"name": "AccountType",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "AccountType",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "ssn",
			"name": "SSN",
			"searchable": true,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "SSN",
			"type": "string",
			"filter_valid_map": null
		},
		{
			"data": "avatar",
			"name": "Avatar",
			"searchable": false,
			"sortable": false,
			"visible": true,
			"filter": false,
			"title": "Avatar",
			"type": "string",
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
			"type": "string",
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
	assert.Nil(json.Unmarshal(Users_listtmp, &listUsers_listDefinition))
}
