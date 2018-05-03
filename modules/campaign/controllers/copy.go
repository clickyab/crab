package controllers

import (
	"context"
	"database/sql"
	"net/http"

	"time"

	adOrm "clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/campaign/errors"
	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/fatih/structs"
)

// copy a campaign by id
// @Rest {
// 		url = /copy/:id
//		protected = true
// 		method = patch
//		resource = copy_campaign:self
// }
func (c Controller) copyCampaign(ctx context.Context, r *http.Request) (*orm.Campaign, error) {
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

	oldID := baseData.campaign.ID
	baseData.campaign.Title += " - copy"
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

	err = copySchedule(oldID, baseData.campaign.ID)
	if err != nil {
		return nil, err
	}

	err = copyAttributes(oldID, baseData.campaign.ID)
	if err != nil {
		return nil, err
	}

	err = copyReportsReceivers(oldID, baseData.campaign.ID)
	if err != nil {
		return nil, err
	}

	err = copyCreativesAssets(oldID, baseData.campaign.ID)
	if err != nil {
		return nil, err
	}

	return baseData.campaign, nil
}

func copySchedule(oldID, newID int64) error {
	db := orm.NewOrmManager()

	sc, err := db.GetSchedule(oldID)
	if err != nil {
		return errors.NotFoundSchedule
	}

	if sc != nil {
		sc.CampaignID = newID
		sc.ID = 0

		err = db.CreateSchedule(sc)
		if err != nil {
			return errors.TimeScheduleError
		}
	}

	return nil
}

func copyAttributes(oldID, newID int64) error {
	db := orm.NewOrmManager()

	attrs, err := db.FindCampaignAttributesByCampaignID(oldID)
	if err != nil && err != sql.ErrNoRows {
		return errors.NotFoundAttributes
	}

	if attrs != nil {
		attrs.CampaignID = newID
		_, err := db.AttachCampaignAttributes(*attrs)
		if err != nil {
			return errors.UpdateError
		}
	}

	return nil
}

func copyReportsReceivers(oldID, newID int64) error {
	db := orm.NewOrmManager()

	recs := db.ListCampaignReportReceiversWithFilter("campaign_id = ?", oldID)
	if len(recs) > 0 {
		var ids []int64
		for _, v := range recs {
			ids = append(ids, v.UserID)
		}
		err := db.UpdateReportReceivers(ids, newID)
		if err != nil {
			return errors.UpdateError
		}
	}

	return nil
}

func copyCreativesAssets(oldID, newID int64) error {
	adDb := adOrm.NewOrmManager()
	creatives := adDb.ListCreativesWithFilter("campaign_id=?", oldID)

	for _, cr := range creatives {
		assets := adDb.ListAssetsWithFilter("creative_id=?", cr.ID)

		cr.ID = 0
		cr.CampaignID = newID

		var ap []*adOrm.Asset
		for k := range assets {
			assets[k].CreativeID = 0
			ap = append(ap, &assets[k])
		}

		_, err := adDb.AddCreative(&cr, ap)
		if err != nil {
			return errors.CreateError
		}
	}

	return nil
}
