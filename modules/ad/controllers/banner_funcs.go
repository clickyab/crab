package controllers

import (
	"context"
	"unicode/utf8"

	"clickyab.com/crab/modules/ad/controllers/advalidator"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	campaignOrm "clickyab.com/crab/modules/campaign/orm"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	uploadOrm "clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/permission"
)

func validateBannerImage(image string, campaignKind campaignOrm.CampaignKind) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(image, "banner-image")
	if err != nil {
		return nil, errors.FileNotFound("image")
	}
	if uploadFile.Attr.Banner == nil {
		return nil, errors.InvalidImageSize
	}

	width, height := uploadFile.Attr.Banner.Width, uploadFile.Attr.Banner.Height
	input := advalidator.InputData{
		Width:    float64(width),
		Height:   float64(height),
		Size:     uploadFile.Size,
		Duration: 0,
		Ext:      uploadFile.MIME,
	}
	return uploadFile, validateBannerImg(input, campaignKind)
}

func validateVideoImage(image string, campaignKind campaignOrm.CampaignKind) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(image, "banner-image")
	if err != nil {
		return nil, errors.FileNotFound("banner-video-image")
	}
	if uploadFile.Attr.Banner == nil {
		return nil, errors.InvalidImageSize
	}

	width, height := uploadFile.Attr.Banner.Width, uploadFile.Attr.Banner.Height
	input := advalidator.InputData{
		Width:    float64(width),
		Height:   float64(height),
		Size:     uploadFile.Size,
		Duration: 0,
		Ext:      uploadFile.MIME,
	}
	return uploadFile, validateVideoImg(input, campaignKind)
}

func validateBannerImg(input advalidator.InputData, campaignKind campaignOrm.CampaignKind) error {
	if campaignKind == campaignOrm.AppCampaign {
		rule := &advalidator.AdValidationConf.FAppBanner.Image
		return rule.Check(input)
	}
	rule := &advalidator.AdValidationConf.FWebBanner.Image
	return rule.Check(input)

}

func validateVideoImg(input advalidator.InputData, campaignKind campaignOrm.CampaignKind) error {
	if campaignKind == campaignOrm.AppCampaign {
		rule := &advalidator.AdValidationConf.FAppBanner.VideoImage
		return rule.Check(input)
	}
	rule := &advalidator.AdValidationConf.FWebBanner.VideoImage
	return rule.Check(input)

}

func validateBannerImages(ctx context.Context, assets *BannerAssetPayload, campaign *campaignOrm.Campaign, domain *domainOrm.Domain, user *aaa.User) error {

	if (orm.CreativeString{}) != assets.BannerImage {
		img, err := validateBannerImage(assets.BannerImage.Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, domain.ID, user); err != nil {
			return err
		}
		img.Label = assets.BannerImage.Label

		assets.bannerImage = img
	}

	if (orm.CreativeString{}) != assets.VideoImage {
		img, err := validateVideoImage(assets.VideoImage.Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, img, domain.ID, user); err != nil {
			return err
		}
		img.Label = assets.VideoImage.Label

		assets.videoImage = img
	}

	return nil
}

func validateCreateBannerFiles(ctx context.Context, p *CreateBannerCreative, campaign *campaignOrm.Campaign, domain *domainOrm.Domain, user *aaa.User) error {
	err := validateBannerImages(ctx, p.Assets, campaign, domain, user)
	if err != nil {
		return err
	}
	return validateBannerVideos(ctx, p.Assets, campaign, domain, user)

}

func validateBannerVideo(video string, kind campaignOrm.CampaignKind) (*uploadOrm.Upload, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindSectionUploadByID(video, "banner-video")
	if err != nil {
		return nil, errors.FileNotFound("banner-video")
	}

	input := advalidator.InputData{
		Size:     uploadFile.Size,
		Duration: int64(uploadFile.Attr.Video.Duration),
		Ext:      uploadFile.MIME,
	}

	if uploadFile.Attr.Video == nil {
		return uploadFile, errors.InvalidImageSize
	}

	if kind == campaignOrm.AppCampaign {
		rule := &advalidator.AdValidationConf.FAppBanner.Video
		return uploadFile, rule.Check(input)
	}
	rule := &advalidator.AdValidationConf.FWebBanner.Video
	return uploadFile, rule.Check(input)
}

func validateBannerVideos(ctx context.Context, assets *BannerAssetPayload, campaign *campaignOrm.Campaign, domain *domainOrm.Domain, user *aaa.User) error {
	if (orm.CreativeString{}) != assets.Video {
		video, err := validateBannerVideo(assets.Video.Val, campaign.Kind)
		if err != nil {
			return err
		}
		if err := fileOwnerCheckPerm(ctx, video, domain.ID, user); err != nil {
			return err
		}
		video.Label = assets.Video.Label
		assets.video = video
	}

	return nil
}

func checkBannerCreatePerm(ctx context.Context, p *createBannerPayload) error {
	// check campaign perm
	_, ok := p.currentUser.HasOn("edit_campaign", p.campaignOwner.ID, p.currentDomain.ID, false, false, permission.ScopeSelf, permission.ScopeGlobal)
	if !ok {
		return errors.AccessDenied
	}

	for i := range p.Creative {
		if err := validateCreateBannerFiles(ctx, p.Creative[i], p.currentCampaign, p.currentDomain, p.currentUser); err != nil {
			return err
		}
	}
	return nil
}

func generateBannerString(asset orm.CreativeString, typ orm.AssetTypes, key string) *orm.Asset {
	tmp := map[string]interface{}{
		"len":   utf8.RuneCountInString(asset.Val),
		"label": asset.Label,
	}
	return &orm.Asset{
		AssetValue: asset.Val,
		AssetType:  typ,
		AssetKey:   key,
		Property:   tmp,
	}
}

// generateBannerAssets generate a slice of banner assets
func generateBannerAssets(data *BannerAssetPayload, videoImage, video, BannerImage *uploadOrm.Upload) []*orm.Asset {
	var assets []*orm.Asset

	if videoImage != nil {
		assets = append(assets, generateBannerMedia(videoImage, orm.AssetImageType, "banner_video_image"))
	}
	if video != nil {
		assets = append(assets, generateBannerMedia(video, orm.AssetVideoType, "banner_video"))
	}
	if BannerImage != nil {
		assets = append(assets, generateBannerMedia(BannerImage, orm.AssetImageType, "banner_image"))
	}

	if (orm.CreativeString{}) != data.CTA {
		assets = append(assets, generateBannerString(data.CTA, orm.AssetTextType, "cta"))
	}
	return assets
}

func generateBannerMedia(asset *uploadOrm.Upload, typ orm.AssetTypes, key string) *orm.Asset {
	return &orm.Asset{
		AssetType: typ,
		Property: func() map[string]interface{} {
			if key == "banner_video" {
				return map[string]interface{}{"label": asset.Label, "duration": asset.Attr.Video.Duration}
			}
			return map[string]interface{}{"width": asset.Attr.Banner.Width, "height": asset.Attr.Banner.Height, "label": asset.Label}
		}(),
		AssetKey:   key,
		AssetValue: asset.ID,
	}
}
