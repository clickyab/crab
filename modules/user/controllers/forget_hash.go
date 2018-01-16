package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"
)

// forgetCallBack is the url coming from sent email
// @Rest {
// 		url = /password/verify/:token
// 		method = get
// }
func (c *Controller) checkForgetHash(ctx context.Context, r *http.Request) (*ResponseLoginOK, error) {
	t := xmux.Param(ctx, "token")
	u, e := verifyCode(ctx, t)
	if e != nil {
		return nil, errors.New("verify code mismatch")

	}
	s, _, e := genVerifyCode(u, "change password")
	assert.Nil(e)
	return c.createLoginResponseWithToken(u, s), nil
}
