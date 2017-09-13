package user

import (
	"context"
	"net/http"

	"time"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type userPayload struct {
	Email       string           `json:"email" validate:"omitempty, email"`
	Avatar      string           `json:"avatar" validate:"omitempty, url"`
	CityID      int64            `json:"city_id" validate:"omitempty"`
	LandLine    string           `json:"land_line" validate:"omitempty"`
	CellPhone   string           `json:"cell_phone" validate:"omitempty"`
	PostalCode  string           `json:"postal_code" validate:"omitempty"`
	FirstName   string           `json:"first_name" validate:"omitempty,gt=2"`
	LastName    string           `json:"last_name" validate:"omitempty,gt=2"`
	Address     string           `json:"address" validate:"omitempty"`
	Gender      aaa.GenderType   `json:"gender" validate:"omitempty"`
	SSN         string           `json:"ssn" validate:"omitempty"`
	Corporation *aaa.Corporation `json:"corporation" validate:"omitempty"`
}

func (u *userPayload) ValidateExtra(ctx context.Context) error {
	if u.Gender != "" {
		if u.Gender.IsValid() {
			return middleware.GroupError{
				"validate": trans.E("Gender is not valid"),
			}
		}
	}

	return nil
}

// EditPersonal route for edit personal profile
// @Route {
// 		url = /personal
//		method = put
//		payload = userPayload
//		middleware = authz.Authenticate
//		200 = responseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) Edit(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*userPayload)
	if !pl.Gender.IsValid() || pl.Gender == aaa.NotSpecifiedGender {
		u.BadResponse(w, middleware.GroupError{
			string(pl.Gender): trans.E("gender is invalid"),
		})
		return
	}

	cu := u.MustGetUser(ctx)
	m := aaa.NewAaaManager()

	var cc *aaa.Corporation
	var e error
	if pl.Corporation != nil {
		cc, e = m.FindCorporationByUserID(cu.ID)
		if e != nil {
			u.BadResponse(w, trans.E("Personal userPayload not allowed to update corporate account"))
		}
	}

	cu.Email = pl.Email
	cu.Avatar = stringToNullString(pl.Avatar)
	cu.CityID = intToNullInt64(pl.CityID)
	cu.LandLine = stringToNullString(pl.LandLine)
	cu.Cellphone = stringToNullString(pl.CellPhone)
	cu.PostalCode = stringToNullString(pl.PostalCode)
	cu.FirstName = pl.FirstName
	cu.LastName = pl.LastName
	cu.Address = stringToNullString(pl.Address)
	cu.Gender = pl.Gender
	cu.SSN = stringToNullString(pl.SSN)
	cu.UpdatedAt = time.Now()

	e = m.UpdateUser(cu)
	if e != nil {
		u.BadResponse(w, trans.EE(e))
		return
	}

	if pl.Corporation != nil {
		cc.LegalName = pl.Corporation.LegalName
		cc.EconomicCode = stringToNullString(pl.Corporation.LegalName)
		cc.LegalRegister = stringToNullString(pl.Corporation.LegalName)

		e = m.UpdateCorporation(cc)
		if e != nil {
			u.BadResponse(w, trans.EE(e))
			return
		}
	}

	u.createLoginResponseWithToken(w, cu, authz.MustGetToken(ctx))
}

func stringToNullString(val string) mysql.NullString {
	return mysql.NullString{String: val, Valid: val == ""}
}

func intToNullInt64(val int64) mysql.NullInt64 {
	return mysql.NullInt64{Int64: val, Valid: val != 0}
}
