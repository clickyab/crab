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
func generateNativeAssets(data NativeAssetPayload, images []*model.Upload, icon, logo, video *model.Upload) []*orm.Asset {
	var assets []*orm.Asset
	var properties map[string]interface{}

	assets = append(assets,
		&orm.Asset{
			AssetType:  orm.AssetTextType,
			AssetKey:   "title",
			AssetValue: data.Title,
			Property:   map[string]interface{}{"len": utf8.RuneCountInString(data.Title)},
		},
		&orm.Asset{
			AssetType:  orm.AssetTextType,
			AssetKey:   "description",
			AssetValue: data.Description,
			Property:   map[string]interface{}{"len": utf8.RuneCountInString(data.Description)},
		},
		&orm.Asset{
			AssetType:  orm.AssetTextType,
			AssetKey:   "cta",
			AssetValue: data.CTA,
		},
	)

	if icon != nil {
		properties = map[string]interface{}{"width": icon.Attr.Native.Width, "height": icon.Attr.Native.Height}

		tmp := orm.Asset{
			AssetType:  orm.AssetImageType,
			Property:   properties,
			AssetKey:   "icon",
			AssetValue: icon.ID,
		}
		assets = append(assets, &tmp)
	}

	if len(images) > 0 {
		for _, img := range images {
			properties = map[string]interface{}{"width": img.Attr.Native.Width, "height": img.Attr.Native.Height}

			tmp := orm.Asset{
				AssetType:  orm.AssetImageType,
				Property:   properties,
				AssetKey:   "image",
				AssetValue: img.ID,
			}
			assets = append(assets, &tmp)
		}
	}

	if video != nil {

		tmp := orm.Asset{
			AssetType:  orm.AssetVideoType,
			AssetKey:   "video",
			AssetValue: video.ID,
		}
		assets = append(assets, &tmp)
	}

	if logo != nil {
		properties = map[string]interface{}{"width": logo.Attr.Native.Width, "height": logo.Attr.Native.Height}

		tmp := orm.Asset{
			AssetType:  orm.AssetImageType,
			Property:   properties,
			AssetKey:   "logo",
			AssetValue: logo.ID,
		}
		assets = append(assets, &tmp)
	}

	if data.Rating != 0 {

		tmp := orm.Asset{
			AssetType:  orm.AssetNumberType,
			AssetKey:   "rating",
			AssetValue: fmt.Sprintf("%v", data.Rating),
		}
		assets = append(assets, &tmp)
	}

	if data.Price != 0 {

		tmp := orm.Asset{
			AssetType:  orm.AssetNumberType,
			AssetKey:   "price",
			AssetValue: fmt.Sprintf("%v", data.Price),
		}
		assets = append(assets, &tmp)
	}

	if data.SalePrice != 0 {

		tmp := orm.Asset{
			AssetType:  orm.AssetNumberType,
			AssetKey:   "saleprice",
			AssetValue: fmt.Sprintf("%v", data.SalePrice),
		}
		assets = append(assets, &tmp)
	}

	if data.Downloads != 0 {

		tmp := orm.Asset{
			AssetType:  orm.AssetNumberType,
			AssetKey:   "downloads",
			AssetValue: fmt.Sprintf("%d", data.Downloads),
		}
		assets = append(assets, &tmp)
	}

	if data.Phone != "" {

		tmp := orm.Asset{
			AssetType:  orm.AssetNumberType,
			AssetKey:   "phone",
			AssetValue: data.Phone,
		}
		assets = append(assets, &tmp)
	}

	return assets
}
