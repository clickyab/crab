package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/xlog"
	"golang.org/x/crypto/bcrypt"
)

// @Validate {
// }
type loginPayload struct {
	Email    string `json:"email" validate:"required,email" error:"email is invalid"`
	Password string `json:"password" validate:"required,gt=5" error:"password is too short"`
}

// login userPayload in system
// @Rest {
// 		url = /login
//		method = post
// }
func (c *Controller) login(ctx context.Context, r *http.Request, p *loginPayload) (*ResponseLoginOK, error) {
	uDomain := domain.MustGetDomain(ctx)
	currentUser, err := aaa.NewAaaManager().FindUserByEmail(p.Email)
	if err != nil {
		return nil, errors.InvalidEmailPassError
	}
	if !currentUser.DomainLess {
		_, err := orm.NewOrmManager().FindActiveUserDomainByUserDomain(currentUser.ID, uDomain.ID)
		if err != nil {
			return nil, errors.AccessDenied
		}
	}
	if bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(p.Password)) != nil {
		return nil, errors.InvalidEmailPassError
	}

	if currentUser.Status == aaa.RegisteredUserStatus {
		return nil, errors.UserNotVerifiedError
	}

	if currentUser.Status == aaa.BlockedUserStatus {
		return nil, errors.UserBlockedError
	}

	userPerms, err := currentUser.GetAllUserPerms(uDomain.ID)
	if err != nil {
		xlog.Get(ctx).Error("get user permissions error:", err)
		return nil, t9e.G("get user permissions error")
	}

	token := aaa.GetNewToken(currentUser)
	return &ResponseLoginOK{
		Token:   token,
		Account: c.createUserResponse(currentUser, userPerms, nil),
	}, nil
}
