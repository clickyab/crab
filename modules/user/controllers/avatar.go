package user

import (
	"context"
	"errors"
	"net/http"

	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/assert"
)

type avatarPayload struct {
	Avatar string `json:"avatar"`
}

// route for add/update user avatar
// @Rest {
// 		url = /avatar
//		protected = true
// 		method = put
// }
func (c *Controller) avatar(ctx context.Context, r *http.Request, p *avatarPayload) (*ResponseLoginOK, error) {
	cu := c.MustGetUser(ctx)
	m := aaa.NewAaaManager()
	if p.Avatar == "" {
		cu.Avatar.String = ""
		cu.Avatar.Valid = false
	} else {
		up, err := model.NewModelManager().FindUploadByID(p.Avatar)
		if err != nil {
			return nil, errors.New("avatar not found")
		}
		cu.Avatar = stringToNullString(up.ID)
	}

	err := m.UpdateUser(cu)
	assert.Nil(err)
	return c.createLoginResponse(cu), nil
}
