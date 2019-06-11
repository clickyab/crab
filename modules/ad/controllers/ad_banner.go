package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campignErr "clickyab.com/crab/modules/campaign/errors"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/safe"
	"github.com/fatih/structs"
)

// BannerAssetPayload BannerAssetPayload
// @Validate{
//}
type BannerAssetPayload struct {
	CTA         orm.CreativeString `json:"cta" validate:"omitempty"`
	VideoImage  orm.CreativeString `json:"video_image" validate:"omitempty"`
	Video       orm.CreativeString `json:"video" validate:"omitempty"`
	BannerImage orm.CreativeString `json:"banner_image" validate:"omitempty"`

	bannerImage *uploadOrm.Upload `json:"-"`
	videoImage  *uploadOrm.Upload `json:"-"`
	video       *uploadOrm.Upload `json:"-"`
}

// CreateBannerCreative CreateBannerCreative
// @Validate{
//}
type CreateBannerCreative struct {
	Name       string                 `json:"name" validate:"required"`
	URL        string                 `json:"url" validate:"required"`
	MaxBid     int64                  `json:"max_bid"`
	Attributes map[string]interface{} `json:"attributes"`
	Assets     *BannerAssetPayload    `json:"assets"`
}

// createBannerPayload createBannerPayload
// @Validate{
//}
type createBannerPayload struct {
	CampaignID int64                  `json:"campaign_id" validate:"required"`
	Creative   []CreateBannerCreative `json:"creative" validate:"required"`
	Type       BannerType             `json:"banner_type" validate:"required"`

	currentUser     *aaa.User             `json:"-"`
	currentDomain   *domainOrm.Domain     `json:"-"`
	currentCampaign *campaignOrm.Campaign `json:"-"`
	campaignOwner   *aaa.User             `json:"-"`
}

// BannerType is the banner creative type
type (
	// BannerType is the banner type video/banner
	// @Enum{
	// }
	BannerType string
)

const (
	// BannerImageType for image banner
	BannerImageType BannerType = "image"
	// BannerVideoType for video banner
	BannerVideoType BannerType = "video"
)

func (p *createBannerPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Type.IsValid() {
		return errors.InvalidBannerTypeErr
	}
	for i := range p.Creative {
		if err := p.Creative[i].Validate(ctx, w, r); err != nil {
			return err
		}
		if err := p.Creative[i].Assets.Validate(ctx, w, r); err != nil {
			return err
		}
		// check by request type image/video
		if p.Type == BannerImageType {
			if (orm.CreativeString{}) == p.Creative[i].Assets.BannerImage {
				return errors.InvalidBannerImage
			}

		} else { // video banner selected
			if (orm.CreativeString{}) == p.Creative[i].Assets.Video {
				return errors.InvalidBannerVideo
			}
			if (orm.CreativeString{}) == p.Creative[i].Assets.CTA {
				return errors.InvalidCTA
			}
		}
	}
	currentUser := authz.MustGetUser(ctx)
	p.currentUser = currentUser
	cpManager := campaignOrm.NewOrmManager()
	dmn := domain.MustGetDomain(ctx)
	p.currentDomain = dmn
	targetCampaign, err := cpManager.FindCampaignByIDDomain(p.CampaignID, dmn.ID)
	if err != nil {
		return campignErr.NotFoundError(p.CampaignID)
	}
	p.currentCampaign = targetCampaign
	campaignOwner, err := aaa.NewAaaManager().FindUserByID(targetCampaign.UserID)
	if err != nil {
		return campignErr.NotFoundError(targetCampaign.ID)
	}
	p.campaignOwner = campaignOwner
	return nil
}

type createBannerResponse []orm.CreativeSaveResult

// addBannerCreative to campaign
// @Rest {
// 		url = /banner
//		protected = true
// 		method = post
// 		resource = create_creative:self
// }
func (c Controller) addBannerCreative(ctx context.Context, r *http.Request, p *createBannerPayload) (*createBannerResponse, error) {
	token := authz.MustGetToken(ctx)
	grantedScope, err := checkBannerCreatePerm(ctx, p)
	if err != nil {
		return nil, err
	}
	var multiCreative = make([]orm.MultiCreative, 0)
	for i := range p.Creative {
		creative := &orm.Creative{
			UserID:     p.currentUser.ID,
			CampaignID: p.currentCampaign.ID,
			Status:     orm.PendingCreativeStatus,
			Type:       orm.CreativeBannerType,
			URL:        p.Creative[i].URL,
			Name:       p.Creative[i].Name,
			MaxBid:     mysql.NullInt64{Valid: p.Creative[i].MaxBid != 0, Int64: p.Creative[i].MaxBid},
			Attributes: p.Creative[i].Attributes,
		}
		multiCreative = append(multiCreative, orm.MultiCreative{
			Creative: creative,
			Assets:   generateBannerAssets(p.Creative[i].Assets, p.Creative[i].Assets.videoImage, p.Creative[i].Assets.video, p.Creative[i].Assets.bannerImage),
		})
		err = creative.SetAuditDomainID(p.currentDomain.ID)
		if err != nil {
			return nil, err
		}
		err = creative.SetAuditOwnerID(p.currentUser.ID)
		if err != nil {
			return nil, err
		}
		d := structs.Map(creative)
		err = creative.SetAuditDescribe(d, "add banner creative")
		if err != nil {
			return nil, err
		}
		err = creative.SetAuditUserData(p.currentUser.ID, token, p.currentDomain.ID, "create_creative", grantedScope)
		if err != nil {
			return nil, err
		}
	}
	db := orm.NewOrmManager()
	res, err := db.AddMultiCreative(multiCreative)
	if err != nil {
		return nil, errors.DBError
	}
	// only in development
	if creativeSeed.Bool() {
		for i := range res {
			safe.GoRoutine(ctx, func() {
				Seed(res[i].Creative)
			})
		}
	}
	finalRes := createBannerResponse(res)
	return &finalRes, nil
}
