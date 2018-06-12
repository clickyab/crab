package controllers

import (
	"context"
	"net/http"

	"strconv"

	"time"

	"clickyab.com/crab/modules/ad/errors"
	adOrm "clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/fatih/structs"
	"github.com/rs/xmux"
)

// @Validate{
//}
type archiveCreativePayload struct {
	targetCreative *adOrm.Creative   `json:"-"`
	currentDomain  *domainOrm.Domain `json:"-"`
}

func (p *archiveCreativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm

	orm := adOrm.NewOrmManager()
	p.targetCreative, err = orm.FindCreativeByID(idInt)
	if err != nil {
		return errors.InvalidIDErr
	}
	if p.targetCreative.ArchivedAt.Valid{
		return errors.AlReadyArchivedErr
	}
	return nil
}

// archiveCreative archive creative with creative id
// @Rest {
// 		url = /archive/:id
//		protected = true
// 		method = patch
// 		resource = archive_creative:self
// }
func (c *Controller) archiveCreative(ctx context.Context, r *http.Request, p *archiveCreativePayload) (*adOrm.Creative, error) {
	currentUser := authz.MustGetUser(ctx)
	token := authz.MustGetToken(ctx)

	uScope, ok := currentUser.HasOn("archive_creative", p.targetCreative.UserID, p.currentDomain.ID, false, false, permission.ScopeGlobal, permission.ScopeSelf)
	if !ok {
		return nil, errors.AccessDenied
	}
	// get now date time
	now := time.Now()
	p.targetCreative.ArchivedAt = mysql.NullTime{Valid: true, Time: now}

	err := p.targetCreative.SetAuditUserData(currentUser.ID, token, p.currentDomain.ID, "archive_creative", uScope)
	if err != nil {
		return nil, err
	}

	err = p.targetCreative.SetAuditDomainID(p.currentDomain.ID)
	if err != nil {
		return nil, err
	}
	err = p.targetCreative.SetAuditOwnerID(p.targetCreative.UserID)
	if err != nil {
		return nil, err
	}
	d := structs.Map(p.targetCreative)
	err = p.targetCreative.SetAuditDescribe(d, "archive creative")
	if err != nil {
		return nil, err
	}
	err = p.targetCreative.SetAuditEntity("creative", p.targetCreative.ID)
	if err != nil {
		return nil, err
	}

	// apply archive
	orm := adOrm.NewOrmManager()
	err = orm.UpdateCreative(p.targetCreative)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error when archive creative")
		return nil, errors.ArchiveCreativeDbErr
	}
	return p.targetCreative, nil
}
