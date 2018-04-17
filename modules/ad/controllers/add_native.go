package controllers

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campignErr "clickyab.com/crab/modules/campaign/errors"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	dm "clickyab.com/crab/modules/domain/orm"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
)

var (
	nativeValidImagesSizeConf = config.RegisterString("crab.modules.ad.native.images",
		"image:320x480,image:480x320,icon:512x512",
		"valid images size separated by x",
	)

	lock            = sync.Mutex{}
	validImageSizes []imageSize
)

type imageSize struct {
	width  int
	height int
	kind   string // weather its icon or image
}

// NativeAssetPayload native asset
// @Validate{
//}
type NativeAssetPayload struct {
	//Required assets
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CTA         string `json:"cta" validate:"required"`
	//Optional assets
	Icon      string   `json:"icon" validate:"omitempty"`
	Images    []string `json:"images" validate:"omitempty"`
	Video     string   `json:"video" validate:"omitempty"`
	Logo      string   `json:"logo" validate:"omitempty"`
	Rating    float64  `json:"rating" validate:"omitempty"`
	Price     float64  `json:"price" validate:"omitempty"`
	SalePrice float64  `json:"sale_price" validate:"omitempty"`
	Downloads int64    `json:"downloads" validate:"omitempty"`
	Phone     string   `json:"phone" validate:"omitempty"`
}

// @Validate{
//}
type createNativePayload struct {
	CampaignID      int64                 `json:"campaign_id" validate:"required"`
	URL             string                `json:"url" validate:"required"`
	MaxBid          int64                 `json:"max_bid" validate:"required,gt=0"`
	Assets          NativeAssetPayload    `json:"assets"`
	CurrentUser     *aaa.User             `json:"-"`
	CurrentDomain   *dm.Domain            `json:"-"`
	CurrentCampaign *campaignOrm.Campaign `json:"-"`
	CampaignOwner   *aaa.User             `json:"-"`
	Images          []*uploadOrm.Upload   `json:"-"`
	Logo            *uploadOrm.Upload     `json:"-"`
	Video           *uploadOrm.Upload     `json:"-"`
	Icon            *uploadOrm.Upload     `json:"-"`
}

func (p *createNativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if err := p.Assets.Validate(ctx, w, r); err != nil {
		return err
	}
	currentUser := authz.MustGetUser(ctx)
	p.CurrentUser = currentUser
	cpManager := campaignOrm.NewOrmManager()
	dmn := domain.MustGetDomain(ctx)
	p.CurrentDomain = dmn
	targetCampaign, err := cpManager.FindCampaignByIDDomain(p.CampaignID, dmn.ID)
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

func validateImage(kind string, images ...string) ([]*uploadOrm.Upload, error) {
	var width, height int
	uploadDBM := uploadOrm.NewModelManager()
	var res = make([]*uploadOrm.Upload, 0)
	for _, img := range images {
		uploadFile, err := uploadDBM.FindSectionUploadByID(img, "native")
		if err != nil {
			return nil, errors.FileNotFound("image")
		}
		if uploadFile.Attr.Native == nil {
			return nil, errors.InvalideImageSize
		}

		if kind == "image" || kind == "icon" {
			width, height = uploadFile.Attr.Native.Width, uploadFile.Attr.Native.Height
			if !isValidSize(width, height, kind) {
				return nil, errors.InvalideImageSize
			}
		}
		res = append(res, uploadFile)
	}
	return res, nil
}

func validateVideo(video string) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(video, "video")
	if err != nil {
		return nil, errors.FileNotFound("video")
	}

	//TODO : add extra validation for video later like duration or size ...
	return uploadFile, nil
}

func isValidSize(width, height int, kind string) bool {
	for i := range validImageSizes {
		if validImageSizes[i].width == width && validImageSizes[i].height == height && validImageSizes[i].kind == kind {
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
// 		resource = add_creative:self
// }
func (c Controller) addNativeCreative(ctx context.Context, r *http.Request, p *createNativePayload) (*orm.CreativeSaveResult, error) {
	err := checkCreatePerm(p)
	if err != nil {
		return nil, err
	}
	creative := &orm.Creative{
		URL:        p.URL,
		Status:     orm.PendingCreativeStatus,
		CampaignID: p.CurrentCampaign.ID,
		Type:       orm.CreativeNativeType,
		UserID:     p.CurrentUser.ID,
		MaxBid:     p.MaxBid,
	}

	db := orm.NewOrmManager()
	assets := generateNativeAssets(p.Assets, p.Images, p.Icon, p.Logo, p.Video)
	res, err := db.AddCreative(creative, assets)
	if err != nil {
		return res, errors.DBError
	}

	return res, nil
}

func checkCreatePerm(p *createNativePayload) error {
	// check campaign perm
	_, ok := aaa.CheckPermOn(p.CampaignOwner, p.CurrentUser, "edit_campaign", p.CurrentDomain.ID)
	if !ok {
		return errors.AccessDenied
	}

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

// fileOwnerCheckPerm owner check
func fileOwnerCheckPerm(image *uploadOrm.Upload, d int64, currentUser *aaa.User) error {
	targetFileOwner, err := aaa.NewAaaManager().FindUserWithParentsByID(image.UserID, d)
	if err != nil {
		return err
	}
	// check campaign perm
	_, ok := aaa.CheckPermOn(targetFileOwner, currentUser, "edit_creative", d)
	if !ok {
		return errors.AccessDenied
	}
	return nil
}

func init() {
	lock.Lock()
	defer lock.Unlock()
	// extract valid sizes

	sizesArr := strings.Split(nativeValidImagesSizeConf.String(), ",")

	for i := range sizesArr {
		size := imageSize{}
		sizeArr := strings.Split(sizesArr[i], ":")
		assert.True(len(sizeArr) == 2)
		kind := sizeArr[0]
		widthHeight := strings.Split(sizeArr[1], "x")
		assert.True(len(widthHeight) == 2)
		width, err := strconv.Atoi(widthHeight[0])
		assert.Nil(err)
		height, err := strconv.Atoi(widthHeight[1])
		assert.Nil(err)
		size.width = width
		size.height = height
		size.kind = kind
		validImageSizes = append(validImageSizes, size)
	}

}
