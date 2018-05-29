package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/ucfg"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/trans"
	gom "github.com/go-sql-driver/mysql"
)

// @Validate {
// }
type registerPayload struct {
	Email     string `json:"email" validate:"email" error:"email is invalid"`
	Password  string `json:"password" validate:"gt=5" error:"password is too short"`
	FirstName string `json:"first_name" validate:"required" error:"first name is invalid"`
	Mobile    string `json:"mobile" validate:"lt=15"`
	LastName  string `json:"last_name" validate:"required" error:"last name is invalid"`
	LegalName string `json:"legal_name" validate:"omitempty,gt=5"`
}

// register is for register user
// @Rest {
// 		url = /register
// 		method = post
// }
func (c *Controller) register(ctx context.Context, r *http.Request, p *registerPayload) (*controller.NormalResponse, error) {
	m := aaa.NewAaaManager()
	d := domain.MustGetDomain(ctx)

	user := aaa.User{
		Email:      p.Email,
		Password:   p.Password,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		DomainLess: false,
		Cellphone:  mysql.NullString{String: p.Mobile, Valid: true},
	}

	corp := aaa.Corporation{}
	if p.LegalName != "" {
		corp.LegalName = p.LegalName
	}

	db := aaa.NewAaaManager()
	role, err := db.FindRoleByName(ucfg.DefaultRole.String())
	if err != nil {
		return nil, errors.NotFoundRoleOfDomain(ucfg.DefaultRole.String(), d.ID)
	}

	err = m.RegisterUserWrapper(&user, &corp, d.ID, role.ID)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return nil, trans.E("error registering user")
		}
		if mysqlError.Number == 1062 {
			return nil, trans.E("duplicate email %s", p.Email)
		}
	}
	e := verifyEmail(&user, r)
	if e == errTooSoon {
		return nil, nil
	}
	assert.Nil(e)
	return nil, nil
}
