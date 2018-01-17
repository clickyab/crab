package user

import (
	"context"
	"errors"
	"net/http"

	"fmt"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/framework/controller"
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

	res := aaa.RegisterUserPayload{
		Email:     p.Email,
		Password:  p.Password,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Mobile:    p.Mobile,
		LegalName: p.LegalName,
	}
	usr, err := m.RegisterUser(res, d.ID)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return nil, errors.New("error registering user")
		}
		if mysqlError.Number == 1062 {
			return nil, fmt.Errorf("duplicate email %s", p.Email)
		}
	}
	e := verifyEmail(usr, r)
	if e == errTooSoon {
		return nil, nil
	}
	assert.Nil(e)
	return nil, nil
}
