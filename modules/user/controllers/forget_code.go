package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/gettext/t9e"
)

// @Validate {
// }
type forgetCodePayload struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required"`
}

// forgetCallBack is the url coming from sent email
// @Rest {
// 		url = /password/verify/
//		method = post
// }
func (c *Controller) checkForgetCode(ctx context.Context, r *http.Request, p *forgetCodePayload) (*ResponseLoginOK, error) {
	u, e := verifyCode(ctx, fmt.Sprintf("%s%s%s", hasher(p.Email+passwordVerifyPath.String()), delimiter, p.Code))
	if e != nil || strings.ToLower(p.Email) != strings.ToLower(u.Email) {
		return nil, t9e.G("bad request data")
	}

	s, _, e := genVerifyCode(u, "change password")
	assert.Nil(e)
	return c.createLoginResponseWithToken(u, s), nil
}