package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
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
	currentUser, err := aaa.NewAaaManager().FindUserByEmailDomain(p.Email, uDomain)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(p.Password)) != nil {
		return nil, errors.InvalidEmailPassError
	}

	if currentUser.Status == aaa.RegisteredUserStatus {
		return nil, errors.UserNotVerifiedError
	}

	if currentUser.Status == aaa.BlockedUserStatus {
		return nil, errors.UserBlockedError
	}

	return c.createLoginResponse(currentUser), nil
}
