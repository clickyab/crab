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
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

/* type bannerSize struct {
	Width  int
	Height int
} */

/* var (
	bannerValidSizes = config.RegisterString("crab.modules.ad.banner",
		"120x600,160x600,300x250,336x280,468x60,728x90,120x240,320x50,800x440,300x600,970x90,970x250,250x250,300x1050,320x480,480x320,128x128",
		"valid banner sizes separated by x",
	)
	minNativeBannerRation = config.RegisterFloat64("crab.modules.ad.native.ratio.min", 1.55, "minimum native banner ration")
	maxNativeBannerRation = config.RegisterFloat64("crab.modules.ad.native.ratio.max", 1.65, "maximum native banner ration")
	lock                  = sync.Mutex{}

	bannerValid []bannerSize
) */

// @Validate{
//}
type nativeCreativePayload struct {
	Creative orm.BaseCreativeData `json:"creative"`
	Assets   orm.NativeAssets     `json:"assets"`
}

func (p *nativeCreativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// TODO: check native images ratio
	return nil
}

// addNativeCreative to campaign
// @Rest {
// 		url = /native
//		protected = true
// 		method = post
// }
func (c Controller) addNativeCreative(ctx context.Context, r *http.Request, p *nativeCreativePayload) (*orm.CreativeSaveResult, error) {
	cpManager := campaignOrm.NewOrmManager()
	d := domain.MustGetDomain(ctx)
	ca, err := cpManager.FindCampaignByIDDomain(p.Creative.CampaignID, d.ID)
	if err != nil {
		return nil, campignErr.NotFoundError(d.ID)
	}

	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ca.UserID, d.ID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "add_creative", d.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	p.Creative.Status = orm.PendingCreativeStatus
	p.Creative.Type = orm.CreativeNativeType

	db := orm.NewOrmManager()
	assets := orm.GenerateNativeAssets(p.Assets)
	res, err := db.AddCreative(p.Creative, assets)
	if err != nil {
		return &res, errors.DBError
	}

	return &res, nil
}

// editNativeCreative to campaign
// @Rest {
// 		url = /native/:creative_id
//		protected = true
// 		method = put
// }
func (c Controller) editNativeCreative(ctx context.Context, r *http.Request, p *nativeCreativePayload) (*orm.CreativeSaveResult, error) {
	id, err := strconv.ParseInt(xmux.Param(ctx, "creative_id"), 10, 64)
	if err != nil {
		return nil, errors.InvalidIDErr
	}
	p.Creative.ID = id

	cpManager := campaignOrm.NewOrmManager()
	d := domain.MustGetDomain(ctx)
	ca, err := cpManager.FindCampaignByIDDomain(p.Creative.CampaignID, d.ID)
	if err != nil {
		return nil, campignErr.NotFoundError(d.ID)
	}

	currentUser := authz.MustGetUser(ctx)
	userManager := aaa.NewAaaManager()
	owner, err := userManager.FindUserWithParentsByID(ca.UserID, d.ID)
	assert.Nil(err)
	_, ok := aaa.CheckPermOn(owner, currentUser, "edit_creative", d.ID)
	if !ok {
		return nil, errors.AccessDenied
	}

	if p.Creative.ID == 0 {
		return nil, errors.InvalidIDErr
	}

	db := orm.NewOrmManager()
	cr, err := db.FindCreativeByIDAndType(p.Creative.ID, orm.CreativeNativeType)
	if cr.ID == 0 || err != nil {
		return nil, t9e.G("Creative with identifier %s for you and type %s is not a valid creative. check creative exist or is a native?", p.Creative.ID, orm.CreativeNativeType)
	}

	if p.Creative.Status == "" {
		p.Creative.Status = cr.Status
	}
	p.Creative.Type = orm.CreativeNativeType

	assets := orm.GenerateNativeAssets(p.Assets)
	res, err := db.EditCreative(p.Creative, assets)
	if err != nil {
		return &res, errors.DBError
	}

	return &res, nil
}

/* func checkBannerDimension(width, height int, bannerType orm.CreativeTypes) bool {
	switch bannerType {
	case orm.CreativeBannerType:
		for i := range bannerValid {
			if bannerValid[i].Width == width && bannerValid[i].Height == height {
				return true
			}
		}
	case orm.CreativeNativeType:
		rate := float64(float64(width) / float64(height))
		if rate < maxNativeBannerRation.Float64() && rate > minNativeBannerRation.Float64() {
			return true
		}

	}

	return false
}

func checkBannerImage(srcID string, bannerType orm.CreativeTypes) (mime string, width int, height int, err error) {
	file, err := model.NewModelManager().FindUploadByID(srcID)

	if err != nil || (file.Attr.Banner == nil && file.Attr.Native == nil) {
		err = errors.InvalidUploadedFile
		return
	}
	//check banner type
	if file.Section != string(bannerType) {
		err = errors.NoBannerError
		return
	}
	//TODO check access to file
	mime = file.MIME

	switch bannerType {
	case orm.CreativeBannerType:
		width = file.Attr.Banner.Width
		height = file.Attr.Banner.Height
	case orm.CreativeNativeType:
		width = file.Attr.Native.Width
		height = file.Attr.Native.Height
	}
	ok := checkBannerDimension(width, height, bannerType)
	if !ok {
		err = errors.InvalidDimension
		return
	}
	return
} */

/* func init() {
	lock.Lock()
	defer lock.Unlock()
	// extract valid sizes
	mainArr := strings.Split(bannerValidSizes.String(), ",")
	for i := range mainArr {
		size := bannerSize{}
		sizeArr := strings.Split(mainArr[i], "x")
		assert.True(len(sizeArr) == 2)
		width, err := strconv.Atoi(sizeArr[0])
		assert.Nil(err)
		height, err := strconv.Atoi(sizeArr[1])
		assert.Nil(err)
		size.Width = width
		size.Height = height
		bannerValid = append(bannerValid, size)
	}
}
*/
