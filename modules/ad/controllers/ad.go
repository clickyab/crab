package controllers

import (
	"fmt"

	"unicode/utf8"

	uploadOrm "clickyab.com/crab/modules/upload/model"

	"clickyab.com/crab/modules/ad/controllers/advalidator"
	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/upload/model"
	"github.com/clickyab/services/framework/controller"
)

// Controller is the controller for the ad package
// @Route {
//		group = /ad
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}

// Important: only add here shared func that use in other package routes

// generateNativeAssets generate a slice of native assets
func generateNativeAssets(data *NativeAssetPayload, vImages, hImages, icons, logos, videos []*model.Upload) []*orm.Asset {
	var assets []*orm.Asset

	assets = append(assets, generateNativeString(data.Titles, orm.AssetTextType, "title")...)

	if len(icons) > 0 {
		assets = append(assets, generateNativeMedia(icons, orm.AssetImageType, "icon")...)
	}
	if len(vImages) > 0 {
		assets = append(assets, generateNativeMedia(vImages, orm.AssetImageType, "v_image")...)
	}
	if len(hImages) > 0 {
		assets = append(assets, generateNativeMedia(hImages, orm.AssetImageType, "h_image")...)
	}
	if len(videos) > 0 {
		assets = append(assets, generateNativeMedia(videos, orm.AssetVideoType, "video")...)
	}
	if len(logos) > 0 {
		assets = append(assets, generateNativeMedia(logos, orm.AssetImageType, "logo")...)
	}
	if len(data.Ratings) > 0 {
		assets = append(assets, generateNativeFloat(data.Ratings, orm.AssetTextType, "rating")...)
	}
	if len(data.Prices) > 0 {
		assets = append(assets, generateNativeFloat(data.Prices, orm.AssetNumberType, "price")...)
	}
	if len(data.SalePrices) > 0 {
		assets = append(assets, generateNativeFloat(data.SalePrices, orm.AssetNumberType, "saleprice")...)
	}
	if len(data.Downloads) > 0 {
		assets = append(assets, generateNativeInt(data.Downloads, orm.AssetNumberType, "downloads")...)
	}

	if len(data.Descriptions) > 0 {
		assets = append(assets, generateNativeString(data.Descriptions, orm.AssetTextType, "description")...)
	}

	if len(data.CTAs) > 0 {
		assets = append(assets, generateNativeString(data.CTAs, orm.AssetTextType, "cta")...)
	}

	if len(data.Phones) > 0 {
		assets = append(assets, generateNativeString(data.Phones, orm.AssetNumberType, "phone")...)
	}

	return assets
}

func generateNativeFloat(assets []orm.NativeFloat, typ orm.AssetTypes, key string) []*orm.Asset {
	var res []*orm.Asset
	for _, val := range assets {
		tmp := map[string]interface{}{
			"label": val.Label,
		}
		res = append(res, &orm.Asset{
			AssetValue: fmt.Sprintf("%v", val.Val),
			AssetType:  typ,
			AssetKey:   key,
			Property:   tmp,
		})
	}
	return res
}

func generateNativeInt(assets []orm.NativeInt, typ orm.AssetTypes, key string) []*orm.Asset {
	var res []*orm.Asset
	for _, val := range assets {
		tmp := map[string]interface{}{
			"label": val.Label,
		}
		res = append(res, &orm.Asset{
			AssetValue: fmt.Sprintf("%d", val.Val),
			AssetType:  typ,
			AssetKey:   key,
			Property:   tmp,
		})
	}
	return res
}

func generateNativeString(assets []orm.NativeString, typ orm.AssetTypes, key string) []*orm.Asset {
	var res []*orm.Asset
	for _, val := range assets {
		tmp := map[string]interface{}{
			"len":   utf8.RuneCountInString(val.Val),
			"label": val.Label,
		}
		res = append(res, &orm.Asset{
			AssetValue: val.Val,
			AssetType:  typ,
			AssetKey:   key,
			Property:   tmp,
		})
	}
	return res
}

