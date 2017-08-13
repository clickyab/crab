package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type registerPayload struct {
	Email       string      `json:"email" validate:"email" error:"email is invalid"`
	Password    string      `json:"password" validate:"gt=5" error:"password is too short"`
	FirstName   string      `json:"first_name" validate:"required" error:"first name is invalid"`
	LastName    string      `json:"last_name" validate:"required" error:"last name is invalid"`
	CompanyName string      `json:"company_name"`
	UserType    aaa.UserTyp `json:"user_type" validate:"required"`
}

// @Route {
// 		url = /register
//		method = post
//		payload = registerPayload
//		200 = responseLoginOK
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
	d := domain.MustGetDomain(ctx)
	if pl.UserType == aaa.CorporationUserTyp && pl.CompanyName == "" {
		u.BadResponse(w, trans.E("company name required for corporation users"))
		return
	}
	usr, err := m.RegisterUser(pl.Email, pl.Password, pl.UserType, pl.FirstName, pl.LastName, pl.CompanyName, d.ID)
	if err != nil {
		u.BadResponse(w, trans.E("error registering user"))
		return
	}
	token := aaa.GetNewToken(usr)
	u.createLoginResponse(w, usr, token)

}
