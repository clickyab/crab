package user

import (
	"context"
	"encoding/json"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/location/location"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"clickyab.com/crab/modules/user/ucfg"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/hub"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/mysql"
)

// Controller is the controller for the userPayload package
// @Route {
//		group = /user
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

type auditData struct {
	Username string      `json:"username"`
	Action   string      `json:"action"`
	Class    string      `json:"class"`
	Data     interface{} `json:"data"`
}

type responseLoginOK struct {
	Token   string `json:"token"`
	Account userResponse   `json:"account"`
}

var (
	_ = controller.ErrorResponseMap{}
	_ = controller.ErrorResponseSimple{}
	_ = domain.Access
)

// MustGetUser try to get back the userPayload to system
func (u Controller) MustGetUser(ctx context.Context) *aaa.User {
	return authz.MustGetUser(ctx)
}

//// MustGetPerm try to get back the perm granted to system
//func (u Controller) MustGetPerm(ctx context.Context) base.Permission {
//	return authz.MustGetCurrentPerm(ctx)
//}

//// MustGetScope try to get back the scope granted to system
//func (u Controller) MustGetScope(ctx echo.Context) base.UserScope {
//	return authz.MustGetCurrentScope(ctx)
//}

func (u Controller) storeData(r *http.Request, token string) error {
	err := kv.NewEavStore(token).SetSubKey("ua", r.UserAgent()).Save(ucfg.TokenTimeout.Duration())
	if err != nil {
		return err
	}
	return kv.NewEavStore(token).SetSubKey("ip", framework.RealIP(r)).Save(ucfg.TokenTimeout.Duration())
}

//// PermDoubleCheck try to double check perm on the object base on its owner id
//// TODO : write code generator for this if you can
//func (u Controller) PermDoubleCheck(ctx echo.Context, objectUserID int64) (base.UserScope, bool) {
//	usr := u.MustGetUser(ctx)
//	perm := u.MustGetPerm(ctx)
//	other := usr
//	if usr.ID != objectUserID {
//		var err error
//		other, err = aaa.NewAaaManager().FindUserByID(objectUserID)
//		assert.Nil(err)
//	}
//	return usr.HasPermOn(perm, other.ID, other.ResellerID.Int64)
//}

// String make this one a stringer
func (u auditData) String() string {
	r, _ := json.Marshal(u)

	return string(r)
}

func audit(username, action, class string, data interface{}) {
	hub.Publish(
		"audit",
		auditData{
			Username: username,
			Action:   action,
			Class:    class,
			Data:     data,
		},
	)
}

type userResponse struct {
	ID            int64          `json:"id"`
	Email         string         `json:"email"`
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	Avatar        *string        `json:"avatar, omitempty"`
	CityName      string         `json:"city_name, omitempty"`
	CityID        int64          `json:"city_id, omitempty"`
	ProvinceName  string         `json:"province_name, omitempty"`
	ProvinceID    int64          `json:"province_id, omitempty"`
	CountryName   string         `json:"country_name, omitempty"`
	CountryID     int64          `json:"country_id, omitempty"`
	LandLine      *string        `json:"land_line, omitempty"`
	Cellphone     *string        `json:"cellphone, omitempty"`
	PostalCode    *string        `json:"postal_code, omitempty"`
	Address       *string        `json:"address, omitempty"`
	Gender        aaa.GenderType `json:"gender, omitempty"`
	SSN           *string        `json:"ssn, omitempty"`
	LegalName     string         `json:"legal_name, omitempty"`
	LegalRegister *string        `json:"legal_register, omitempty"`
	EconomicCode  *string        `json:"economic_code, omitempty"`
}

func (c Controller) createLoginResponseWithToken(w http.ResponseWriter, user *aaa.User, token string) {
	u := userResponse{}

	u.ID = user.ID
	u.Email = user.Email
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Avatar = nullableToPointer(user.Avatar)
	u.LandLine = nullableToPointer(user.LandLine)
	u.Cellphone = nullableToPointer(user.Cellphone)
	u.PostalCode = nullableToPointer(user.PostalCode)
	u.Address = nullableToPointer(user.Address)
	u.Gender = user.Gender
	u.SSN = nullableToPointer(user.SSN)
	if user.Corporation != nil {

		u.LegalName = user.Corporation.LegalName
		u.LegalRegister = nullableToPointer(user.Corporation.LegalRegister)
		u.EconomicCode = nullableToPointer(user.Corporation.EconomicCode)
	}

	var l *location.CityInfo
	if user.CityID.Valid {
		m := location.NewLocationManager()
		m.FindAllByCityID(user.CityID.Int64)

		u.CityName = l.CityName
		u.CityID = l.CityID
		u.ProvinceName = l.ProvinceName
		u.ProvinceID = l.ProvinceID
		u.CountryName = l.CountryName
		u.CountryID = l.CountryID
	}
	res := responseLoginOK{
		Token:   token,
		Account: u,
	}

	c.OKResponse(w, res)
}

func nullableToPointer(v mysql.NullString) *string {
	if v.Valid {
		return &v.String
	}
	return nil
}

func (c Controller) createLoginResponse(w http.ResponseWriter, user *aaa.User) {
	token := aaa.GetNewToken(user)
	c.createLoginResponseWithToken(w, user, token)
}
