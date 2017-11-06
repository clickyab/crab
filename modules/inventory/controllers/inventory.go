package controllers

import (
	"context"
	"net/http"

	"strconv"

	"time"

	"regexp"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/inventory/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
	"github.com/rs/xmux"
)

// Controller is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /inventory
// }
type Controller struct {
	controller.Base
}

type whiteBlackLists []orm.WhiteBlackList

// whiteBlackLists return all user inventories
// @Route {
// 		url = /presets
//		method = get
//		200 = whiteBlackLists
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (ctrl *Controller) whiteBlackLists(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := authz.MustGetUser(ctx)
	res := orm.NewOrmManager().ListWhiteBlackListsWithFilter(
		"user_id = ? ", u.ID)
	if len(res) == 0 {
		ctrl.NotFoundResponse(w, trans.E("User doesn't have any list"))
		return
	}
	ctrl.OKResponse(w, res)
}

// whiteBlackList return a user inventory
// @Route {
// 		url = /preset/:id
//		method = get
//		200 = orm.WhiteBlackList
//		404 = controller.ErrorResponseSimple
//		middleware = authz.Authenticate
// }
func (ctrl *Controller) whiteBlackList(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	id, e := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if e != nil {
		ctrl.BadResponse(w, trans.E("not valid id"))
		return
	}
	res, e := orm.NewOrmManager().FindWhiteBlackListByID(id)
	if e != nil {
		ctrl.NotFoundResponse(w, trans.E("Inventory with id %d does not exists!", id))
		return
	}
	ctrl.OKResponse(w, res)
}

//@Validate {
//}
type whiteBlackList struct {
	Label   string   `json:"label" db:"label" validate:"gt=7"`
	Domains []string `json:"domains" db:"domains" validate:"gt=0"`
	// Kind shows if it's a white list (true) or blacklist (false)
	Kind          bool              `json:"kind" db:"kind"`
	PublisherType orm.PublisherType `json:"publisher_type" db:"publisher_type"`
}

// addPreset get a new whitelist blacklist for user
// @Route {
// 		url = /preset
//		method = post
//		200 = orm.WhiteBlackList
//		middleware = authz.Authenticate
//		payload = whiteBlackList
// }
func (ctrl *Controller) addPreset(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*whiteBlackList)
	u := authz.MustGetUser(ctx)
	dm := domain.MustGetDomain(ctx)
	now := time.Now()
	resMap := FillResMap(pl.Domains)
	d := &orm.WhiteBlackList{
		Active:        true,
		UpdatedAt:     now,
		CreatedAt:     now,
		Domains:       mysql.StringMapJSONArray(resMap),
		Label:         pl.Label,
		Kind:          pl.Kind,
		PublisherType: pl.PublisherType,
		UserID:        u.ID,
		DomainID:      dm.ID,
	}
	e := orm.NewOrmManager().CreateWhiteBlackList(d)
	assert.Nil(e)
	ctrl.OKResponse(w, d)
}

// FillResMap fill string map json for black white
func FillResMap(domains []string) map[string][]string {
	var pattern = regexp.MustCompile(`^\d+$`)
	var resMap = make(map[string][]string)
	var ids = []int64{}
	for i := range domains {
		if pattern.Match([]byte(domains[i])) {
			idInt, err := strconv.ParseInt(domains[i], 10, 64)
			assert.Nil(err)
			ids = append(ids, idInt)
		} else {
			resMap[domains[i]] = []string{}
		}
	}
	m := orm.NewOrmManager()
	if len(ids) > 0 {
		for _, v := range m.GetDomainPublishers(ids) {
			if _, k := resMap[v.Publisher]; k {
				resMap[v.Publisher] = append(resMap[v.Publisher], v.Domain)
			} else {
				resMap[v.Publisher] = []string{v.Domain}
			}
		}
	}
	return resMap
}
