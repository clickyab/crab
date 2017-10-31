package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/location/location"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"clickyab.com/crab/modules/user/ucfg"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/kv"
)

// Controller is the controller for the userPayload package
// @Route {
//		group = /user
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

//type auditData struct {
//	Username string      `json:"username"`
//	Action   string      `json:"action"`
//	Class    string      `json:"class"`
//	Data     interface{} `json:"data"`
//}

// ResponseLoginOK login or ping or other response
type ResponseLoginOK struct {
	Token   string       `json:"token"`
	Account userResponse `json:"account"`
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
//func (u auditData) String() string {
//	r, _ := json.Marshal(u)
//
//	return string(r)
//}

//func audit(username, action, class string, data interface{}) {
//	hub.Publish(
//		"audit",
//		auditData{
//			Username: username,
//			Action:   action,
//			Class:    class,
//			Data:     data,
//		},
//	)
//}

type userResponse struct {
	ID            int64          `json:"id"`
	Email         string         `json:"email"`
	FirstName     string         `json:"first_name"`
	LastName      string         `json:"last_name"`
	Avatar        string         `json:"avatar,omitempty"`
	CityName      string         `json:"city_name,omitempty"`
	CityID        int64          `json:"city_id,omitempty"`
	ProvinceName  string         `json:"province_name,omitempty"`
	ProvinceID    int64          `json:"province_id,omitempty"`
	CountryName   string         `json:"country_name,omitempty"`
	CountryID     int64          `json:"country_id,omitempty"`
	LandLine      string         `json:"land_line,omitempty"`
	Cellphone     string         `json:"cellphone,omitempty"`
	PostalCode    string         `json:"postal_code,omitempty"`
	Address       string         `json:"address,omitempty"`
	Gender        aaa.GenderType `json:"gender,omitempty"`
	SSN           string         `json:"ssn,omitempty"`
	LegalName     string         `json:"legal_name,omitempty"`
	LegalRegister string         `json:"legal_register,omitempty"`
	EconomicCode  string         `json:"economic_code,omitempty"`
}

func (u Controller) createLoginResponseWithToken(w http.ResponseWriter, user *aaa.User, token string) {
	us := userResponse{}

	us.ID = user.ID
	us.Email = user.Email
	us.FirstName = user.FirstName
	us.LastName = user.LastName
	us.Avatar = user.Avatar.String
	us.LandLine = user.LandLine.String
	us.Cellphone = user.Cellphone.String
	us.PostalCode = user.PostalCode.String
	us.Address = user.Address.String
	if user.Gender != aaa.NotSpecifiedGender {
		us.Gender = user.Gender
	}
	us.SSN = user.SSN.String

	if c, e := aaa.NewAaaManager().FindCorporationByUserID(us.ID); e == nil {
		us.LegalName = c.LegalName
		us.LegalRegister = c.LegalRegister.String
		us.EconomicCode = c.EconomicCode.String
	}

	if user.CityID.Valid {
		m := location.NewLocationManager()
		l := m.FindAllByCityID(user.CityID.Int64)

		us.CityName = l.CityName
		us.CityID = l.CityID
		us.ProvinceName = l.ProvinceName
		us.ProvinceID = l.ProvinceID
		us.CountryName = l.CountryName
		us.CountryID = l.CountryID
	}
	res := ResponseLoginOK{
		Token:   token,
		Account: us,
	}

	u.OKResponse(w, res)
}

func (u Controller) createLoginResponse(w http.ResponseWriter, user *aaa.User) {
	token := aaa.GetNewToken(user)
	u.createLoginResponseWithToken(w, user, token)
}
