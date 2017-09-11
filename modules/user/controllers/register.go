package user

import (
	"context"
	"math/rand"
	"net/http"

	"fmt"
	"time"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/kv"
	"github.com/clickyab/services/notification"
	"github.com/clickyab/services/trans"
)

// @Validate {
// }
type registerPayload struct {
	Email     string `json:"email" validate:"email" error:"email is invalid"`
	Password  string `json:"password" validate:"gt=5" error:"password is too short"`
	FirstName string `json:"first_name" validate:"required" error:"first name is invalid"`
	Mobile    string `json:"mobile"`
	LastName  string `json:"last_name" validate:"required" error:"last name is invalid"`
	LegalName string `json:"legal_name" validate:"omitempty,gt=5"`
}

// @Route {
// 		url = /register
//		method = post
//		payload = registerPayload
//		400 = controller.ErrorResponseSimple
// }
func (u *Controller) register(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*registerPayload)
	m := aaa.NewAaaManager()
	d := domain.MustGetDomain(ctx)

	res := aaa.RegisterUserPayload{
		Email:     pl.Email,
		Password:  pl.Password,
		FirstName: pl.FirstName,
		LastName:  pl.LastName,
		Mobile:    pl.Mobile,
		LegalName: pl.LegalName,
	}
	usr, err := m.RegisterUser(res, d.ID)
	if err != nil {
		u.BadResponse(w, trans.E("error registering userPayload"))
		return
	}
	//generate activation code
	c := rand.Intn(99999) + 100000
	//save in redis
	for i := 0; i < 10; i++ {
		if len(kv.NewEavStore(fmt.Sprintf("%s%s%d", active, seperator, c)).AllKeys()) == 0 {
			break
		}
		c = rand.Intn(99999) + 100000
	}

	err = kv.NewEavStore(fmt.Sprintf("%s%s%d", active, seperator, c)).
		SetSubKey(userID, fmt.Sprintf("%d", usr.ID)).
		Save(2 * time.Hour)
	assert.Nil(err)
	//send mail to user
	a := notification.Packet{Platform: notification.MailType, To: []string{pl.Email}}
	message := fmt.Sprintf("welcome to crab this is your activation code\n%d", c)
	go func() {
		notification.Send("welcome", message, a)
	}()
	u.OKResponse(w, "user has been created")

}
