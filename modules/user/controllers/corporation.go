package user

import (
	"context"
	"net/http"

	"errors"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// @Validate {
// }
type corporation struct {
	Name         string `json:"name" validate:"omitempty,gt=2"`
	FirstName    string `json:"first_name" db:"first_name" validate:"omitempty,gt=2"`
	LastName     string `json:"last_name" db:"last_name" validate:"omitempty,gt=2"`
	Cellphone    string `json:"cellphone" validate:"omitempty,numeric"`
	Phone        string `json:"phone" validate:"omitempty,numeric"`
	Address      string `json:"address" validate:"omitempty,gt=2"`
	EconomicCode string `json:"economic_code" validate:"omitempty,numeric"`
	RegisterCode string `json:"register_code" validate:"omitempty,numeric"`
	CityID       int64  `json:"city_id" validate:"omitempty,numeric"`
	ProvinceID   int64  `json:"province_id" validate:"omitempty,numeric"`
	CountryID    int64  `json:"country_id" validate:"omitempty,numeric"`
}

// editCorporation is used for current user update
// @Route {
// 		url = /corporation
//		method = put
//		payload = corporation
//		middleware = authz.Authenticate
//      200 = aaa.UserCorporation
//		403 = controller.ErrorResponseSimple
// }
func (c Controller) editCorporation(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	payload := c.MustGetPayload(ctx).(*corporation)
	currentUser := c.MustGetUser(ctx)
	println(currentUser.UserType)

	if !currentUser.UserType.IsValid() || currentUser.UserType != aaa.CorporationUserTyp {
		c.BadResponse(w, errors.New("incompatible user type"))
		return
	}

	dbInsert := &aaa.UserCorporation{
		UserID:       currentUser.ID,
		FirstName:    mysql.NullString{String: payload.FirstName, Valid: payload.FirstName != ""},
		Name:         mysql.NullString{String: payload.Name, Valid: payload.Name != ""},
		Address:      mysql.NullString{String: payload.Address, Valid: payload.Address != ""},
		Cellphone:    mysql.NullString{String: payload.Cellphone, Valid: payload.Cellphone != ""},
		EconomicCode: mysql.NullString{String: payload.EconomicCode, Valid: payload.EconomicCode != ""},
		Phone:        mysql.NullString{String: payload.Phone, Valid: payload.Phone != ""},
		RegisterCode: mysql.NullString{String: payload.RegisterCode, Valid: payload.RegisterCode != ""},
		CityID:       mysql.NullInt64{Int64: payload.CityID, Valid: payload.CityID != 0},
		ProvinceID:   mysql.NullInt64{Int64: payload.ProvinceID, Valid: payload.ProvinceID != 0},
		CountryID:    mysql.NullInt64{Int64: payload.CountryID, Valid: payload.CountryID != 0},
	}

	err := aaa.NewAaaManager().UpdateUserCorporation(dbInsert)
	assert.Nil(err)

	c.OKResponse(w, *dbInsert)
}
