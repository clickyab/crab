package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/financial/errors"
	"clickyab.com/crab/modules/financial/orm"
	"github.com/clickyab/services/xlog"
	gom "github.com/go-sql-driver/mysql"
	"github.com/rs/xmux"
)

// @Validate{
//}
type editGatewayPayload struct {
	Name          string            `json:"name" validation:"required"`
	Status        orm.GatewayStatus `json:"status" validation:"required"`
	targetGateway *orm.Gateway      `json:"-"`
}

func (p *editGatewayPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidGatewayStatusErr
	}
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	m := orm.NewOrmManager()
	gateway, err := m.FindGatewayByID(idInt)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.targetGateway = gateway
	return nil
}

// editGateway to financial
// @Rest {
// 		url = /gateways/:id
//		protected = true
// 		method = put
// 		resource = edit_gateway:superGlobal
// }
func (c *Controller) editGateway(ctx context.Context, r *http.Request, p *editGatewayPayload) (*orm.Gateway, error) {
	p.targetGateway.Status = p.Status
	p.targetGateway.Name = p.Name
	m := orm.NewOrmManager()
	err := m.UpdateGateway(p.targetGateway)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			xlog.GetWithError(ctx, err).Debug("can't edit gateway")
			return nil, errors.EditGatewayErr
		}
		if mysqlError.Number == 1062 {
			return nil, errors.GatewayAlreadyExistErr
		}
		xlog.GetWithError(ctx, err).Debug("can't edit gateway")
		return nil, errors.EditGatewayErr
	}
	return p.targetGateway, nil
}
