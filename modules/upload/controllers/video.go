package controllers

import (
	"context"
	"encoding/base64"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// getVideoReady find video into the system
// @Route {
// 		url = /video/:id
//		method = get
//		middleware = authz.Authenticate
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (c Controller) getVideoReady(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	authz.MustGetUser(ctx)
	m := model.NewModelManager()
	decURL, err := base64.URLEncoding.DecodeString(xmux.Param(ctx, "id"))
	if err != nil {
		c.BadResponse(w, errors.New("wrong id"))
		return
	}
	file, err := m.FindUploadByID(string(decURL))
	if err != nil {
		c.BadResponse(w, errors.New("wrong id"))
		return
	}
	//check if file ready or not
	_, err = os.Stat(filepath.Join(UPath.String(), file.ID))
	if err != nil {
		c.JSON(w, http.StatusOK, struct {
			Ready string `json:"ready"`
		}{
			Ready: "pending",
		})
		return
	}
	c.JSON(w, http.StatusOK, struct {
		Ready string `json:"ready"`
	}{
		Ready: "done",
	})
}
