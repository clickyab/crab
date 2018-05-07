package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/ad/orm"
)

type rejectReasons []orm.CreativeRejectReasons

// getCreativeRejectReasons to get list of creative reject reasons
// @Rest {
// 		url = /creative-reject-reasons
//		protected = true
// 		method = get
// 		resource = get_creative_reject_reason:self
// }
func (c Controller) getCreativeRejectReasons(ctx context.Context, r *http.Request) (rejectReasons, error) {
	m := orm.NewOrmManager()
	res := m.ListCreativeRejectReasons()
	return rejectReasons(res), nil
}
