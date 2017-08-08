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

	valid, err := aaa.NewAaaManager().IsOldPassword(id, payload.NewPassword)
	assert.Nil(err)

	if !valid {
		c.BadResponse(w, trans.E("password were used before"))
		return
	}

	err = aaa.NewAaaManager().UpdateOldPassword(id, payload.NewPassword)
	assert.Nil(err)

	c.OKResponse(w, nil)
}
