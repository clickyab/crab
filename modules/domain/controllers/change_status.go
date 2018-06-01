package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type changeDomainStatusPayload struct {
	DomainStatus  orm.DomainStatus `json:"domain_status" validate:"required"`
	currentDomain *orm.Domain      `json:"-"`
	targetDomain  *orm.Domain      `json:"-"`
}

func (p *changeDomainStatusPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.DomainStatus.IsValid() {
		return errors.InvalidDomainStatus
	}
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm
	// find target domain by id
	targetDomain, err := orm.NewOrmManager().FindDomainByID(idInt)
	if err != nil {
		return errors.DomainNotFoundError(idInt)
	}
	p.targetDomain = targetDomain
	return nil
}

// changeDomainStatus change domain status by id, status can be enable or disable
// @Rest {
// 		url = /change-domain-status/:id
//		protected = true
// 		method = put
// 		resource = change_domain_status:superGlobal
// }
func (c *Controller) changeDomainStatus(ctx context.Context, r *http.Request, p *changeDomainStatusPayload) (*orm.Domain, error) {
	m := orm.NewOrmManager()
	// apply status
	p.targetDomain.Status = p.DomainStatus
	err := m.UpdateDomain(p.targetDomain)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error when change domain status")
		return nil, errors.UpdateStatusDbErr
	}

	return p.targetDomain, nil
}
