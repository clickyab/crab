package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/mailer"
	"clickyab.com/crab/modules/user/middleware/authz"
	"clickyab.com/crab/modules/user/services"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/xlog"
)

// @Validate {
// }
type addUserToWhitelabelPayload struct {
	Email           string               `json:"email" validate:"email" error:"email is invalid"`
	Password        string               `json:"password" validate:"gt=5" error:"password is too short"`
	RoleID          int64                `json:"role_id" validate:"required" error:"role id is required"`
	FirstName       string               `json:"first_name" validate:"required" error:"first name is invalid"`
	LastName        string               `json:"last_name" validate:"required" error:"last name is invalid"`
	Mobile          string               `json:"mobile" validate:"omitempty"`
	AccountType     aaa.AccountType      `json:"account_type" validate:"required"`
	NotifyUser      bool                 `json:"notify_user"`
	Advantage       int64                `json:"advantage" validate:"max=99,min=0" error:"should be in 0-99"`
	CorporationInfo *CorporationInfoType `json:"corporation_info" validate:"omitempty"`

	role          *aaa.Role
	currentDomain *orm.Domain
	corporation   *aaa.Corporation
}

// CorporationInfoType corporation info type
type CorporationInfoType struct {
	LegalName     string `json:"legal_name"`
	LegalRegister string `json:"legal_register"`
	EconomicCode  string `json:"economic_code"`
}

func (p *addUserToWhitelabelPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	d := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	p.currentDomain = d
	if !p.AccountType.IsValid() {
		return errors.InvalidAccountType
	}

	if p.AccountType == aaa.CorporationUser {
		if p.CorporationInfo == nil || p.CorporationInfo.LegalName == "" {
			return errors.InvalidCorporationLegalName
		}
		corp := &aaa.Corporation{}
		corp.LegalName = p.CorporationInfo.LegalName
		corp.LegalRegister = mysql.NullString{String: p.CorporationInfo.LegalRegister, Valid: p.CorporationInfo.LegalRegister != ""}
		corp.EconomicCode = mysql.NullString{String: p.CorporationInfo.EconomicCode, Valid: p.CorporationInfo.EconomicCode != ""}
		p.corporation = corp
	}
	m := aaa.NewAaaManager()
	// validate if role ids is valid and not in forbidden keys
	role, err := m.FindRoleByID(p.RoleID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error, can't find role id, or role is forbidden")
		return errors.InvalidOrForbiddenRoleErr
	}

	currentUser.SetUserRole(d.ID)
	if currentUser.Role.Level <= role.Level {
		return errors.InvalidOrForbiddenRoleErr
	}
	p.role = role
	return nil
}

// register is for register user, account_type can be : personal or corporation
// @Rest {
// 		url = /whitelabel/add
//		protected = true
// 		method = post
//		resource = add_to_whitelabel_user:global
// }
func (c *Controller) registerToWhitelabel(ctx context.Context, r *http.Request, p *addUserToWhitelabelPayload) (*userResponse, error) {
	user := aaa.User{
		Email:      p.Email,
		Password:   p.Password,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		Advantage:  mysql.NullInt64{Valid: p.Advantage != 0, Int64: p.Advantage},
		DomainLess: false,
		Status:     aaa.ActiveUserStatus,
		Cellphone:  mysql.NullString{String: p.Mobile, Valid: true},
	}

	err := services.WhiteLabelAddUser(ctx, &user, p.corporation, p.currentDomain.ID, p.role)
	if err != nil {
		return nil, err
	}
	if p.NotifyUser {
		go func() {
			err = mailer.LoginInfoEmail(&user, p.Password, r)
			assert.Nil(err)
		}()
	}
	user.Role = p.role
	res := c.createUserResponse(&user, nil, nil)
	return &res, nil
}
