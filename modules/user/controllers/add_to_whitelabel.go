package user

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/clickyab/services/xlog"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
	gom "github.com/go-sql-driver/mysql"
)

// AccountType is the user account type
type (
	// AccountType is the user account type
	// @Enum{
	// }
	AccountType string
)

const (
	// PersonalUser male
	PersonalUser AccountType = "personal"
	// CorporationUser corporation
	CorporationUser AccountType = "corporation"
)

// @Validate {
// }
type addUserToWhitelabelPayload struct {
	Email           string      `json:"email" validate:"email" error:"email is invalid"`
	Password        string      `json:"password" validate:"gt=5" error:"password is too short"`
	RoleID          int64       `json:"role_id" validate:"required" error:"role id is required"`
	FirstName       string      `json:"first_name" validate:"required" error:"first name is invalid"`
	LastName        string      `json:"last_name" validate:"required" error:"last name is invalid"`
	Mobile          string      `json:"mobile" validate:"required,lt=15"`
	AccountType     AccountType `json:"account_type" validate:"required"`
	NotifyUser      bool        `json:"notify_user" validate:"required"`
	CorporationInfo struct {
		LegalName     string `json:"legal_name" validate:"omitempty"`
		LegalRegister string `json:"legal_register" validate:"omitempty"`
		EconomicCode  string `json:"economic_code" validate:"omitempty"`
	} `json:"corporation_info" validate:"omitempty`
}

func (l *addUserToWhitelabelPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !l.AccountType.IsValid() {
		return errors.InvalidAccountType
	}

	if l.AccountType == CorporationUser {
		if l.CorporationInfo.LegalName == "" {
			return errors.InvalidCorporationLegalName
		}
	}

	return nil
}

// register is for register user
// @Rest {
// 		url = /add-to/whitelabel
//		protected = true
// 		method = post
//		resource = add_to_whitelabel_user:global
// }
func (c *Controller) registerToWhitelabel(ctx context.Context, r *http.Request, p *addUserToWhitelabelPayload) (*controller.NormalResponse, error) {
	m := aaa.NewAaaManager()
	d := domain.MustGetDomain(ctx)

	user := aaa.User{
		Email:     p.Email,
		Password:  p.Password,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Status:    aaa.ActiveUserStatus,
		Cellphone: mysql.NullString{String: p.Mobile, Valid: true},
	}

	corp := aaa.Corporation{}
	if p.AccountType == CorporationUser {
		corp.LegalName = p.CorporationInfo.LegalName
		corp.LegalRegister = mysql.NullString{String: p.CorporationInfo.LegalRegister, Valid: (p.CorporationInfo.LegalRegister != "")}
		corp.EconomicCode = mysql.NullString{String: p.CorporationInfo.EconomicCode, Valid: (p.CorporationInfo.EconomicCode != "")}
	}

	_, err := m.FindRoleByIDAndDomain(p.RoleID, d.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			xlog.GetWithError(ctx, err).Debug("can't find role id in add user to whitelabel route")
		}
		return nil, errors.InvalidRoleIDErr
	}

	err = m.RegisterUser(&user, &corp, d.ID, p.RoleID)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return nil, trans.E("error registering user")
		}
		if mysqlError.Number == 1062 {
			return nil, trans.E("duplicate email %s", p.Email)
		}

		xlog.GetWithError(ctx, err).Debug("error in add new user to whitelabel route")
		return nil, errors.DBError
	}

	if p.NotifyUser {
		go func() {
			err = mailer.LoginInfoEmail(&user, p.Password, r)
			assert.Nil(err)
		}()
	}

	return nil, nil
}
