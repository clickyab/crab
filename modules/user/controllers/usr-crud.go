package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/config"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/eav"
	"github.com/clickyab/services/random"
	"golang.org/x/crypto/bcrypt"
)

// @Validate {
// }
type loginPayload struct {
	Email    string `json:"email" validate:"email" error:"email is invalid"`
	Password string `json:"password" validate:"gt=5" error:"password is too short"`
}

type responseLoginOK struct {
	ID       int64       `json:"id"`
	Email    string      `json:"email"`
	Token    string      `json:"token"`
	UserType aaa.UserTyp `json:"user_type"`
}

// login user in system
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
	domain := middleware.MustGetDomain(ctx)
	currentUser, err := aaa.NewAaaManager().FindUserByEmailDomian(pl.Email, domain)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(currentUser.Password), []byte(pl.Password)) != nil {
		c.ForbiddenResponse(w, errors.New("wrong password"))
		return
	}
	token := getNewToken(currentUser)
	c.OKResponse(w, responseLoginOK{
		ID:       currentUser.ID,
		Email:    currentUser.Email,
		Token:    token,
		UserType: currentUser.UserType,
	})
}

func getNewToken(user *aaa.User) string {
	t := fmt.Sprintf("%d:%s", user.ID, <-random.ID)
	generated := eav.NewEavStore(t).SetSubKey("token", user.AccessToken)
	assert.Nil(generated.Save(ucfg.TokenTimeout.Duration()))
	return t
}
