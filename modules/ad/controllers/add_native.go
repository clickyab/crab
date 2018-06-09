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
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/safe"
	"github.com/clickyab/services/xlog"
	"github.com/fatih/structs"
)

var creativeSeed = config.RegisterBoolean("crab.modules.creative.seed", true, "insert detail after creative created")

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
	Titles       []orm.NativeString `json:"title" validate:"required"`
	Descriptions []orm.NativeString `json:"description"`
	CTAs         []orm.NativeString `json:"cta"`
	//Optional assets
	Icons      []orm.NativeString `json:"icon" validate:"omitempty"`
	Images     []orm.NativeString `json:"images" validate:"omitempty"`
	Videos     []orm.NativeString `json:"video" validate:"omitempty"`
	Logos      []orm.NativeString `json:"logo" validate:"omitempty"`
	Ratings    []orm.NativeFloat  `json:"rating" validate:"omitempty"`
	Prices     []orm.NativeFloat  `json:"price" validate:"omitempty"`
	SalePrices []orm.NativeFloat  `json:"sale_price" validate:"omitempty"`
	Downloads  []orm.NativeInt    `json:"downloads" validate:"omitempty"`
	Phones     []orm.NativeString `json:"phone" validate:"omitempty"`
}

// @Validate{
//}
type createNativePayload struct {
	CampaignID      int64                  `json:"campaign_id" validate:"required"`
	Name            string                 `json:"name" validate:"required"`
	URL             string                 `json:"url" validate:"required"`
	MaxBid          int64                  `json:"max_bid"`
	Attributes      map[string]interface{} `json:"attributes"`
	Assets          NativeAssetPayload     `json:"assets"`
	CurrentUser     *aaa.User              `json:"-"`
	CurrentDomain   *dm.Domain             `json:"-"`
	CurrentCampaign *campaignOrm.Campaign  `json:"-"`
	CampaignOwner   *aaa.User              `json:"-"`
	Images          []*uploadOrm.Upload    `json:"-"`
	Logos           []*uploadOrm.Upload    `json:"-"`
	Videos          []*uploadOrm.Upload    `json:"-"`
	Icons           []*uploadOrm.Upload    `json:"-"`
}

func (p *createNativePayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
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
	targetCampaign, err := cpManager.FindCampaignByIDDomain(p.CampaignID, dmn.ID)
	if err != nil {
		return campignErr.NotFoundError(p.CampaignID)
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
			return errors.ImageRequiredErr
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

func validateImage(kind string, image string) (*uploadOrm.Upload, error) {
	var width, height int
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(image, "native-image")
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
	return uploadFile, nil
}

func validateVideo(video string) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(video, "native-video")
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
// 		resource = create_creative:self
// }
func (c Controller) addNativeCreative(ctx context.Context, r *http.Request, p *createNativePayload) (*orm.CreativeSaveResult, error) {
	err := checkCreatePerm(ctx, p)
	token := authz.MustGetToken(ctx)
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
	assets := generateNativeAssets(p.Assets, p.Images, p.Icons, p.Logos, p.Videos)
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

func checkCreatePerm(ctx context.Context, p *createNativePayload) error {
	// check campaign perm
	_, ok := p.CurrentUser.HasOn("edit_campaign", p.CampaignOwner.ID, p.CurrentDomain.ID, false, false, permission.ScopeSelf)
	if !ok {
		return errors.AccessDenied
	}

	var images = make([]*uploadOrm.Upload, 0)

	for i := range p.Assets.Images {
		img, err := validateImage("image", p.Assets.Images[i].Val)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, p.CurrentDomain.ID, p.CurrentUser); err != nil {
			return err
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
			return err
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
			return err
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
			return err
		}
		video.Label = p.Assets.Videos[i].Label

		videos = append(videos, video)
	}

	p.Videos = videos
	return nil
}

func emptyValNotPermitted(payload NativeAssetPayload) error {
	err := emptyStringNotAllowed(payload.Titles)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.Descriptions)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.Images)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.Phones)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.CTAs)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.Videos)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.Logos)
	if err != nil {
		return err
	}
	err = emptyStringNotAllowed(payload.Icons)
	if err != nil {
		return err
	}
	err = emptyIntNotAllowed(payload.Downloads)
	if err != nil {
		return err
	}
	err = emptyFloatNotAllowed(payload.Ratings)
	if err != nil {
		return err
	}
	err = emptyFloatNotAllowed(payload.Prices)
	if err != nil {
		return err
	}
	err = emptyFloatNotAllowed(payload.SalePrices)
	return err
}

func emptyStringNotAllowed(data []orm.NativeString) error {
	for _, v := range data {
		if v.Val == "" {
			return errors.EmptyValErr
		}
	}
	return nil
}

func emptyIntNotAllowed(data []orm.NativeInt) error {
	for _, v := range data {
		if v.Val == 0 {
			return errors.EmptyValErr
		}
	}
	return nil
}

func emptyFloatNotAllowed(data []orm.NativeFloat) error {
	for _, v := range data {
		if v.Val == 0 {
			return errors.EmptyValErr
		}
	}
	return nil
}

// fileOwnerCheckPerm owner check
func fileOwnerCheckPerm(ctx context.Context, image *uploadOrm.Upload, d int64, currentUser *aaa.User) error {
	targetFileOwner, err := aaa.NewAaaManager().FindUserWithParentsByID(image.UserID, d)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error on FindUserWithParentsByID")
		return errors.AssetsPermErr
	}
	// check campaign perm
	_, ok := currentUser.HasOn("edit_creative", targetFileOwner.ID, d, false, false, permission.ScopeSelf)
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
