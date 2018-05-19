package user

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

// @Validate {
// }
type changePass struct {
	Password      string      `json:"password" validate:"gt=5"`
	targetUser    *aaa.User   `json:"-"`
	currentDomain *orm.Domain `json:"-"`
}

func (p *changePass) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}
	currentDomain := domain.MustGetDomain(ctx)
	p.currentDomain = currentDomain
	// find target user
	targetUser, err := aaa.NewAaaManager().FindUserWithParentsByID(id, currentDomain.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.targetUser = targetUser
	return nil
}

// changeAdminPassword change password (Admin)
// @Rest {
// 		url = /admin/password/change/:id
// 		method = patch
//		protected = true
//		resource = edit_user:global
// }
func (c *Controller) changeAdminPassword(ctx context.Context, r *http.Request, p *changePass) (*controller.NormalResponse, error) {
	currentUser := authz.MustGetUser(ctx)
	_, ok := aaa.CheckPermOn(p.targetUser, currentUser, "edit_user", p.currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// update password
	err := p.targetUser.ChangePassword(p.Password)
	if err != nil {
		switch err {
		case aaa.ErrorOldPass:
			return nil, err
		default:
			assert.Nil(err)
		}
	}
	return nil, nil
}
