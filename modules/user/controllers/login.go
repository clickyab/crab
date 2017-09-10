package user

import (
	"context"
	"errors"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"golang.org/x/crypto/bcrypt"
)

// @Validate {
// }
type loginPayload struct {
	Email    string `json:"email" validate:"email" error:"email is invalid"`
	Password string `json:"password" validate:"gt=5" error:"password is too short"`
}

// login userPayload in system
// @Route {
// 		url = /login
//		method = post
//      payload = loginPayload
//		200 = responseLoginOK
//		400 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) login(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := c.MustGetPayload(ctx).(*loginPayload)
	uDomain := domain.MustGetDomain(ctx)
	currentUser, err := aaa.NewAaaManager().FindUserByEmailDomain(pl.Email, uDomain)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(pl.Password)) != nil {
		c.ForbiddenResponse(w, errors.New("wrong password"))
		return
	}
	token := aaa.GetNewToken(currentUser)
	c.createLoginResponse(w, currentUser, token)

}
