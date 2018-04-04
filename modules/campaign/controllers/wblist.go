package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type whiteBlackPayload struct {
	ListID   int64            `json:"list_id"`
	WhiteTyp *bool            `json:"white_typ" validate:"required"` //true for white false for black
	Exchange orm.ExchangeType `json:"exchange" validate:"required"`
}

// updateWhiteBlackList will update campaign white/black list
// @Rest {
// 		url = /wb/:id
//		protected = true
// 		method = put
// }
func (c *Controller) updateWhiteBlackList(ctx context.Context, r *http.Request, p *whiteBlackPayload) (*orm.Campaign, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}

	if err != nil || id < 1 {
		return nil, errors.InvalidIDErr
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		return o, err
	}

	err = db.UpdateCampaignWhiteBlackList(p.ListID, p.Exchange, p.WhiteTyp, o)
	if err != errors.ErrInventoryID.(error) {
		xlog.GetWithError(ctx, err).Debug("can not update white/black list")
	}

	return o, t9e.G("can't update campaign white/black list")
}

// deleteWhiteBlackList will update campaign white/black list
// @Rest {
// 		url = /wblist/:id
//		protected = true
// 		method = delete
// }
func (c *Controller) deleteWhiteBlackList(ctx context.Context, r *http.Request) (*orm.Campaign, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, err
	}
	db := orm.NewOrmManager()
	o, err := db.FindCampaignByID(id)
	if err != nil {
		return o, err
	}

	err = db.DeleteCampaignWhiteBlackList(o)
	if err != errors.ErrInventoryID.(error) {
		xlog.GetWithError(ctx, err).Debug("can not delete white/black list")
	}

	return o, t9e.G("can't delete campaign white/black list")
}
