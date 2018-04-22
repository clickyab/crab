package orm

import (
	"fmt"
	"time"

	"strconv"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// AssetTypes is the creative type
type (
	// AssetTypes is the creative type
	// @Enum{
	// }
	AssetTypes string
)

const (
	// AssetImageType for image asset type
	AssetImageType AssetTypes = "image"
	// AssetVideoType for video asset type
	AssetVideoType AssetTypes = "video"
	// AssetTextType for text asset type
	AssetTextType AssetTypes = "text"
	// AssetNumberType for number asset type
	AssetNumberType AssetTypes = "number"
)

// Asset model in database
// @Model {
//		table = assets
//		primary = true, id
//		find_by = id,creative_id
//		list = yes
// }
type Asset struct {
	ID         int64                  `json:"id" db:"id"`
	CreativeID int64                  `json:"creative_id" db:"creative_id"`
	AssetType  AssetTypes             `json:"asset_type" db:"asset_type"`
	Property   mysql.GenericJSONField `json:"property" db:"property"`
	AssetKey   string                 `json:"asset_key" db:"asset_key"`
	AssetValue string                 `json:"asset_value" db:"asset_value"`
	CreatedAt  time.Time              `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time              `json:"updated_at" db:"updated_at"`
}

// DeleteAllCreativeAssets to delete all assets of a creative
func (m *Manager) DeleteAllCreativeAssets(crID int64) error {
	_, err := m.GetWDbMap().Exec(
		fmt.Sprintf("delete from %s "+
			"where creative_id=?",
			AssetTableFull,
		),
		crID,
	)

	return err
}

// NativeString native string payload
type NativeString struct {
	Label string `json:"label,omitempty"`
	Val   string `json:"val"`
}

// NativeFloat native float payload
type NativeFloat struct {
	Label string  `json:"label,omitempty"`
	Val   float64 `json:"val"`
}

// NativeInt native int payload
type NativeInt struct {
	Label string `json:"label,omitempty"`
	Val   int64  `json:"val"`
}

// beautyAsset make assets of a creative beaty and key=>value for response
func beautyAsset(assets []Asset) map[string][]interface{} {
	var bAssets = make(map[string][]interface{})
	for _, asset := range assets {

		typ := asset.AssetType
		if typ == AssetNumberType {
			val, err := strconv.ParseFloat(asset.AssetValue, 10)
			assert.Nil(err)
			tmp := NativeFloat{
				Val:   val,
				Label: asset.Property["label"].(string),
			}
			bAssets[asset.AssetKey] = append(bAssets[asset.AssetKey], tmp)
		} else {
			tmp := NativeString{
				Val:   asset.AssetValue,
				Label: asset.Property["label"].(string),
			}
			bAssets[asset.AssetKey] = append(bAssets[asset.AssetKey], tmp)
		}

	}

	return bAssets
}

// beautyAsset make assets of a creative beaty and key=>value for response
//func beautyAsset1(assets []Asset) map[string]interface{} {
//	var bAssets = make(map[string]interface{})
//
//	var imgAsset = make([]NativeString, 0)
//	var logoAsset = make([]NativeString, 0)
//	var videoAsset = make([]NativeString, 0)
//	var iconAsset = make([]NativeString, 0)
//	var titleAsset = make([]NativeString, 0)
//	var describAsset = make([]NativeString, 0)
//	var ctaAsset = make([]NativeString, 0)
//	var ratingAsset = make([]NativeFloat, 0)
//	var priceAsset = make([]NativeFloat, 0)
//	var salePriceAsset = make([]NativeFloat, 0)
//	var downloadAsset = make([]NativeInt, 0)
//	var phoneAsset = make([]NativeString, 0)
//	for _, asset := range assets {
//
//		switch asset.AssetKey {
//		case "image":
//			imgAsset = append(imgAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "video":
//			videoAsset = append(videoAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "logo":
//			logoAsset = append(logoAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "icon":
//			iconAsset = append(iconAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "title":
//			titleAsset = append(titleAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "description":
//			describAsset = append(describAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "cta":
//			ctaAsset = append(ctaAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "rating":
//			val, _ := strconv.ParseFloat(asset.AssetValue, 10)
//			ratingAsset = append(ratingAsset, NativeFloat{
//				Val:   val,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "price":
//			val, _ := strconv.ParseFloat(asset.AssetValue, 10)
//			priceAsset = append(priceAsset, NativeFloat{
//				Val:   val,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "saleprice":
//			val, _ := strconv.ParseFloat(asset.AssetValue, 10)
//			salePriceAsset = append(salePriceAsset, NativeFloat{
//				Val:   val,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "downloads":
//			val, _ := strconv.ParseInt(asset.AssetValue, 10, 0)
//			downloadAsset = append(downloadAsset, NativeInt{
//				Val:   val,
//				Label: asset.Property["label"].(string),
//			})
//
//		case "phone":
//			phoneAsset = append(phoneAsset, NativeString{
//				Val:   asset.AssetValue,
//				Label: asset.Property["label"].(string),
//			})
//
//		}
//	}
//
//	if len(imgAsset) > 0 {
//		bAssets["image"] = imgAsset
//	}
//	if len(videoAsset) > 0 {
//		bAssets["video"] = videoAsset
//	}
//	if len(logoAsset) > 0 {
//		bAssets["logo"] = logoAsset
//	}
//	if len(iconAsset) > 0 {
//		bAssets["icon"] = iconAsset
//	}
//	if len(titleAsset) > 0 {
//		bAssets["title"] = titleAsset
//	}
//	if len(describAsset) > 0 {
//		bAssets["description"] = describAsset
//	}
//	if len(ctaAsset) > 0 {
//		bAssets["cta"] = ctaAsset
//	}
//	if len(ratingAsset) > 0 {
//		bAssets["rating"] = ratingAsset
//	}
//	if len(priceAsset) > 0 {
//		bAssets["price"] = priceAsset
//	}
//	if len(salePriceAsset) > 0 {
//		bAssets["saleprice"] = salePriceAsset
//	}
//	if len(downloadAsset) > 0 {
//		bAssets["downloads"] = downloadAsset
//	}
//	if len(phoneAsset) > 0 {
//		bAssets["phone"] = phoneAsset
//	}
//
//	return bAssets
//}
