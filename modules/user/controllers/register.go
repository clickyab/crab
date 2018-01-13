package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type registerPayload struct {
	Email     string `json:"email" validate:"email" error:"email is invalid"`
	Password  string `json:"password" validate:"gt=5" error:"password is too short"`
	FirstName string `json:"first_name" validate:"required" error:"first name is invalid"`
	Mobile    string `json:"mobile"`
	LastName  string `json:"last_name" validate:"required" error:"last name is invalid"`
	LegalName string `json:"legal_name" validate:"omitempty,gt=5"`
}

// @Route {
// 		url = /register
//		method = post
//		payload = registerPayload
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) register(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*registerPayload)
	m := aaa.NewAaaManager()
	d := domain.MustGetDomain(ctx)

	isUnique, err := m.CheckEmailUniqueness(pl.Email)
	if err != nil {
		u.BadResponse(w, trans.E(err.Error()))
		return
	}

	if !isUnique {
		u.BadResponse(w, trans.E("Duplicate email"))
		return
	}

	res := aaa.RegisterUserPayload{
		Email:     pl.Email,
		Password:  pl.Password,
		FirstName: pl.FirstName,
		LastName:  pl.LastName,
		Mobile:    pl.Mobile,
		LegalName: pl.LegalName,
	}
	usr, err := m.RegisterUser(res, d.ID)
	if err != nil {
		u.BadResponse(w, trans.E("error registering userPayload"))
		return
	}
	e := verifyEmail(usr, r)
	if e == errTooSoon {
		u.OKResponse(w, "user has been created")
		return
	}
	assert.Nil(e)
	u.OKResponse(w, nil)

}
