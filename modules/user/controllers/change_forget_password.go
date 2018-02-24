package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate {
// }
type callBackPayload struct {
	NewPassword string `json:"new_password" validate:"gt=5" error:"password is too short"`
}

// changeForgetPassword change forget password
// @Rest {
// 		url = /password/change/:token
// 		method = put
// }
func (c *Controller) changeForgetPassword(ctx context.Context, r *http.Request, p *callBackPayload) (*ResponseLoginOK, error) {
	t := xmux.Param(ctx, "token")

	u, e := verifyCode(ctx, t)
	if e != nil {
		return nil, t9e.G("error while verifying code")
	}

	err := u.ChangePassword(p.NewPassword)
	if err != nil {
		if err == aaa.ErrorOldPass {
			return nil, err
		}

		xlog.GetWithError(ctx, err).Debug("database error on change user password")
		return nil, t9e.G("cant change password")
	}

	return c.createLoginResponse(u), nil
}
