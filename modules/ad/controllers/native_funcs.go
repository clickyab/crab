package controllers

import (
	"context"

	"clickyab.com/crab/modules/ad/controllers/advalidator"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	dm "clickyab.com/crab/modules/domain/orm"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/permission"
)

const (
	nativeVerticalImage   = "v-image"
	nativeHorizontalImage = "h-image"
	nativeLogo            = "logo"
	nativeIcon            = "icon"
)

// NativeAssetPayload native asset
// @Validate{
//}
type NativeAssetPayload struct {
	//Required assets
	Titles       []orm.CreativeString `json:"title" validate:"required"`
	Descriptions []orm.CreativeString `json:"description"`
	CTAs         []orm.CreativeString `json:"cta"`
	//Optional assets
	Icons      []orm.CreativeString `json:"icon" validate:"omitempty"`
	HImages    []orm.CreativeString `json:"h_image" validate:"omitempty"`
	VImages    []orm.CreativeString `json:"v_image" validate:"omitempty"`
	Videos     []orm.CreativeString `json:"video" validate:"omitempty"`
	Logos      []orm.CreativeString `json:"logo" validate:"omitempty"`
	Ratings    []orm.NativeFloat    `json:"rating" validate:"omitempty"`
	Prices     []orm.NativeFloat    `json:"price" validate:"omitempty"`
	SalePrices []orm.NativeFloat    `json:"sale_price" validate:"omitempty"`
	Downloads  []orm.NativeInt      `json:"downloads" validate:"omitempty"`
	Phones     []orm.CreativeString `json:"phone" validate:"omitempty"`

	vImages []*uploadOrm.Upload `json:"-"`
	hImages []*uploadOrm.Upload `json:"-"`
	logos   []*uploadOrm.Upload `json:"-"`
	videos  []*uploadOrm.Upload `json:"-"`
	icons   []*uploadOrm.Upload `json:"-"`
}

func emptyNativeValNotPermitted(payload *NativeAssetPayload) error {
	err := emptyNativeStringNotAllowed(payload.Titles)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.Descriptions)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.VImages)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.HImages)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.Phones)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.CTAs)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.Videos)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.Logos)
	if err != nil {
		return err
	}
	err = emptyNativeStringNotAllowed(payload.Icons)
	if err != nil {
		return err
	}
	err = emptyNativeIntNotAllowed(payload.Downloads)
	if err != nil {
		return err
	}
	err = emptyNativeFloatNotAllowed(payload.Ratings)
	if err != nil {
		return err
	}
	err = emptyNativeFloatNotAllowed(payload.Prices)
	if err != nil {
		return err
	}
	err = emptyNativeFloatNotAllowed(payload.SalePrices)
	return err
}

func emptyNativeStringNotAllowed(data []orm.CreativeString) error {
	for _, v := range data {
		if v.Val == "" {
			return errors.EmptyValErr
		}
	}
	return nil
}

func emptyNativeIntNotAllowed(data []orm.NativeInt) error {
	for _, v := range data {
		if v.Val == 0 {
			return errors.EmptyValErr
		}
	}
	return nil
}

func emptyNativeFloatNotAllowed(data []orm.NativeFloat) error {
	for _, v := range data {
		if v.Val == 0 {
			return errors.EmptyValErr
		}
	}
	return nil
}

func checkNativeCreatePerm(ctx context.Context, p *createNativePayload) error {
	// check campaign perm
	_, ok := p.CurrentUser.HasOn("edit_campaign", p.CampaignOwner.ID, p.CurrentDomain.ID, false, false, permission.ScopeSelf, permission.ScopeGlobal)
	if !ok {
		return errors.AccessDenied
	}

	return validateCreateNativeFiles(ctx, p)
}

func validateCreateNativeFiles(ctx context.Context, p *createNativePayload) error {
	err := validateNativeImages(ctx, p.Assets, p.CurrentCampaign, p.CurrentDomain, p.CurrentUser)
	if err != nil {
		return err
	}
	return validateNativeVideos(ctx, p.Assets, p.CurrentCampaign, p.CurrentDomain, p.CurrentUser)

}

func validateNativeImages(ctx context.Context, assets *NativeAssetPayload, campaign *campaignOrm.Campaign, domain *dm.Domain, user *aaa.User) error {
	var vImages = make([]*uploadOrm.Upload, 0)

	for i := range assets.VImages {
		img, err := validateNativeImage(nativeVerticalImage, assets.VImages[i].Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, domain.ID, user); err != nil {
			return err
		}
		img.Label = assets.VImages[i].Label

		vImages = append(vImages, img)
	}

	assets.vImages = vImages

	var hImages = make([]*uploadOrm.Upload, 0)

	for i := range assets.HImages {
		img, err := validateNativeImage(nativeHorizontalImage, assets.HImages[i].Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, domain.ID, user); err != nil {
			return err
		}
		img.Label = assets.HImages[i].Label

		hImages = append(hImages, img)
	}

	assets.hImages = hImages

	var icons = make([]*uploadOrm.Upload, 0)

	for i := range assets.Icons {
		img, err := validateNativeImage(nativeIcon, assets.Icons[i].Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, domain.ID, user); err != nil {
			return err
		}
		img.Label = assets.Icons[i].Label

		icons = append(icons, img)
	}

	assets.icons = icons

	var logos = make([]*uploadOrm.Upload, 0)

	for i := range assets.Logos {
		img, err := validateNativeImage(nativeLogo, assets.Logos[i].Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, domain.ID, user); err != nil {
			return err
		}
		img.Label = assets.Logos[i].Label

		logos = append(logos, img)
	}

	assets.logos = logos
	return nil
}

