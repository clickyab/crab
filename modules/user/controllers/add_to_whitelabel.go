package user

import (
	"context"
	"net/http"

	"strings"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/mailer"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
)

var forbiddenUserRoles = config.RegisterString("crab.modules.user.forbidden.roles",
	"Admin,SuperAdmin,Owner",
	"forbidden roles",
)

// @Validate {
// }
type addUserToWhitelabelPayload struct {
	Email           string               `json:"email" validate:"email" error:"email is invalid"`
	Password        string               `json:"password" validate:"gt=5" error:"password is too short"`
	RolesID         []int64              `json:"roles_id" validate:"required" error:"roles id is required"`
	FirstName       string               `json:"first_name" validate:"required" error:"first name is invalid"`
	LastName        string               `json:"last_name" validate:"required" error:"last name is invalid"`
	Mobile          string               `json:"mobile" validate:"omitempty"`
	AccountType     aaa.AccountType      `json:"account_type" validate:"required"`
	NotifyUser      bool                 `json:"notify_user"`
	Advantage       int                  `json:"advantage" validate:"max=99,min=0" error:"should be in 0-99"`
	CorporationInfo *CorporationInfoType `json:"corporation_info" validate:"omitempty"`

	roles         []*aaa.Role
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
	roles, err := m.FindRolesByDomainExclude(p.RolesID, d.ID, getForbiddenRoles()...)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error, can't find role id, or role is forbidden")
		return errors.InvalidOrForbiddenRoleErr
	}

	p.roles = roles
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
	m := aaa.NewAaaManager()

	currentUser := authz.MustGetUser(ctx)

	_, ok := aaa.CheckPermOn(currentUser, currentUser, "add_to_whitelabel_user", p.currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}

	user := aaa.User{
		Email:     p.Email,
		Password:  p.Password,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Advantage: p.Advantage,
		Status:    aaa.ActiveUserStatus,
		Cellphone: mysql.NullString{String: p.Mobile, Valid: true},
	}

	err := m.WhiteLabelAddUserRoles(ctx, &user, p.corporation, p.currentDomain.ID, p.roles)
	if err != nil {
		return nil, err
	}
	if p.NotifyUser {
		go func() {
			err = mailer.LoginInfoEmail(&user, p.Password, r)
			assert.Nil(err)
		}()
	}
	user.Roles = p.roles
	res := c.createUserResponse(&user, nil, nil)
	return &res, nil
}

func getForbiddenRoles() []string {
	return strings.Split(forbiddenUserRoles.String(), ",")
}
