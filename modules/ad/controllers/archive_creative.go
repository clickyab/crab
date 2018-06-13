package controllers

import (
	"context"
	"net/http"

	"strconv"

	"time"

	"clickyab.com/crab/modules/ad/errors"
	adOrm "clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/fatih/structs"
	"github.com/rs/xmux"
)

func validateArchiveCreative(ctx context.Context) (*domainOrm.Domain, *adOrm.Creative, error) {
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return nil, nil, errors.InvalidIDErr
	}
	dm := domain.MustGetDomain(ctx)

	targetCreative, err := adOrm.NewOrmManager().FindCreativeByID(idInt)
	if err != nil {
		return nil, nil, errors.InvalidIDErr
	}
	_, err = orm.NewOrmManager().FindCampaignByIDDomain(targetCreative.CampaignID, dm.ID)
	if err != nil {
		return nil, nil, errors.InvalidIDErr
	}

	if targetCreative.ArchivedAt.Valid {
		return nil, nil, errors.AlReadyArchivedErr
	}
	return dm, targetCreative, nil
}

// archiveCreative archive creative with creative id
// @Rest {
// 		url = /archive/:id
//		protected = true
// 		method = patch
// 		resource = archive_creative:self
// }
func (c *Controller) archiveCreative(ctx context.Context, r *http.Request) (*adOrm.Creative, error) {
	currentUser := authz.MustGetUser(ctx)
	token := authz.MustGetToken(ctx)

	currentDomain, targetCreative, err := validateArchiveCreative(ctx)
	if err != nil {
		return nil, err
	}

	uScope, ok := currentUser.HasOn("archive_creative", targetCreative.UserID, currentDomain.ID, false, false, permission.ScopeGlobal, permission.ScopeSelf)
	if !ok {
		return nil, errors.AccessDenied
	}
	// get now date time
	targetCreative.ArchivedAt = mysql.NullTime{Valid: true, Time: time.Now()}

	err = targetCreative.SetAuditUserData(currentUser.ID, token, currentDomain.ID, "archive_creative", uScope)
	if err != nil {
		return nil, err
	}

	err = targetCreative.SetAuditDomainID(currentDomain.ID)
	if err != nil {
		return nil, err
	}
	err = targetCreative.SetAuditOwnerID(targetCreative.UserID)
	if err != nil {
		return nil, err
	}
	d := structs.Map(targetCreative)
	err = targetCreative.SetAuditDescribe(d, "archive creative")
	if err != nil {
		return nil, err
	}
	err = targetCreative.SetAuditEntity("creative", targetCreative.ID)
	if err != nil {
		return nil, err
	}
	// apply archive
	err = adOrm.NewOrmManager().UpdateCreative(targetCreative)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error when archive creative")
		return nil, errors.ArchiveCreativeDbErr
	}
	return targetCreative, nil
}
