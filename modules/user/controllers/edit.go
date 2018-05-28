package user

import (
	"context"
	"net/http"

	"time"

	"database/sql"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
	"github.com/clickyab/services/xlog"
)

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

	corporation *aaa.Corporation
}

// edit route for edit personal profile
// @Rest {
// 		url = /update
//		protected = true
//		method = put
// }
func (c *Controller) edit(ctx context.Context, r *http.Request, p *userPayload) (*ResponseLoginOK, error) {
	cu := c.MustGetUser(ctx)
	m := aaa.NewAaaManager()

	cc, e := m.FindCorporationByUserID(cu.ID)

	if e != nil {
		if e != sql.ErrNoRows {
			xlog.GetWithError(ctx, errors.DBError)
			return nil, errors.DBError
		}
		// corporation not found user is personal
		if !p.Gender.IsValid() || p.Gender == aaa.NotSpecifiedGender {
			return nil, errors.GenderInvalid
		}
		cu.Gender = p.Gender
	} else { // user is corporation
		if p.LegalName == "" {
			return nil, errors.LegalEmptyErr
		}
		p.corporation = cc
		p.corporation.LegalName = p.LegalName
		p.corporation.EconomicCode = stringToNullString(p.EconomicCode)
		p.corporation.LegalRegister = stringToNullString(p.LegalRegister)

		e = m.UpdateCorporation(p.corporation)
		if e != nil {
			xlog.Get(ctx).Error("update user corporation: ", e)
			return nil, errors.UpdateCorporationErr
		}
	}

	cu.CityID = intToNullInt64(p.CityID)
	cu.LandLine = stringToNullString(p.LandLine)
	cu.Cellphone = stringToNullString(p.CellPhone)
	cu.PostalCode = stringToNullString(p.PostalCode)
	cu.FirstName = p.FirstName
	cu.LastName = p.LastName
	cu.Address = stringToNullString(p.Address)
	cu.SSN = stringToNullString(p.SSN)
	cu.UpdatedAt = time.Now()

	e = m.UpdateUser(cu)
	if e != nil {
		xlog.Get(ctx).Error("error message when updating user:", e)
		return nil, errors.UpdateUserErr
	}
	return &ResponseLoginOK{
		Token:   authz.MustGetToken(ctx),
		Account: c.createUserResponse(cu, nil, nil),
	}, nil
}

func stringToNullString(val string) mysql.NullString {
	return mysql.NullString{String: val, Valid: val != ""}
}

func intToNullInt64(val int64) mysql.NullInt64 {
	return mysql.NullInt64{Int64: val, Valid: val != 0}
}
