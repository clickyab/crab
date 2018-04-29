package controllers

import (
	"context"
	"net/http"
	"time"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/framework/controller"
	"github.com/clickyab/services/mysql"
	"github.com/fatih/structs"
)

// archive will archive campaign
// @Rest {
// 		url = /archive/:id
//		protected = true
// 		method = patch
//		resource = archive_campaign:self
// }
func (c *Controller) archive(ctx context.Context, r *http.Request) (*controller.NormalResponse, error) {
	baseData, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return nil, err
	}

	uScope, ok := aaa.CheckPermOn(baseData.owner, baseData.currentUser, "archive_campaign", baseData.campaign.DomainID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err = baseData.campaign.SetAuditUserData(baseData.currentUser.ID, false, 0, "archive_campaign", uScope)
	if err != nil {
		return nil, err
	}

	// if campaign current mode is archive nothing can be done
	if baseData.campaign.ArchivedAt.Valid && baseData.campaign.ArchivedAt.Time.Before(time.Now()) {
		//nothing can do
		return nil, errors.ChangeArchiveError
	}

	baseData.campaign.ArchivedAt = mysql.NullTime{Valid: true, Time: time.Now()}

	d := structs.Map(baseData.campaign)
	err = baseData.campaign.SetAuditDescribe(d, "archive campaign")
	if err != nil {
		return nil, err
	}

	db := orm.NewOrmManager()
	err = db.UpdateCampaign(baseData.campaign)
	if err != nil {
		return nil, errors.UpdateCampaignErr
	}

	return nil, nil
}
