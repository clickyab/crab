package controllers

import (
	"fmt"

	"unicode/utf8"

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
func generateNativeAssets(data NativeAssetPayload, images, icons, logos, videos []*model.Upload) []*orm.Asset {
	var assets []*orm.Asset

	assets = append(assets, generateNativeString(data.Titles, orm.AssetTextType, "title")...)

	if len(icons) > 0 {
		assets = append(assets, generateNativeMedia(icons, orm.AssetImageType, "icon")...)
	}
	if len(images) > 0 {
		assets = append(assets, generateNativeMedia(images, orm.AssetImageType, "image")...)
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
					return map[string]interface{}{"label": val.Label}
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
