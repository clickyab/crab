package user

import (
	"context"
	"net/http"

	"time"

	"errors"

	upload "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
)

type avatarPayload struct {
	Avatar string `json:"avatar"`
}

// route for add/update user avatar
// @Route {
// 		url = /avatar
//		method = put
//		payload = avatarPayload
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) avatar(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*avatarPayload)

	cu := u.MustGetUser(ctx)
	m := aaa.NewAaaManager()
	if pl.Avatar == "" {
		cu.Avatar.String = ""
		cu.Avatar.Valid = false
	} else {

		up, err := upload.NewModelManager().FindUploadByID(pl.Avatar)
		if err != nil {
			u.NotFoundResponse(w, errors.New("avatar not found"))
			return
		}
		cu.Avatar = stringToNullString(up.ID)
	}

	err := m.UpdateUser(cu)
	assert.Nil(err)
	u.createLoginResponse(w, cu)
}

// @Validate {
// }
type userPayload struct {
	CityID        int64          `json:"city_id" validate:"omitempty"`
	LandLine      string         `json:"land_line" validate:"omitempty"`
	CellPhone     string         `json:"cell_phone" validate:"omitempty"`
	PostalCode    string         `json:"postal_code" validate:"omitempty"`
	FirstName     string         `json:"first_name" validate:"required,gt=2"`
	LastName      string         `json:"last_name" validate:"required,gt=2"`
	Address       string         `json:"address" validate:"omitempty"`
	Gender        aaa.GenderType `json:"gender" validate:"omitempty"`
	SSN           string         `json:"ssn" validate:"omitempty"`
	LegalName     string         `json:"legal_name" validate:"omitempty"`
	LegalRegister string         `json:"legal_register" validate:"omitempty"`
	EconomicCode  string         `json:"economic_code" validate:"omitempty"`
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

// edit route for edit personal profile
// @Route {
// 		url = /update
//		method = put
//		payload = userPayload
//		middleware = authz.Authenticate
//		200 = ResponseLoginOK
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) edit(ctx context.Context, w http.ResponseWriter, r *http.Request) {
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
	if pl.LegalName != "" {
		cc, e = m.FindCorporationByUserID(cu.ID)
		if e != nil {
			u.BadResponse(w, trans.E("Personal userPayload not allowed to update corporate account"))
		}
	}

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

	if pl.LegalName != "" {
		cc.LegalName = pl.LegalName
		cc.EconomicCode = stringToNullString(pl.LegalName)
		cc.LegalRegister = stringToNullString(pl.LegalName)

		e = m.UpdateCorporation(cc)
		if e != nil {
			u.BadResponse(w, trans.EE(e))
			return
		}
	}

	u.createLoginResponseWithToken(w, cu, authz.MustGetToken(ctx))
}

func stringToNullString(val string) mysql.NullString {
	return mysql.NullString{String: val, Valid: val != ""}
}

func intToNullInt64(val int64) mysql.NullInt64 {
	return mysql.NullInt64{Int64: val, Valid: val != 0}
}
