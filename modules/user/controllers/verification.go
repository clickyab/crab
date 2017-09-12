package user

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"errors"

	"time"

	"crypto/sha1"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"

	"github.com/clickyab/services/notification"
	"github.com/clickyab/services/random"
	"github.com/rs/xmux"
)

const (
	verifyKey      = "VERIFY"
	verifySub      = "ID"
	dump           = "d"
	activeTemplate = "Email"
)

type verifyIdResponse responseLoginOK

// verifyId is verify code
// @Route {
// 		url = /verify/:hash/:key
//		method = get
//		200 = verifyIdResponse
//		401 = controller.ErrorResponseSimple
//		403 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) verifyEmail(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	h := xmux.Param(ctx, "hash")
	k := xmux.Param(ctx, "key")
	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKey, h))

	if kw.SubKey(k) != dump {
		ctrl.BadResponse(w, errors.New("Code is not valid"))
		return
	}
	id, e := strconv.ParseInt(kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKey, h)).SubKey(verifySub), 10, 64)
	assert.Nil(e)

	m := aaa.NewAaaManager()
	cu, e := m.FindUserByID(id)
	assert.Nil(e)
	if cu.Status != aaa.RegisteredUserStatus {
		ctrl.ForbiddenResponse(w, nil)
		return
	}
	cu.Status = aaa.ActiveUserStatus
	m.UpdateUser(cu)
	ctrl.createLoginResponse(w, cu)
}

// @Validate{
// }
type verifyResendPayload struct {
	Email    string `json:"email_string" validate:"required|email"`
}

// verifyResend will send an email again
// @Route {
// 		url = /verify/resend
//		method = post
//		200 = verifyResendPayload
//		404 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) verifyResend(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := ctrl.MustGetPayload(ctx).(*verifyResendPayload)
	u, e := aaa.NewAaaManager().FindUserByEmail(pl.Email)
	if e != nil {
		ctrl.NotFoundResponse(w, nil)
		return
	}

	sendVerifyCode(u)
}

func sendVerifyCode(u *aaa.User) {
	hash := func() string {
		h := sha1.New()
		h.Write([]byte(u.Email))
		return fmt.Sprintf("%x", h.Sum(nil))
	}()

	kw := kv.NewEavStore(fmt.Sprintf("%s_%s", verifyKey, hash))
	// TODO: get it from config
	exp := 5 * time.Hour

	if len(kw.AllKeys()) != 0 {
		t := time.Now().Add(exp).Add(-2 * time.Minute).Sub(time.Now())
		if t < kw.TTL() {
			return
		}
	}

	key := fmt.Sprintf("%s%s%s", <-random.ID, <-random.ID, <-random.ID)
	strconv.FormatInt(u.ID, 64)
	kw.SetSubKey(verifySub, strconv.FormatInt(u.ID, 64))
	kw.SetSubKey(key, dump)
	assert.Nil(kw.Save(exp))

	go func() {
		a := notification.Packet{Platform: notification.MailType, To: []string{u.Email}}
		notification.Send(fmt.Sprintf("%s/%s", hash, key), fmt.Sprintf("%s/%s", hash, key), a)
	}()
}
