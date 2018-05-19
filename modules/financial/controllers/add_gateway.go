package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	gom "github.com/go-sql-driver/mysql"
)

// @Validate{
//}
type addGatewayPayload struct {
	Name   string            `json:"name" validation:"required"`
	Status orm.GatewayStatus `json:"status" validation:"required"`
}

func (p *addGatewayPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidGatewayStatusErr
	}
	return nil
}

// addGateway to financial
// @Rest {
// 		url = /gateways
//		protected = true
// 		method = post
// 		resource = add_gateway:global
// }
func (c *Controller) addGateway(ctx context.Context, r *http.Request, p *addGatewayPayload) (*orm.Gateway, error) {
	currentDomain := domain.MustGetDomain(ctx)
	currentUser := authz.MustGetUser(ctx)
	// check permission
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "add_gateway", currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	gatewayObj := &orm.Gateway{
		Name:      p.Name,
		Status:    p.Status,
		IsDefault: orm.IsNotDefault,
	}
	m := orm.NewOrmManager()
	err := m.CreateGateway(gatewayObj)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			xlog.GetWithError(ctx, err).Debug("can't insert new gateway")
			return nil, errors.CreateGatewayErr
		}
		if mysqlError.Number == 1062 {
			return nil, errors.GatewayAlreadyExistErr
		}
	}
	return gatewayObj, nil
}
