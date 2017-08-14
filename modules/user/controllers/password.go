package user

import (
	"context"
	"net/http"
	"strconv"

	"github.com/clickyab/services/assert"
	"github.com/rs/xmux"

	"errors"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/trans"
)

type forget struct {
	Email string `json:"email"`
}

// forgetPassword
// @Route {
//		url = /password/forget
//      method = post
//      payload = forget
//      200 = controller.NormalResponse
//      400 = controller.ErrorResponseSimple
// }
func (c Controller) forgetPassword(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	payload, _ := c.MustGetPayload(ctx).(*forget)
	mail := payload.Email

	user, err := aaa.NewAaaManager().FindUserByEmail(mail)
	if err != nil {
		c.BadResponse(w, errors.New("Email not found"))
		return
	}

	// need token for notif
	aaa.GetNewToken(user)

	// todo: Send Notification

	c.OKResponse(w, nil)
}

// @Validate {
// }
type callBack struct {
	Token       string `json:"token" validate:"required"`
	NewPassword string `json:"new_password" validate:"gt=5" error:"password is too short"`
}

// forgetCallBack is the url coming from sent email
// 		@Route {
// 		url = /password/callback
// 		method = put
//      payload = callBack
//		200 = controller.NormalResponse
// 		400 = controller.ErrorResponseSimple
// }
func (c Controller) forgetCallBack(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	key := xmux.Param(ctx, "key")
	payload := c.MustGetPayload(ctx).(*callBack)

	var userID string
	if userID = kv.NewEavStore(key).SubKey("userID"); userID != "" {
		c.BadResponse(w, trans.E("expired key"))
		return
	}

	id, err := strconv.ParseInt(userID, 10, 0)
	assert.Nil(err)

	u, e := aaa.NewAaaManager().FindUserByID(id)
	assert.Nil(e)

	e = u.ChangePassword(payload.NewPassword)
	if e != nil {
		switch e {
		case aaa.ErrorWrongPassword:
			c.BadResponse(w, trans.EE(e))
			return
		case aaa.ErrorOldPass:
			c.BadResponse(w, trans.EE(e))
			return
		default:
			assert.Nil(e)
		}
	}

	c.OKResponse(w, nil)
}

// @Validate {
// }
type changePassword struct {
	CurrentPassword string `json:"current_password" validate:"gt=5"`
	NewPassword     string `json:"new_password" validate:"gt=5"`
}

// changePassword
// @Route {
//		url = /password/change
//      method = put
//      payload = changePassword
//		middleware = authz.Authenticate
//      200 = controller.NormalResponse
//      400 = controller.ErrorResponseSimple
// }
func (c Controller) changePassword(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	u := c.MustGetUser(ctx)
	payload := c.MustGetPayload(ctx).(*changePassword)
	e := u.UpdatePassword(payload.CurrentPassword, payload.NewPassword)
	if e != nil {
		switch e {
		case aaa.ErrorWrongPassword:
			c.BadResponse(w, trans.EE(e))
			return
		case aaa.ErrorOldPass:
			c.BadResponse(w, trans.EE(e))
			return
		default:
			assert.Nil(e)
		}
	}

	c.OKResponse(w, nil)
}
