package controllers

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

// getDomainDetail get domain detail by domain id
// @Rest {
// 		url = /:id
//		protected = true
// 		method = get
//		resource = get_detail_domain:global
// }
func (c *Controller) getDomainDetail(ctx context.Context, r *http.Request) (*orm.Domain, error) {
	userDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	// check access
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "get_detail_domain", userDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// get domain id from url params
	domainID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	// find domain
	m := orm.NewOrmManager()
	domainObj, err := m.FindDomainByID(domainID)
	if err != nil {
		return nil, errors.DomainNotFoundError(domainID)
	}
	return domainObj, nil
}
