package controllers

import (
	campaignErr "clickyab.com/crab/modules/campaign/errors"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	uploadOrm "clickyab.com/crab/modules/upload/model"

	"context"
	"net/http"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/fatih/structs"
)

// @Validate{
//}
type vastAssetPayload struct {
	Media orm.VastString `json:"media" validate:"omitempty"`
	Cta   orm.VastString `json:"cta" validate:"omitempty"`
}

// @Validate{
//}
type createVastPayLoad struct {
	CampaignID int64                  `json:"campaign_id" validate:"required"`
	Name       string                 `json:"name" validate:"required"`
	URL        string                 `json:"url" validate:"required"`
	MaxBid     int64                  `json:"max_bid"`
	Attributes map[string]interface{} `json:"attributes"`

	Assets vastAssetPayload `json:"assets" validate:"required"`

	media     *uploadOrm.Upload `json:"-"`
	mediaKind orm.AssetTypes    `json:"-"`

	currentUser    *aaa.User             `json:"-"`
	currentDomain  *domainOrm.Domain     `json:"-"`
	targetCampaign *campaignOrm.Campaign `json:"-"`
	campaignOwner  *aaa.User             `json:"-"`
}

func (p *createVastPayLoad) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	currentUser := authz.MustGetUser(ctx)
	p.currentUser = currentUser
	cpManager := campaignOrm.NewOrmManager()
	dmn := domain.MustGetDomain(ctx)
	p.currentDomain = dmn

	targetCampaign, err := cpManager.FindCampaignByIDDomain(p.CampaignID, dmn.ID)
	if err != nil {
		return campaignErr.NotFoundError(targetCampaign.ID)
	}

	p.targetCampaign = targetCampaign
	campaignOwner, err := aaa.NewAaaManager().FindUserByID(targetCampaign.UserID)
	if err != nil {
		return campaignErr.NotFoundError(p.targetCampaign.UserID)
	}
	p.campaignOwner = campaignOwner

	// extra fields required for app campaigns required
	if targetCampaign.Kind == campaignOrm.AppCampaign {
		// TODO : implement app vast
		return errors.NotImplementedVastAppErr
	}
	return nil
}

// addVastCreative to campaign
// @Rest {
// 		url = /vast
//		protected = true
// 		method = post
// 		resource = create_creative:self
// }
func (c Controller) addVastCreative(ctx context.Context, r *http.Request, p *createVastPayLoad) (*orm.CreativeSaveResult, error) {

	token := authz.MustGetToken(ctx)
	// check permission
	uScope, ok := p.currentUser.HasOn("create_creative", p.campaignOwner.ID, p.currentDomain.ID, false, false, permission.ScopeSelf, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	// check file perm and validate theme
	err := checkVastAssetsPerm(p)
	if err != nil {
		return nil, err
	}
	creative := &orm.Creative{
		URL:        p.URL,
		Status:     orm.PendingCreativeStatus,
		CampaignID: p.targetCampaign.ID,
		Type:       orm.CreativeVastType,
		UserID:     p.currentUser.ID,
		MaxBid:     mysql.NullInt64{Valid: p.MaxBid != 0, Int64: p.MaxBid},
		Attributes: p.Attributes,
		Name:       p.Name,
	}
	err = creative.SetAuditDomainID(p.currentDomain.ID)
	if err != nil {
		return nil, err
	}
	err = creative.SetAuditOwnerID(p.currentUser.ID)
	if err != nil {
		return nil, err
	}
	d := structs.Map(creative)
	err = creative.SetAuditDescribe(d, "add vast creative")
	if err != nil {
		return nil, err
	}
	err = creative.SetAuditUserData(p.currentUser.ID, token, p.currentDomain.ID, "create_creative", uScope)
	if err != nil {
		return nil, err
	}
	assets := generateVastAssets(p)
	db := orm.NewOrmManager()
	res, err := db.AddCreative(creative, assets)
	if err != nil {
		return nil, errors.DBError
	}
	return res, nil
}
