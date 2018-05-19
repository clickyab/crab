package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate{
//}
type editDomainPayload struct {
	Description  string                 `json:"description" validate:"omitempty"`
	Attributes   map[string]interface{} `json:"attributes" validate:"omitempty"`
	Status       orm.DomainStatus       `json:"status" validate:"required"`
	targetDomain *orm.Domain            `json:"-"`
}

func (p *editDomainPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidDomainStatus
	}
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	m := orm.NewOrmManager()
	domainObj, err := m.FindDomainByID(idInt)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.targetDomain = domainObj
	return nil
}

// editDomain to domain
// @Rest {
// 		url = /edit/:id
//		protected = true
// 		method = put
//		resource = god:global
// }
func (c *Controller) editDomain(ctx context.Context, r *http.Request, p *editDomainPayload) (*orm.Domain, error) {
	currentUser := authz.MustGetUser(ctx)
	currentDomain := domain.MustGetDomain(ctx)
	// check permission
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "god", currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDeniedErr
	}
	p.targetDomain.Status = p.Status
	p.targetDomain.Description = mysql.NullString{Valid: p.Description != "", String: p.Description}
	p.targetDomain.Attributes = p.Attributes
	m := orm.NewOrmManager()
	err := m.UpdateDomain(p.targetDomain)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("can't update domain")
		return nil, errors.UpdateDomainErr
	}
	return p.targetDomain, nil
}
