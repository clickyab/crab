package user

import (
	"context"
	"net/http"

	"fmt"
	"math/rand"

	"time"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/notification"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type sendActivePayload struct {
	Email string `json:"email" validate:"required|email"`
}

const (
	active    string = "ACTIVE"
	seperator string = "_"
	userID    string = "USERID"
)

const activeTemplate string = ``

// @Route {
// 		url = /active
//		method = post
//		payload = sendActivePayload
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (u *Controller) sendActive(ctx context.Context, w http.ResponseWriter, r *http.Request) { //send only for registered user
	pl := u.MustGetPayload(ctx).(*sendActivePayload)
	//find user by email
	m := aaa.NewAaaManager()
	user, err := m.FindUserByEmail(pl.Email)
	if err != nil {
		// user not found (not registered)
		u.NotFoundResponse(w, trans.E("user not found"))
		return
	}
	if user.Status != aaa.RegisteredUserStatus {
		// user is blocked oa already active
		u.BadResponse(w, trans.E("user already activated or blocked"))
		return
	}
	c := rand.Intn(9999) + 10000
	//save in redis
	if len(kv.NewEavStore(fmt.Sprintf("%s%s%d", active, seperator, c)).AllKeys()) != 0 {
		u.BadResponse(w, trans.E("check your mail please"))
		return
	}
	err = kv.NewEavStore(fmt.Sprintf("%s%s%d", active, seperator, c)).
		SetSubKey(userID, fmt.Sprintf("%d", user.ID)).
		Save(10 * time.Minute)
	assert.Nil(err)
	a := notification.Packet{Platform: notification.MailType, To: []string{pl.Email}}
	go func() {
		notification.Send(fmt.Sprintf("%d", c), activeTemplate, a)
	}()
	u.OKResponse(w, nil)
}
