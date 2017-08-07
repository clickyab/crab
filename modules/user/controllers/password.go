package user

import (
	"context"
	"net/http"

	"errors"

	"clickyab.com/crab/modules/user/aaa"
)

type forget struct {
	Email string `json:"email"`
}

type resp struct {
	Status string `json:"status"`
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
