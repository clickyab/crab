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

// SearchUser resp
type SearchUser struct {
	Email string `json:"email"`
	ID    int64  `json:"id"`
}

// SearchUserArr resp
type SearchUserArr []SearchUser

// searchByMail search user by email
// @Rest {
// 		url = /search/mail
//		protected = true
// 		method = post
// }
func (c *Controller) searchByMail(ctx context.Context, r *http.Request, p *searchUserPayload) (SearchUserArr, error) {
	dm := domain.MustGetDomain(ctx)
	foundUsers := aaa.NewAaaManager().ListUserByEmail(p.Email, dm.ID)
	if len(foundUsers) == 0 {
		return nil, errors.NotFoundWithEmail(p.Email)
	}
	var res = make([]SearchUser, 0)
	for i := range foundUsers {
		res = append(res, SearchUser{
			Email: foundUsers[i].Email,
			ID:    foundUsers[i].ID,
		})
	}
	return SearchUserArr(res), nil
}
