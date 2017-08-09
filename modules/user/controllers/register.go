package user

import (
	"context"
	"net/http"

	middleware2 "clickyab.com/crab/modules/domain/middleware"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/trans"
)

type registerPayload struct {
	Email    string      `json:"email" validate:"email" error:"email is invalid"`
	Password string      `json:"password" validate:"gt=5" error:"password is too short"`
	UserType aaa.UserTyp `json:"user_type"`
}

type responseRegister struct {
	ID       int64       `json:"id"`
	Email    string      `json:"email"`
	Token    string      `json:"token"`
	UserType aaa.UserTyp `json:"user_type"`
}

// @Route {
// 		url = /register
//		method = post
//      payload = registerPayload
//		middleware = middleware2.Access
//		200 = responseRegister
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) Register(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*registerPayload)
	if !pl.UserType.IsValid() {
		u.BadResponse(w, middleware.GroupError{
			string(pl.UserType): trans.E("user type is invalid"),
		})
		return
	}
	m := aaa.NewAaaManager()
	domain := middleware2.MustGetDomain(ctx)
	user, err := m.RegisterUser(pl.Email, pl.Password, pl.UserType, domain.ID)
	if err != nil {
		u.BadResponse(w, trans.E("error registering user"))
		return
	}
	token := aaa.GetNewToken(user)
	u.OKResponse(w, responseRegister{
		ID:       user.ID,
		Email:    user.Email,
		Token:    token,
		UserType: user.UserType,
	})

}
