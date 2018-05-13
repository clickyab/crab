package user

import (
	"context"
	"net/http"

	domain2 "clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
)

// ping is for ping
// @Rest {
// 		url = /ping
//		protected = true
// 		method = get
// }
func (c *Controller) ping(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	user := authz.MustGetUser(ctx)
	domain := domain2.MustGetDomain(ctx)
	userPerms, err := user.GetAllUserPerms(domain.ID)
	if err != nil {
		return nil, err
	}
	return &ResponseLoginOK{
		Token:   authz.MustGetToken(ctx),
		Account: c.createUserResponse(user),
		Perms:   userPerms,
	}, nil
}
