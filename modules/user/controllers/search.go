package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
)

type searchUserPayload struct {
	Email string `json:"email" validate:"required"`
}

type userSearchResp []aaa.UserSearchResult

// searchByMail search user by email
// @Rest {
// 		url = /search/mail
//		protected = true
// 		method = post
// }
func (c *Controller) searchByMail(ctx context.Context, r *http.Request, p *searchUserPayload) (userSearchResp, error) {
	dm := domain.MustGetDomain(ctx)
	foundUsers := aaa.NewAaaManager().ListUserByEmail(p.Email, dm.ID)
	if len(foundUsers) == 0 {
		return nil, errors.NotFoundWithEmail(p.Email)
	}
	return userSearchResp(foundUsers), nil
}
