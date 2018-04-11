package controllers

import (
	"context"
	"strconv"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	userError "clickyab.com/crab/modules/user/errors"
	"github.com/clickyab/services/framework/controller"
	"github.com/rs/xmux"
)

// Controller is the controller for the user package
// @Route {
//		group = /campaign
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

// Importatn: only add here shared func that use in other package routes

// BaseData basic data needed in almost all routes
type BaseData struct {
	campaign *orm.Campaign
	domain   *domainOrm.Domain
	owner    *aaa.User
}

// CheckUserCamapignDomain func to check campaign exist and is for current user and domain and return base data
func CheckUserCamapignDomain(ctx context.Context) (*BaseData, error) {
	caID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 64)
	if err != nil || caID < 1 {
		return nil, errors.InvalidIDErr
	}

	d := domain.MustGetDomain(ctx)
	db := orm.NewOrmManager()

	ca, err := db.FindCampaignByIDDomain(caID, d.ID)
	if err != nil {
		return nil, errors.NotFoundError(caID)
	}

	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ca.UserID, d.ID)
	if err != nil {
		return nil, userError.NotFoundWithDomainError(d.Name)
	}

	res := BaseData{
		campaign: ca,
		domain:   d,
		owner:    owner,
	}

	return &res, nil
}
