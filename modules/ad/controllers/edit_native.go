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
	"github.com/clickyab/services/mysql"
	"github.com/fatih/structs"
	"github.com/rs/xmux"
)

// @Validate{
//}
type editNativePayload struct {
	Name            string                 `json:"name" validate:"required"`
	URL             string                 `json:"url" validate:"required"`
	MaxBid          int64                  `json:"max_bid"`
	Attributes      map[string]interface{} `json:"attributes"`
	Assets          NativeAssetPayload     `json:"assets"`
	CurrentUser     *aaa.User              `json:"-"`
	CurrentDomain   *dm.Domain             `json:"-"`
	CurrentCampaign *campaignOrm.Campaign  `json:"-"`
	CurrentCreative *orm.Creative          `json:"-"`
	CampaignOwner   *aaa.User              `json:"-"`
	CreativeOwner   *aaa.User              `json:"-"`
	Images          []*uploadOrm.Upload    `json:"-"`
	Logos           []*uploadOrm.Upload    `json:"-"`
	Videos          []*uploadOrm.Upload    `json:"-"`
	Icons           []*uploadOrm.Upload    `json:"-"`
}

func (p *editNativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := p.Assets.Validate(ctx, w, r); err != nil {
		return err
	}
	err := emptyValNotPermitted(p.Assets)
	if err != nil {
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

	// extra fields required for app campaigns required
	if targetCampaign.Kind == campaignOrm.AppCampaign {
		if len(p.Assets.Images) == 0 {
			return errors.InvalidIDErr
		}
		if len(p.Assets.Icons) == 0 {
			return errors.IconRequiredErr
		}
		if len(p.Assets.CTAs) == 0 {
			return errors.CtaRequiredErr
		}
		if len(p.Assets.Videos) == 0 {
			return errors.VideoRequiredErr
		}
	}

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

	userToken := authz.MustGetToken(ctx)
	err := checkEditPerm(ctx, p, userToken)

	if err != nil {
		return nil, err
	}
	p.CurrentCreative.URL = p.URL
	p.CurrentCreative.Status = orm.PendingCreativeStatus
	p.CurrentCreative.Type = orm.CreativeNativeType
	p.CurrentCreative.MaxBid = mysql.NullInt64{Valid: p.MaxBid != 0, Int64: p.MaxBid}
	p.CurrentCreative.Attributes = p.Attributes
	p.CurrentCreative.Name = p.Name
	err = p.CurrentCreative.SetAuditDomainID(p.CurrentDomain.ID)
	if err != nil {
		return nil, err
	}
	err = p.CurrentCreative.SetAuditOwnerID(p.CurrentCreative.UserID)
	if err != nil {
		return nil, err
	}
	d := structs.Map(p.CurrentCreative)
	err = p.CurrentCreative.SetAuditDescribe(d, "edit creative")
	if err != nil {
		return nil, err
	}
	err = p.CurrentCreative.SetAuditEntity("creative", p.CurrentCreative.ID)
	if err != nil {
		return nil, err
	}

	db := orm.NewOrmManager()
	assets := generateNativeAssets(p.Assets, p.Images, p.Icons, p.Logos, p.Videos)
	res, err := db.EditCreative(p.CurrentCreative, assets)
	if err != nil {
		return res, errors.DBError
	}
	return res, nil
}

func checkEditPerm(ctx context.Context, p *editNativePayload, userToken string) error {
	// check creative perm
	_, ok := p.CurrentUser.HasOn("edit_creative", p.CreativeOwner.ID, p.CurrentDomain.ID, false, false)
	if !ok {
		return errors.AccessDenied
	}
	// check campaign perm
	uScope, ok := p.CurrentUser.HasOn("edit_campaign", p.CampaignOwner.ID, p.CurrentDomain.ID, false, false)
	if !ok {
		return errors.AccessDenied
	}
	err := p.CurrentCreative.SetAuditUserData(p.CurrentUser.ID, userToken, p.CurrentDomain.ID, "edit_creative,edit_campaign", uScope)
	if err != nil {
		return err
	}

	return checkFilePerm(ctx, p)
}

func checkFilePerm(ctx context.Context, p *editNativePayload) error {

	var images = make([]*uploadOrm.Upload, 0)

	for i := range p.Assets.Images {
		img, err := validateImage("image", p.Assets.Images[i].Val)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, p.CurrentDomain.ID, p.CurrentUser); err != nil {
			return errors.AssetsPermErr
		}
		img.Label = p.Assets.Images[i].Label

		images = append(images, img)
	}

	p.Images = images

	var icons = make([]*uploadOrm.Upload, 0)

	for i := range p.Assets.Icons {
		img, err := validateImage("image", p.Assets.Images[i].Val)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, p.CurrentDomain.ID, p.CurrentUser); err != nil {
			return errors.AssetsPermErr
		}
		img.Label = p.Assets.Images[i].Label

		icons = append(icons, img)
	}

	p.Icons = icons

	var logos = make([]*uploadOrm.Upload, 0)

	for i := range p.Assets.Logos {
		img, err := validateImage("logo", p.Assets.Logos[i].Val)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, p.CurrentDomain.ID, p.CurrentUser); err != nil {
			return errors.AssetsPermErr
		}
		img.Label = p.Assets.Logos[i].Label

		logos = append(logos, img)
	}

	p.Logos = logos

	var videos = make([]*uploadOrm.Upload, 0)

	for i := range p.Assets.Videos {
		video, err := validateVideo(p.Assets.Videos[i].Val)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, video, p.CurrentDomain.ID, p.CurrentUser); err != nil {
			return errors.AssetsPermErr
		}
		video.Label = p.Assets.Videos[i].Label

		videos = append(videos, video)
	}

	p.Videos = videos
	return nil
}
