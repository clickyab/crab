package user

import (
	"context"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/location/location"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/framework/controller"
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

// ResponseLoginOK login or ping or other response
type ResponseLoginOK struct {
	Token             string       `json:"token"`
	Account           userResponse `json:"account"`
	Perms             *[]string    `json:"perms,omitempty"`
	ImpersonatorToken string       `json:"impersonator_token,omitempty"`
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

type userResponse struct {
	ID            int64                  `json:"id"`
	Email         string                 `json:"email"`
	FirstName     string                 `json:"first_name"`
	LastName      string                 `json:"last_name"`
	Avatar        string                 `json:"avatar,omitempty"`
	CityName      string                 `json:"city_name,omitempty"`
	CityID        int64                  `json:"city_id,omitempty"`
	ProvinceName  string                 `json:"province_name,omitempty"`
	Province      string                 `json:"province,omitempty"`
	CountryName   string                 `json:"country_name,omitempty"`
	CountryID     int64                  `json:"country_id,omitempty"`
	LandLine      string                 `json:"land_line,omitempty"`
	Cellphone     string                 `json:"cellphone,omitempty"`
	PostalCode    string                 `json:"postal_code,omitempty"`
	Address       string                 `json:"address,omitempty"`
	Gender        aaa.GenderType         `json:"gender,omitempty"`
	SSN           string                 `json:"ssn,omitempty"`
	Perms         *[]string              `json:"perms,omitempty"`
	LegalName     string                 `json:"legal_name,omitempty"`
	LegalRegister string                 `json:"legal_register,omitempty"`
	EconomicCode  string                 `json:"economic_code,omitempty"`
	Balance       int64                  `json:"balance"`
	Attributes    mysql.GenericJSONField `json:"attributes,omitempty"`
	Role          *aaa.Role              `json:"role,omitempty"`
	Managers      []aaa.ManagerUser      `json:"managers,omitempty"`
	Advantage     int                    `json:"advantage,omitempty"`
}

func (u Controller) createUserResponse(user *aaa.User, perms *[]string, managers []aaa.ManagerUser) userResponse {
	us := userResponse{}
	us.ID = user.ID
	us.Email = user.Email
	us.FirstName = user.FirstName
	us.LastName = user.LastName
	us.Avatar = user.Avatar.String
	us.LandLine = user.LandLine.String
	us.Cellphone = user.Cellphone.String
	us.PostalCode = user.PostalCode.String
	us.Attributes = user.Attributes
	us.Address = user.Address.String
	us.Balance = user.Balance
	us.Perms = perms
	us.Role = user.Role
	us.Managers = managers
	us.Advantage = user.Advantage
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
		us.Province = l.Province
		us.CountryName = l.CountryName
		us.CountryID = l.CountryID
	}
	return us
}

func (u Controller) createLoginResponse(user *aaa.User) *ResponseLoginOK {
	token := aaa.GetNewToken(user)
	return &ResponseLoginOK{
		Token:   token,
		Account: u.createUserResponse(user, nil, nil),
	}
}
