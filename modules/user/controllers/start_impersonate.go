package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
)

// @Validate {
// }
type startImpersonatePayload struct {
	TargetUserID  int64       `json:"user_id" validate:"required"`
	targetUser    *aaa.User   `json:"-"`
	currentDomain *orm.Domain `json:"-"`
}

func (p *startImpersonatePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm
	//find user
	m := aaa.NewAaaManager()
	targetUser, err := m.FindUserWithParentsByID(p.TargetUserID, p.currentDomain.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.targetUser = targetUser
	return nil
}

// startImpersonate start impersonate for user
// @Rest {
// 		url = /start-impersonate
//		protected = true
// 		method = post
// 		resource = impersonate_user:global
// }
func (c *Controller) startImpersonate(ctx context.Context, r *http.Request, p *startImpersonatePayload) (*ResponseLoginOK, error) {
	currentUser := authz.MustGetUser(ctx)
	//check permission
	_, ok := aaa.CheckPermOn(p.targetUser, currentUser, "impersonate_user", p.currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	//generate impersonate token for target user
	targetToken := aaa.GetImpersonateToken(p.targetUser, currentUser.AccessToken)

	userPerms, err := p.targetUser.GetAllUserPerms(p.currentDomain.ID)
	if err != nil {
		return nil, err
	}
	result := &ResponseLoginOK{
		Token:   targetToken,
		Account: c.createUserResponse(p.targetUser, userPerms),
	}
	return result, nil
}