func validateNativeVideos(ctx context.Context, assets *NativeAssetPayload, campaign *campaignOrm.Campaign, domain *dm.Domain, user *aaa.User) error {
	var videos = make([]*uploadOrm.Upload, 0)
	for i := range assets.Videos {
		video, err := validateNativeVideo(assets.Videos[i].Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, video, domain.ID, user); err != nil {
			return err
		}
		video.Label = assets.Videos[i].Label

		videos = append(videos, video)
	}

	assets.videos = videos
	return nil
}

func validateNativeVideo(video string, kind campaignOrm.CampaignKind) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(video, "native-video")
	if err != nil {
		return nil, errors.FileNotFound("video")
	}

	input := advalidator.InputData{
		Size:     uploadFile.Size,
		Duration: int64(uploadFile.Attr.Video.Duration),
		Ext:      uploadFile.MIME,
	}

	if kind == campaignOrm.AppCampaign {
		rule := &advalidator.AdValidationConf.FAppNative.Video
		return uploadFile, rule.Check(input)
	}
	rule := &advalidator.AdValidationConf.FWebNative.Video
	return uploadFile, rule.Check(input)
}

func validateNativeImg(input advalidator.InputData, kind string, campaignKind campaignOrm.CampaignKind) error {
	if campaignKind == campaignOrm.AppCampaign {
		switch kind {
		case nativeVerticalImage:
			rule := &advalidator.AdValidationConf.FAppNative.VImage
			return rule.Check(input)
		case nativeIcon:
			rule := &advalidator.AdValidationConf.FAppNative.Icon
			return rule.Check(input)
		case nativeHorizontalImage:
			rule := &advalidator.AdValidationConf.FAppNative.HImage
			return rule.Check(input)
		case nativeLogo:
			rule := &advalidator.AdValidationConf.FAppNative.Logo
			return rule.Check(input)

		}
	} else {
		switch kind {
		case nativeVerticalImage:
			rule := &advalidator.AdValidationConf.FWebNative.VImage
			return rule.Check(input)
		case nativeIcon:
			rule := &advalidator.AdValidationConf.FWebNative.Icon
			return rule.Check(input)
		case nativeHorizontalImage:
			rule := &advalidator.AdValidationConf.FWebNative.HImage
			return rule.Check(input)
		case nativeLogo:
			rule := &advalidator.AdValidationConf.FWebNative.Logo
			return rule.Check(input)

		}
	}
	return errors.InvalidUploadedFile
}

func validateNativeImage(kind string, image string, campaignKind campaignOrm.CampaignKind) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(image, "native-image")
	if err != nil {
		return nil, errors.FileNotFound("image")
	}
	if uploadFile.Attr.Native == nil {
		return nil, errors.InvalidImageSize
	}
	if kind == nativeVerticalImage || kind == nativeIcon || kind == nativeHorizontalImage || kind == nativeLogo {
		width, height := uploadFile.Attr.Native.Width, uploadFile.Attr.Native.Height
		input := advalidator.InputData{
			Width:    float64(width),
			Height:   float64(height),
			Size:     uploadFile.Size,
			Duration: 0,
			Ext:      uploadFile.MIME,
		}
		return uploadFile, validateNativeImg(input, kind, campaignKind)

	}
	return uploadFile, errors.InvalidUploadedFile
}

func validateEditNativeFiles(ctx context.Context, p *editNativePayload) error {
	err := validateNativeImages(ctx, p.Assets, p.CurrentCampaign, p.CurrentDomain, p.CurrentUser)
	if err != nil {
		return err
	}
	return validateNativeVideos(ctx, p.Assets, p.CurrentCampaign, p.CurrentDomain, p.CurrentUser)

}

func checkNativeEditPerm(ctx context.Context, p *editNativePayload, userToken string) error {
	// check creative perm
	_, ok := p.CurrentUser.HasOn("edit_creative", p.CreativeOwner.ID, p.CurrentDomain.ID, false, false, permission.ScopeSelf, permission.ScopeGlobal)
	if !ok {
		return errors.AccessDenied
	}
	// check campaign perm
	uScope, ok := p.CurrentUser.HasOn("edit_campaign", p.CampaignOwner.ID, p.CurrentDomain.ID, false, false, permission.ScopeSelf, permission.ScopeGlobal)
	if !ok {
		return errors.AccessDenied
	}
	err := p.CurrentCreative.SetAuditUserData(p.CurrentUser.ID, userToken, p.CurrentDomain.ID, "edit_creative,edit_campaign", uScope)
	if err != nil {
		return err
	}

	return validateEditNativeFiles(ctx, p)
}
