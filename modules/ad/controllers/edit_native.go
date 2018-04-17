package controllers

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campignErr "clickyab.com/crab/modules/campaign/errors"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	dm "clickyab.com/crab/modules/domain/orm"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/rs/xmux"
)

// @Validate{
//}
type editNativePayload struct {
	URL             string                `json:"url" validate:"required"`
	MaxBid          int64                 `json:"max_bid" validate:"required,gt=0"`
	Assets          NativeAssetPayload    `json:"assets"`
	CurrentUser     *aaa.User             `json:"-"`
	CurrentDomain   *dm.Domain            `json:"-"`
	CurrentCampaign *campaignOrm.Campaign `json:"-"`
	CurrentCreative *orm.Creative         `json:"-"`
	CampaignOwner   *aaa.User             `json:"-"`
	CreativeOwner   *aaa.User             `json:"-"`
	Images          []*uploadOrm.Upload   `json:"-"`
	Logo            *uploadOrm.Upload     `json:"-"`
	Video           *uploadOrm.Upload     `json:"-"`
	Icon            *uploadOrm.Upload     `json:"-"`
}

func (p *editNativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := p.Assets.Validate(ctx, w, r); err != nil {
		return err
	}
	currentUser := authz.MustGetUser(ctx)
	p.CurrentUser = currentUser
	cpManager := campaignOrm.NewOrmManager()
	dmn := domain.MustGetDomain(ctx)
	p.CurrentDomain = dmn
	cID, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	currentCreative, err := orm.NewOrmManager().FindCreativeByIDAndType(cID, orm.CreativeNativeType)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.CurrentCreative = currentCreative
	creativeOwner, err := aaa.NewAaaManager().FindUserWithParentsByID(currentCreative.UserID, dmn.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.CreativeOwner = creativeOwner
	targetCampaign, err := cpManager.FindCampaignByIDDomain(currentCreative.CampaignID, dmn.ID)
	if err != nil {
		return campignErr.NotFoundError(targetCampaign.ID)
	}
	p.CurrentCampaign = targetCampaign
	campaignOwner, err := aaa.NewAaaManager().FindUserWithParentsByID(targetCampaign.UserID, dmn.ID)
	if err != nil {
		return campignErr.NotFoundError(targetCampaign.ID)
	}
	p.CampaignOwner = campaignOwner

	return nil
}

// addNativeCreative to campaign
// @Rest {
// 		url = /native/:id
//		protected = true
// 		method = put
// 		resource = edit_creative:self
// }
func (c Controller) editNativeCreative(ctx context.Context, r *http.Request, p *editNativePayload) (*orm.CreativeSaveResult, error) {
	err := checkEditPerm(p)
	if err != nil {
		return nil, err
	}
	p.CurrentCreative.URL = p.URL
	p.CurrentCreative.Status = orm.PendingCreativeStatus
	p.CurrentCreative.Type = orm.CreativeNativeType
	p.CurrentCreative.MaxBid = p.MaxBid

	db := orm.NewOrmManager()
	assets := generateNativeAssets(p.Assets, p.Images, p.Icon, p.Logo, p.Video)
	res, err := db.EditCreative(p.CurrentCreative, assets)
	if err != nil {
		return res, errors.DBError
	}

	return res, nil
}

func checkEditPerm(p *editNativePayload) error {
	// check creative perm
	_, ok := aaa.CheckPermOn(p.CreativeOwner, p.CurrentUser, "edit_creative", p.CurrentDomain.ID)
	if !ok {
		return errors.AccessDenied
	}
	// check campaign perm
	_, ok = aaa.CheckPermOn(p.CampaignOwner, p.CurrentUser, "edit_campaign", p.CurrentDomain.ID)
	if !ok {
		return errors.AccessDenied
	}

	return checkFilePerm(p)
}

func checkFilePerm(p *editNativePayload) error {
	if len(p.Assets.Images) > 0 {
		images, err := validateImage("image", p.Assets.Images...)
		if err != nil {
			return err
		}
		p.Images = images
		for i := range images { //check perm on upload file
			if fileOwnerCheckPerm(images[i], p.CurrentDomain.ID, p.CurrentUser) != nil {
				return err
			}
		}
	}

	if p.Assets.Icon != "" {
		icon, err := validateImage("icon", p.Assets.Icon)
		if err != nil {
			return err
		}
		if fileOwnerCheckPerm(icon[0], p.CurrentDomain.ID, p.CurrentUser) != nil {
			return err
		}
		p.Icon = icon[0]
	}

	if p.Assets.Logo != "" {
		logo, err := validateImage("logo", p.Assets.Logo)
		if err != nil {
			return err
		}
		if fileOwnerCheckPerm(logo[0], p.CurrentDomain.ID, p.CurrentUser) != nil {
			return err
		}
		p.Logo = logo[0]
	}

	if p.Assets.Video != "" {
		video, err := validateVideo(p.Assets.Video)
		if err != nil {
			return err
		}
		if fileOwnerCheckPerm(video, p.CurrentDomain.ID, p.CurrentUser) != nil {
			return err
		}
		p.Video = video
	}
	return nil
}
