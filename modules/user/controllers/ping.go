package user

import (
	"context"
	"net/http"

	domain2 "clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/xlog"
)

// ping is for ping
// @Rest {
// 		url = /ping
//		protected = true
// 		method = get
// }
func (c *Controller) ping(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	user := authz.MustGetUser(ctx)
	userToken := authz.MustGetToken(ctx)
	domain := domain2.MustGetDomain(ctx)
	userPerms, err := user.GetAllUserPerms(domain.ID)
	impersonatorToken := aaa.ImpersonatorToken(userToken)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error on get user permissions")
		return nil, errors.GetUserPermsDbErr
	}
	res := &ResponseLoginOK{
		Token:   authz.MustGetToken(ctx),
		Account: c.createUserResponse(user, userPerms),
	}
	if impersonatorToken != "" {
		res.ImpersonatorToken = impersonatorToken
	}
	return res, nil
}
