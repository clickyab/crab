package user

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// getUserDetail get user detail  by id
// @Rest {
// 		url = /get/:id
//		protected = true
// 		method = get
//		resource = get_detail_user:global
// }
func (c *Controller) getUserDetail(ctx context.Context, r *http.Request) (*userResponse, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	// get user id from url params
	userID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	m := aaa.NewAaaManager()
	// find user
	userObj, err := m.FindUserByIDDomain(userID, userDomain.ID)
	if err != nil {
		return nil, errors.NotFoundWithDomainError(userDomain.DomainBase)
	}
	// check access
	_, ok := currentUser.HasOn("get_detail_user", userObj.ID, userDomain.ID, true, true)
	if !ok {
		return nil, errors.AccessDenied
	}

	// find user managers
	managers, err := m.FindUserManagers(userObj.ID, userDomain.ID)
	if err != nil {
		return nil, errors.GetUserManagerDbErr
	}
	// load user roles into its model for current domain
	userObj.SetUserRole(userDomain.ID)
	userRes := c.createUserResponse(userObj, nil, managers)
	return &userRes, nil
}
