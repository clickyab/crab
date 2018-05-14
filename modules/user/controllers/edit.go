package user

import (
	"context"
	"net/http"

	"time"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
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
}

// edit route for edit personal profile
// @Rest {
// 		url = /update
//		protected = true
//		method = put
// }
func (c *Controller) edit(ctx context.Context, r *http.Request, p *userPayload) (*ResponseLoginOK, error) {
	if !p.Gender.IsValid() || p.Gender == aaa.NotSpecifiedGender {
		return nil, t9e.G("invalid user gender")
	}

	cu := c.MustGetUser(ctx)
	m := aaa.NewAaaManager()

	var cc *aaa.Corporation
	var e error
	if p.LegalName != "" {
		cc, e = m.FindCorporationByUserID(cu.ID)
		if e != nil {
			return nil, t9e.G("personal userPayload not allowed to update corporate account")
		}
	}

	cu.CityID = intToNullInt64(p.CityID)
	cu.LandLine = stringToNullString(p.LandLine)
	cu.Cellphone = stringToNullString(p.CellPhone)
	cu.PostalCode = stringToNullString(p.PostalCode)
	cu.FirstName = p.FirstName
	cu.LastName = p.LastName
	cu.Address = stringToNullString(p.Address)
	cu.Gender = p.Gender
	cu.SSN = stringToNullString(p.SSN)
	cu.UpdatedAt = time.Now()

	e = m.UpdateUser(cu)
	if e != nil {
		return nil, e
	}

	if p.LegalName != "" {
		cc.LegalName = p.LegalName
		cc.EconomicCode = stringToNullString(p.LegalName)
		cc.LegalRegister = stringToNullString(p.LegalName)

		e = m.UpdateCorporation(cc)
		if e != nil {
			return nil, e
		}
	}

	return &ResponseLoginOK{
		Token:   authz.MustGetToken(ctx),
		Account: c.createUserResponse(cu, nil),
	}, nil
}

func stringToNullString(val string) mysql.NullString {
	return mysql.NullString{String: val, Valid: val != ""}
}

func intToNullInt64(val int64) mysql.NullInt64 {
	return mysql.NullInt64{Int64: val, Valid: val != 0}
}
