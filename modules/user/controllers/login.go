package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/gettext/t9e"
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
		return nil, t9e.G("wrong email or password")
	}

	if currentUser.Status == aaa.RegisteredUserStatus {
		return nil, t9e.G("not verified")
	}

	if currentUser.Status == aaa.BlockedUserStatus {
		return nil, t9e.G("this account has been blocked")
	}

	return c.createLoginResponse(currentUser), nil
}
