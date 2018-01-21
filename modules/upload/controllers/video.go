package controllers

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"path/filepath"

	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

type getVideoResponse struct {
	Ready string `json:"ready"`
}

// getVideoReady find video into the system
// @Rest {
// 		url = /video/:id
//		protected = true
// 		method = get
// }
func (c *Controller) getVideoReady(ctx context.Context, r *http.Request) (*getVideoResponse, error) {
	authz.MustGetUser(ctx)
	m := model.NewModelManager()
	decURL, err := base64.URLEncoding.DecodeString(xmux.Param(ctx, "id"))
	if err != nil {
		return nil, t9e.G("wrong id")
	}
	file, err := m.FindUploadByID(string(decURL))
	if err != nil {
		return nil, t9e.G("wrong id")
	}
	//check if file ready or not
	_, err = os.Stat(filepath.Join(UPath.String(), file.ID))
	if err != nil {
		return &getVideoResponse{Ready: "pending"}, nil
	}
	return &getVideoResponse{Ready: "done"}, nil
}
