package user

import (
	"context"
	"fmt"
	"net/http"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type checkActivePayload struct {
	Email  string `json:"email" validate:"required|email"`
	Number int    `json:"number" validate:"required"`
}

// @Route {
// 		url = /active
//		method = patch
//		payload = checkActivePayload
//		200 = controller.NormalResponse
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (u *Controller) checkActive(ctx context.Context, w http.ResponseWriter, r *http.Request) { //send only for registered userPayload
	pl := u.MustGetPayload(ctx).(*checkActivePayload)
	//find userPayload by email
	c := pl.Number
	res := kv.NewEavStore(fmt.Sprintf("%s%s%d", active, seperator, c)).SubKey(userID)
	if res == "" {
		u.BadResponse(w, trans.E("wrong number entered"))
		return
	}
	ID := res
	if ID == "" {
		u.BadResponse(w, trans.E("wrong number entered"))
		return
	}
	m := aaa.NewAaaManager()
	user, err := m.FindUserByEmail(pl.Email)
	if err != nil {
		// userPayload not found (not registered)
		u.NotFoundResponse(w, trans.E("userPayload not found"))
		return
	}
	if user.Status != aaa.RegisteredUserStatus {
		// userPayload is blocked oa already active
		u.BadResponse(w, trans.E("userPayload already activated or blocked"))
		return
	}
	//compare
	if fmt.Sprintf("%d", user.ID) != ID {
		u.BadResponse(w, trans.E("wrong number entered"))
		return
	}
	//all good change active status
	user.Status = aaa.ActiveUserStatus
	assert.Nil(m.UpdateUser(user))
	u.OKResponse(w, nil)
}
