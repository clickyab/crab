package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/ad/controllers/advalidator"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campignErr "clickyab.com/crab/modules/campaign/errors"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	dm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/safe"
	"github.com/fatih/structs"
)

var creativeSeed = config.RegisterBoolean("crab.modules.creative.seed", false, "insert detail after creative created")

// @Validate{
//}
type createNativePayload struct {
	CampaignID      int64                  `json:"campaign_id" validate:"required"`
	Name            string                 `json:"name" validate:"required"`
	URL             string                 `json:"url" validate:"required"`
	MaxBid          int64                  `json:"max_bid"`
	Attributes      map[string]interface{} `json:"attributes"`
	Assets          *NativeAssetPayload    `json:"assets"`
	CurrentUser     *aaa.User              `json:"-"`
	CurrentDomain   *dm.Domain             `json:"-"`
	CurrentCampaign *campaignOrm.Campaign  `json:"-"`
	CampaignOwner   *aaa.User              `json:"-"`
}

func (p *createNativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := p.Assets.Validate(ctx, w, r); err != nil {
		return err
	}
	err := emptyNativeValNotPermitted(p.Assets)
	if err != nil {
		return err
	}
	currentUser := authz.MustGetUser(ctx)
	p.CurrentUser = currentUser
	cpManager := campaignOrm.NewOrmManager()
	dmn := domain.MustGetDomain(ctx)
	p.CurrentDomain = dmn
	targetCampaign, err := cpManager.FindCampaignByIDDomain(p.CampaignID, dmn.ID)
	if err != nil {
		return campignErr.NotFoundError(p.CampaignID)
	}
	p.CurrentCampaign = targetCampaign
	campaignOwner, err := aaa.NewAaaManager().FindUserByID(targetCampaign.UserID)
	if err != nil {
		return campignErr.NotFoundError(targetCampaign.ID)
	}
	p.CampaignOwner = campaignOwner

	// extra fields required for app campaigns required
	if targetCampaign.Kind == campaignOrm.AppCampaign {
		if len(p.Assets.VImages) == 0 {
			return errors.VerticalImageRequiredErr
		}
		if len(p.Assets.HImages) == 0 {
			return errors.HorizontalImageRequiredErr
		}
		if len(p.Assets.Icons) == 0 {
			return errors.IconRequiredErr
		}
		if len(p.Assets.CTAs) == 0 {
			return errors.CtaRequiredErr
		}
	}

	return nil
}

// addNativeCreative to campaign
// @Rest {
// 		url = /native
//		protected = true
// 		method = post
// 		resource = create_creative:self
// }
func (c Controller) addNativeCreative(ctx context.Context, r *http.Request, p *createNativePayload) (*orm.CreativeSaveResult, error) {
	token := authz.MustGetToken(ctx)
	err := checkNativeCreatePerm(ctx, p)
	if err != nil {
		return nil, err
	}
	creative := &orm.Creative{
		URL:        p.URL,
		Status:     orm.PendingCreativeStatus,
		CampaignID: p.CurrentCampaign.ID,
		Type:       orm.CreativeNativeType,
		UserID:     p.CurrentUser.ID,
		MaxBid:     mysql.NullInt64{Valid: p.MaxBid != 0, Int64: p.MaxBid},
		Attributes: p.Attributes,
		Name:       p.Name,
	}
	err = creative.SetAuditDomainID(p.CurrentDomain.ID)
	if err != nil {
		return nil, err
	}
	err = creative.SetAuditOwnerID(p.CurrentUser.ID)
	if err != nil {
		return nil, err
	}
	d := structs.Map(creative)
	err = creative.SetAuditDescribe(d, "add native creative")
	if err != nil {
		return nil, err
	}
	err = creative.SetAuditUserData(p.CurrentUser.ID, token, p.CurrentDomain.ID, "add_creative", permission.ScopeSelf)
	if err != nil {
		return nil, err
	}
	db := orm.NewOrmManager()
	assets := generateNativeAssets(p.Assets, p.Assets.vImages, p.Assets.hImages, p.Assets.icons, p.Assets.logos, p.Assets.videos)
	res, err := db.AddCreative(creative, assets)
	if err != nil {
		return res, errors.DBError
	}
	// only in development
	if creativeSeed.Bool() {
		safe.GoRoutine(ctx, func() {
			Seed(creative)
		})
	}

	return res, nil
}

func init() {
	advalidator.RegisterAdValidationRules()
}
