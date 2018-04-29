package controllers

import (
	"context"
	"net/http"

	"time"

	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/fatih/structs"
)

// @Validate{
//}
type copyCampaignPayload struct {
	Title string `json:"title" validate:"gt=3"`
}

// copy a campaign by id
// @Rest {
// 		url = /copy/:id
//		protected = true
// 		method = patch
//		resource = copy_campaign:self
// }
func (c Controller) copyCampaign(ctx context.Context, r *http.Request, p *copyCampaignPayload) (*orm.Campaign, error) {
	baseData, err := CheckUserCamapignDomain(ctx)
	if err != nil {
		return nil, err
	}
	uScope, ok := aaa.CheckPermOn(baseData.owner, baseData.currentUser, "copy_campaign", baseData.campaign.DomainID)
	if !ok {
		return nil, errors.AccessDenied
	}

	err = baseData.campaign.SetAuditUserData(baseData.currentUser.ID, false, 0, "copy_campaign", uScope)
	if err != nil {
		return nil, err
	}

	// check for archive campaign
	if baseData.campaign.ArchivedAt.Valid && baseData.campaign.ArchivedAt.Time.Before(time.Now()) {
		return nil, errors.ArchivedEditError
	}

	baseData.campaign.Title = p.Title
	baseData.campaign.ID = 0

	d := structs.Map(baseData.campaign)
	err = baseData.campaign.SetAuditDescribe(d, "copy new campaign")
	if err != nil {
		return nil, err
	}

	db := orm.NewOrmManager()
	err = db.CreateCampaign(baseData.campaign)
	if err != nil {
		return nil, errors.DuplicateNameError
	}

	return baseData.campaign, nil
}
