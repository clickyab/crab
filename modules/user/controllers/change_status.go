package user

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	orm2 "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type changeUserStatus struct {
	Status        aaa.UserValidStatus `json:"status" validate:"required"`
	targetUser    *aaa.User           `json:"-"`
	currentDomain *orm2.Domain        `json:"-"`
}

// ChangeUserStatusResult change user status result
type ChangeUserStatusResult struct {
	UserID    int64               `json:"user_id"`
	NewStatus aaa.UserValidStatus `json:"new_status"`
}

func (p *changeUserStatus) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidUserStatusErr
	}
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm

	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	user, userErr := aaa.NewAaaManager().FindUserWithParentsByID(idInt, p.currentDomain.ID)
	if userErr != nil {
		return userErr
	}
	p.targetUser = user
	return nil
}

// changeUserStatus to user
// @Rest {
// 		url = /change-user-status/:id
//		protected = true
// 		method = patch
// 		resource = change_user_status:global
// }
func (c *Controller) changeUserStatus(ctx context.Context, r *http.Request, p *changeUserStatus) (*ChangeUserStatusResult, error) {
	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	//check permission
	_, ok := aaa.CheckPermOn(p.targetUser, currentUser, "change_user_status", p.currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// change user status
	p.targetUser.Status = p.Status
	err := userManager.UpdateUser(p.targetUser)
	if err != nil {
		xlog.GetWithError(ctx, err).Errorf("update user status error from db: userID:%d, status:%s", p.targetUser.ID, p.Status)
		return nil, errors.UpdateUserStatusErr
	}
	res := &ChangeUserStatusResult{
		UserID:    p.targetUser.ID,
		NewStatus: p.Status,
	}
	return res, nil
}
