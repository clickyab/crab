package user

import (
	"encoding/json"

	"github.com/clickyab/services/eav"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/hub"

	"context"

	"clickyab.com/crab/modules/user/config"
	"clickyab.com/crab/modules/user/middleware"

	"net/http"

	"clickyab.com/crab/modules/user/aaa"
)

// Controller is the controller for the user package
// @Route {
//		group = /user
// }
type Controller struct {
	controller.Base
}

type userAudit struct {
	Username string      `json:"username"`
	Action   string      `json:"action"`
	Class    string      `json:"class"`
	Data     interface{} `json:"data"`
}

var (
	_ = controller.ErrorResponseMap{}
	_ = controller.ErrorResponseSimple{}
)

// MustGetUser try to get back the user to system
func (u Controller) MustGetUser(ctx context.Context) *aaa.User {
	return authz.MustGetUser(ctx)
}

//// MustGetPerm try to get back the perm granted to system
//func (u Controller) MustGetPerm(ctx context.Context) base.Permission {
//	return authz.MustGetCurrentPerm(ctx)
//}

//// MustGetScope try to get back the scope granted to system
//func (u Controller) MustGetScope(ctx echo.Context) base.UserScope {
//	return authz.MustGetCurrentScope(ctx)
//}

func (u Controller) storeData(r *http.Request, token string) error {
	err := eav.NewEavStore(token).SetSubKey("ua", r.UserAgent()).Save(ucfg.TokenTimeout.Duration())
	if err != nil {
		return err
	}
	return eav.NewEavStore(token).SetSubKey("ip", framework.RealIP(r)).Save(ucfg.TokenTimeout.Duration())
}

//// PermDoubleCheck try to double check perm on the object base on its owner id
//// TODO : write code generator for this if you can
//func (u Controller) PermDoubleCheck(ctx echo.Context, objectUserID int64) (base.UserScope, bool) {
//	usr := u.MustGetUser(ctx)
//	perm := u.MustGetPerm(ctx)
//	other := usr
//	if usr.ID != objectUserID {
//		var err error
//		other, err = aaa.NewAaaManager().FindUserByID(objectUserID)
//		assert.Nil(err)
//	}
//	return usr.HasPermOn(perm, other.ID, other.ResellerID.Int64)
//}

// String make this one a stringer
func (u userAudit) String() string {
	r, _ := json.Marshal(u)

	return string(r)
}

func audit(username, action, class string, data interface{}) {
	hub.Publish(
		"audit",
		userAudit{
			Username: username,
			Action:   action,
			Class:    class,
			Data:     data,
		},
	)
}
