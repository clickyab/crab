package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	upload "clickyab.com/crab/modules/upload/controllers"

	"clickyab.com/crab/libs"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campignErr "clickyab.com/crab/modules/campaign/errors"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/gettext/t9e"
	"github.com/rs/xmux"
)

var (
	nativeValidImagesSizeConf = config.RegisterString("crab.modules.ad.native.images",
		"320x480,480x320",
		"valid images size separated by x",
	)
	nativeIconSize = config.RegisterString("crab.modules.ad.native.icon",
		"512x512",
		"valid icon size separated by x",
	)

	lock            = sync.Mutex{}
	validImageSizes []imageSize
)

type imageSize struct {
	Width  int
	Height int
}

// @Validate{
//}
type nativeCreativePayload struct {
	Creative orm.BaseCreativeData `json:"creative"`
	Assets   orm.NativeAssets     `json:"assets"`
}

func (p *nativeCreativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if len(p.Assets.Images) > 0 {
		err := validateImage(p.Assets.Images)
		if err != nil {
			return err
		}
	}

	if p.Assets.Icon != "" {
		icon := upload.UPath.String() + p.Assets.Icon
		width, height := libs.GetImageDimension(icon)

		if size := fmt.Sprint("%sx%s", width, height); size != nativeIconSize.String() {
			return errors.InvalideImageSize
		}

		uploadDBM := uploadOrm.NewModelManager()
		_, err := uploadDBM.FindUploadByID(p.Assets.Icon)
		if err != nil {
			return errors.FileNotFound("icon")
		}
	}

	return nil
}

func validateImage(images []string) error {
	var width, height int
	var imgPath string

	uploadDBM := uploadOrm.NewModelManager()

	for _, img := range images {
		imgPath = upload.UPath.String() + img
		width, height = libs.GetImageDimension(imgPath)

		if !isValidSize(width, height) {
			return errors.InvalideImageSize
		}

		_, err := uploadDBM.FindUploadByID(img)
		if err != nil {
			return errors.FileNotFound("image")
		}
	}

	return nil
}

func isValidSize(width, height int) bool {
	for i := range validImageSizes {
		if validImageSizes[i].Width == width && validImageSizes[i].Height == height {
			return true
		}
	}

	return false
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
	assets := orm.GenerateNativeAssets(p.Assets, upload.UPath.String())
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

	assets := orm.GenerateNativeAssets(p.Assets, upload.UPath.String())
	res, err := db.EditCreative(p.Creative, assets)
	if err != nil {
		return &res, errors.DBError
	}

	return &res, nil
}

func init() {
	lock.Lock()
	defer lock.Unlock()
	// extract valid sizes
	sizes := strings.Split(nativeValidImagesSizeConf.String(), ",")
	for i := range sizes {
		size := imageSize{}
		widthHeight := strings.Split(sizes[i], "x")
		assert.True(len(widthHeight) == 2)
		width, err := strconv.Atoi(widthHeight[0])
		assert.Nil(err)
		height, err := strconv.Atoi(widthHeight[1])
		assert.Nil(err)
		size.Width = width
		size.Height = height
		validImageSizes = append(validImageSizes, size)
	}
}