func generateNativeMedia(assets []*model.Upload, typ orm.AssetTypes, key string) []*orm.Asset {
	var res []*orm.Asset

	for _, val := range assets {
		tmp := orm.Asset{
			AssetType: typ,
			Property: func() map[string]interface{} {
				if key == "video" {
					return map[string]interface{}{"width": val.Attr.Video.Width, "height": val.Attr.Video.Height, "label": val.Label}
				}
				return map[string]interface{}{"width": val.Attr.Native.Width, "height": val.Attr.Native.Height, "label": val.Label}
			}(),
			AssetKey:   key,
			AssetValue: val.ID,
		}
		res = append(res, &tmp)
	}
	return res
}

func validateVastMedia(image string) (*uploadOrm.Upload, orm.AssetTypes, error) {
	uploadDBM := uploadOrm.NewModelManager()
	uploadFile, err := uploadDBM.FindUploadByID(image)
	if err != nil {
		return nil, "", errors.FileNotFound(image)
	}
	if uploadFile.Section != "vast-image" && uploadFile.Section != "vast-video" {
		return nil, "", errors.InvalidUploadSectionErr
	}
	if uploadFile.MIME == string(uploadOrm.VideoMime) {
		width, height := uploadFile.Attr.Video.Width, uploadFile.Attr.Video.Height
		input := advalidator.InputData{
			Width:    float64(width),
			Height:   float64(height),
			Size:     uploadFile.Size,
			Duration: int64(uploadFile.Attr.Video.Duration),
			Ext:      uploadFile.MIME,
		}
		rule := &advalidator.AdValidationConf.FWebVast.Video
		return uploadFile, orm.AssetVideoType, rule.Check(input)
	}
	if uploadFile.Attr.Banner == nil {
		return nil, orm.AssetImageType, errors.InvalideImageSize
	}
	width, height := uploadFile.Attr.Banner.Width, uploadFile.Attr.Banner.Height
	input := advalidator.InputData{
		Width:    float64(width),
		Height:   float64(height),
		Size:     uploadFile.Size,
		Duration: 0,
		Ext:      uploadFile.MIME,
	}
	rule := &advalidator.AdValidationConf.FWebVast.Image
	return uploadFile, orm.AssetImageType, rule.Check(input)
}

func checkVastAssetsPerm(p *createVastPayLoad) error {

	if p.Assets.Media.Val != "" {
		image, kind, err := validateVastMedia(p.Assets.Media.Val)
		if err != nil {
			return err
		}
		p.Media = image
		p.MediaKind = kind
	}
	return nil
}

func generateVastAssets(p *createVastPayLoad) []*orm.Asset {

	var assets []*orm.Asset
	if p.MediaKind == orm.AssetVideoType {
		assets = append(assets, generateVastMedia(p.Media, orm.AssetVideoType, "video"))
	}
	if p.MediaKind == orm.AssetImageType {
		assets = append(assets, generateVastMedia(p.Media, orm.AssetImageType, "image"))
	}
	if p.Assets.Cta.Val != "" {
		assets = append(assets, generateVastString(p.Assets.Cta, orm.AssetTextType, "cta"))
	}
	return assets
}

func generateVastMedia(asset *uploadOrm.Upload, typ orm.AssetTypes, key string) *orm.Asset {
	a := &orm.Asset{
		AssetType: typ,
		Property: func() map[string]interface{} {
			if key == "video" {
				return map[string]interface{}{"width": asset.Attr.Video.Width, "height": asset.Attr.Video.Height, "label": asset.Label}
			}
			return map[string]interface{}{"width": asset.Attr.Banner.Width, "height": asset.Attr.Banner.Height, "label": asset.Label}
		}(),
		AssetKey:   key,
		AssetValue: asset.ID,
	}
	return a
}

func generateVastString(asset orm.NativeString, typ orm.AssetTypes, key string) *orm.Asset {
	tmp := map[string]interface{}{
		"len":   utf8.RuneCountInString(asset.Val),
		"label": asset.Label,
	}
	res := &orm.Asset{
		AssetValue: asset.Val,
		AssetType:  typ,
		AssetKey:   key,
		Property:   tmp,
	}
	return res
}
