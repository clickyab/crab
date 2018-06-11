package controllers

import (
	"context"

	"clickyab.com/crab/modules/ad/errors"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
)

// fileOwnerCheckPerm owner check
func fileOwnerCheckPerm(ctx context.Context, image *uploadOrm.Upload, d int64, currentUser *aaa.User) error {
	targetFileOwner, err := aaa.NewAaaManager().FindUserByID(image.UserID)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error on FindUserWithParentsByID")
		return errors.AssetsPermErr
	}
	// check campaign perm
	_, ok := currentUser.HasOn("edit_creative", targetFileOwner.ID, d, false, false, permission.ScopeSelf, permission.ScopeGlobal)
	if !ok {
		return errors.AccessDenied
	}
	return nil
}
