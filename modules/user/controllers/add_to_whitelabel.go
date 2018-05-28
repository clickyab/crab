package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/mailer"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
)

// @Validate {
// }
type addUserToWhitelabelPayload struct {
	Email           string               `json:"email" validate:"email" error:"email is invalid"`
	Password        string               `json:"password" validate:"gt=5" error:"password is too short"`
	RolesID         []int64              `json:"roles_id" validate:"required" error:"roles id is required"`
	FirstName       string               `json:"first_name" validate:"required" error:"first name is invalid"`
	LastName        string               `json:"last_name" validate:"required" error:"last name is invalid"`
	Mobile          string               `json:"mobile" validate:"required,lt=15"`
	AccountType     aaa.AccountType      `json:"account_type" validate:"required"`
	NotifyUser      bool                 `json:"notify_user"`
	Advantage       int                  `json:"advantage" validate:"max=99,min=0" error:"should be in 0-99"`
	CorporationInfo *CorporationInfoType `json:"corporation_info" validate:"omitempty"`
}

// CorporationInfoType corporation info type
type CorporationInfoType struct {
	LegalName     string `json:"legal_name"`
	LegalRegister string `json:"legal_register"`
	EconomicCode  string `json:"economic_code"`
}

func (l *addUserToWhitelabelPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !l.AccountType.IsValid() {
		return errors.InvalidAccountType
	}

	if l.AccountType == aaa.CorporationUser {
		if l.CorporationInfo.LegalName == "" {
			return errors.InvalidCorporationLegalName
		}
	}
	return nil
}

// register is for register user, account_type can be : personal or corporation
// @Rest {
// 		url = /whitelabel/add
//		protected = true
// 		method = post
//		resource = add_to_whitelabel_user:global
// }
func (c *Controller) registerToWhitelabel(ctx context.Context, r *http.Request, p *addUserToWhitelabelPayload) (*controller.NormalResponse, error) {
	m := aaa.NewAaaManager()
	d := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)

	_, ok := aaa.CheckPermOn(currentUser, currentUser, "add_to_whitelabel_user", d.ID, permission.ScopeGlobal)
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

	corp := aaa.Corporation{}
	if p.AccountType == aaa.CorporationUser {
		corp.LegalName = p.CorporationInfo.LegalName
		corp.LegalRegister = mysql.NullString{String: p.CorporationInfo.LegalRegister, Valid: (p.CorporationInfo.LegalRegister != "")}
		corp.EconomicCode = mysql.NullString{String: p.CorporationInfo.EconomicCode, Valid: (p.CorporationInfo.EconomicCode != "")}
	}

	err := m.WhiteLabelAddUserRoles(ctx, &user, &corp, d.ID, p.RolesID)
	if err != nil {
		return nil, err
	}
	if p.NotifyUser {
		go func() {
			err = mailer.LoginInfoEmail(&user, p.Password, r)
			assert.Nil(err)
		}()
	}
	return nil, nil
}
