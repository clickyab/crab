package user

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

// getUserDetail get user detail  by id
// @Rest {
// 		url = /detail/:id
//		protected = true
// 		method = get
//		resource = get_detail_user:global
// }
func (c *Controller) getUserDetail(ctx context.Context, r *http.Request) (*userResponse, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	// check access
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "get_detail_user", userDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// get user id from url params
	id := xmux.Param(ctx, "id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	// find user
	userObj, err := aaa.NewAaaManager().FindUserByIDDomain(userID, userDomain.ID)
	if err != nil {
		return nil, errors.NotFoundWithDomainError(userDomain.Name)
	}
	res := c.createUserResponse(userObj, nil)
	return &res, nil
}
