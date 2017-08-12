package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type personalPayload struct {
	FirstName string         `json:"first_name" validate:"omitempty,gt=2"`
	LastName  string         `json:"last_name" validate:"omitempty,gt=2"`
	Gender    aaa.GenderType `json:"gender" validate:"required"`
	CellPhone string         `json:"cellphone" validate:"omitempty,numeric"`
	Phone     string         `json:"phone" validate:"omitempty,numeric"`
	Address   string         `json:"address" validate:"omitempty,gt=5"`
	CityID    int64          `json:"city_id" validate:"omitempty,numeric"`
}

// EditPersonal route for edit personal profile
// @Route {
// 		url = /personal
//		method = put
//		payload = personalPayload
//		middleware = authz.Authenticate
//		200 = aaa.UserPersonal
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) EditPersonal(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*personalPayload)
	if !pl.Gender.IsValid() {
		u.BadResponse(w, middleware.GroupError{
			string(pl.Gender): trans.E("gender is invalid"),
		})
		return
	}
	// TODO check for valid city id
	currentUser := u.MustGetUser(ctx)
	// the user should be personal type only
	if currentUser.UserType != aaa.PersonalUserTyp {
		u.BadResponse(w, trans.E("user is not personal"))
		return
	}
	m := aaa.NewAaaManager()
	up := &aaa.UserPersonal{
		UserID:    currentUser.ID,
		FirstName: mysql.NullString{String: pl.FirstName, Valid: pl.FirstName != ""},
		LastName:  mysql.NullString{String: pl.LastName, Valid: pl.LastName != ""},
		Gender:    pl.Gender,
		Cellphone: mysql.NullString{String: pl.CellPhone, Valid: pl.CellPhone != ""},
		Phone:     mysql.NullString{String: pl.Phone, Valid: pl.Phone != ""},
		Address:   mysql.NullString{String: pl.Address, Valid: pl.Address != ""},
		CityID:    mysql.NullInt64{Int64: pl.CityID, Valid: pl.CityID != 0},
	}
	err := m.UpdateUserPersonal(up)
	if err != nil {
		u.BadResponse(w, trans.EE(err))
		return
	}
	u.OKResponse(w, up)
}
