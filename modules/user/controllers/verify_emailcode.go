package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
)

// @Validate {
// }
type verifyEmailCodePayload struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

// verifyEmailCode is verify email
// @Rest {
// 		url = /email/verify
// 		method = post
// }
func (c *Controller) verifyEmailCode(ctx context.Context, r *http.Request, p *verifyEmailCodePayload) (*ResponseLoginOK, error) {
	u, e := verifyCode(ctx, fmt.Sprintf("%s%s%s", hasher(p.Email+emailVerifyPath.String()), delimiter, p.Code))

	if e != nil || u.Status != aaa.RegisteredUserStatus || strings.ToLower(p.Email) != strings.ToLower(u.Email) {
		return nil, errors.New("bad request data")
	}

	u.Status = aaa.ActiveUserStatus
	assert.Nil(aaa.NewAaaManager().UpdateUser(u))
	return c.createLoginResponse(u), nil
}
