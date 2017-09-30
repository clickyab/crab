package user

import (
	"context"
	"errors"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/models"
	"golang.org/x/crypto/bcrypt"
)

// @Validate {
// }
type loginPayload struct {
	Email    string `json:"email" validate:"required,email" error:"email is invalid"`
	Password string `json:"password" validate:"required,gt=5" error:"password is too short"`
}

// login userPayload in system
// @Route {
// 		url = /login
//		method = post
//		payload = loginPayload
//		200 = ResponseLoginOK
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*loginPayload)
	uDomain := domain.MustGetDomain(ctx)
	currentUser, err := models.NewModelsManager().FindUserByEmailDomain(pl.Email, uDomain)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(pl.Password)) != nil {
		c.ForbiddenResponse(w, errors.New("wrong email or password"))
		return
	}

	if currentUser.Status == models.RegisteredUserStatus {
		c.ForbiddenResponse(w, errors.New("not verified"))
		return
	}

	if currentUser.Status == models.BlockedUserStatus {
		c.ForbiddenResponse(w, errors.New("this account has been blocked"))
		return
	}

	c.createLoginResponse(w, currentUser)

}
