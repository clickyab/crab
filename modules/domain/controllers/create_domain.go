package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	gom "github.com/go-sql-driver/mysql"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"github.com/clickyab/services/xlog"
)

// @Validate{
//}
type createDomainPayload struct {
	Name        string                 `json:"name" validate:"required"`
	Description string                 `json:"description" validate:"omitempty"`
	Attributes  map[string]interface{} `json:"attributes" validate:"omitempty"`
	Status      orm.DomainStatus       `json:"status" validate:"required"`
}

// createDomainResult create domain result
type createDomainResult struct {
	DomainID int64            `json:"domain_id"`
	Name     string           `json:"name"`
	Status   orm.DomainStatus `json:"status"`
}

func (p *createDomainPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidDomainStatus
	}
	return nil
}

// createDomain to domain
// @Rest {
// 		url = /create
//		protected = true
// 		method = post
//		resource = create_new_domain:global
// }
func (c *Controller) createDomain(ctx context.Context, r *http.Request, p *createDomainPayload) (*createDomainResult, error) {
	currentUser := authz.MustGetUser(ctx)
	currentDomain := domain.MustGetDomain(ctx)
	// check permission
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "create_new_domain", currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDeniedErr
	}
	// create domain object
	newDomain := &orm.Domain{
		Name:        p.Name,
		Description: mysql.NullString{Valid: true, String: p.Description},
		Attributes:  p.Attributes,
		Status:      p.Status,
	}
	m := orm.NewOrmManager()
	// create new domain
	err := m.CreateDomain(newDomain)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			xlog.GetWithError(ctx, err).Debug("can't insert new domain")
			return nil, errors.CreateDomainErr
		}
		if mysqlError.Number == 1062 {
			return nil, errors.AlreadyExistErr
		}
	}
	res := &createDomainResult{
		DomainID: newDomain.ID,
		Name:     newDomain.Name,
		Status:   newDomain.Status,
	}
	return res, nil
}
