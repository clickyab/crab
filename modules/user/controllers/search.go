package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
)

type searchUserPayload struct {
	Email string `json:"email" validate:"required"`
}

type userSearchResp []aaa.UserSearchResult

// searchByMail search user by email
// @Rest {
// 		url = /search/user/mail
//		protected = true
// 		method = post
// }
func (c *Controller) searchByMail(ctx context.Context, r *http.Request, p *searchUserPayload) (userSearchResp, error) {
	dm := domain.MustGetDomain(ctx)
	foundUsers := aaa.NewAaaManager().ListUserByEmail(p.Email, dm.ID)
	if len(foundUsers) == 0 {
		return userSearchResp{}, nil
	}
	return userSearchResp(foundUsers), nil
}

// searchMangerByMail search manager user by email
// @Rest {
// 		url = /search/managers/mail
//		protected = true
// 		method = post
// }
func (c *Controller) searchMangerByMail(ctx context.Context, r *http.Request, p *searchUserPayload) (userSearchResp, error) {
	dm := domain.MustGetDomain(ctx)
	foundUsers := aaa.NewAaaManager().SearchAccountByMailDomain(p.Email, dm.ID)
	if len(foundUsers) == 0 {
		return userSearchResp{}, nil
	}
	return userSearchResp(foundUsers), nil
}

// searchAdvertiserByMail search advertiser user by email
// @Rest {
// 		url = /search/advertiser/mail
//		protected = true
// 		method = post
// }
func (c *Controller) searchAdvertiserByMail(ctx context.Context, r *http.Request, p *searchUserPayload) (userSearchResp, error) {
	dm := domain.MustGetDomain(ctx)
	foundUsers := aaa.NewAaaManager().SearchAdvertiserByMailDomain(p.Email, dm.ID)
	if len(foundUsers) == 0 {
		return userSearchResp{}, nil
	}
	return userSearchResp(foundUsers), nil
}
