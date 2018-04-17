package orm

import (
	"fmt"
	"reflect"
	"time"

	"clickyab.com/crab/libs"
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

// NativeAssets list of valid assets for native
type NativeAssets struct {
	//Required assets
	Title       string `json:"title" validation:"required"`
	Description string `json:"description" validation:"required"`
	CTA         string `json:"cta" validation:"required"`
	//Optional assets
	Icon      string   `json:"icon" validation:"omitempty"`
	Images    []string `json:"images" validation:"omitempty"`
	Video     string   `json:"video" validation:"omitempty"`
	Logo      string   `json:"logo" validation:"omitempty"`
	Rating    float64  `json:"rating" validation:"omitempty"`
	Price     float64  `json:"price" validation:"omitempty"`
	Saleprice float64  `json:"saleprice" validation:"omitempty"`
	Downloads int64    `json:"downloads" validation:"omitempty"`
	Phone     string   `json:"phone" validation:"omitempty"`
}

// GenerateNativeAssets generate a slice of native assets
func GenerateNativeAssets(data NativeAssets, baseUploadPath string) []Asset {
	var assets []Asset
	var properties map[string]interface{}
	var width, height int
	var imgPath string

	tmp := Asset{
		AssetType:  AssetTextType,
		Property:   properties,
		AssetKey:   "title",
		AssetValue: data.Title,
		CreatedAt:  time.Now(),
	}
	assets = append(assets, tmp)

	tmp = Asset{
		AssetType:  AssetTextType,
		Property:   properties,
		AssetKey:   "description",
		AssetValue: data.Description,
		CreatedAt:  time.Now(),
	}
	assets = append(assets, tmp)

	tmp = Asset{
		AssetType:  AssetTextType,
		Property:   properties,
		AssetKey:   "cta",
		AssetValue: data.CTA,
		CreatedAt:  time.Now(),
	}
	assets = append(assets, tmp)

	if data.Icon != "" {
		imgPath = baseUploadPath + data.Icon
		width, height = libs.GetImageDimension(imgPath)
		properties = map[string]interface{}{"width": width, "height": height}

		tmp = Asset{
			AssetType:  AssetImageType,
			Property:   properties,
			AssetKey:   "icon",
			AssetValue: data.Icon,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if len(data.Images) > 0 {
		for _, img := range data.Images {
			imgPath = baseUploadPath + img
			width, height = libs.GetImageDimension(imgPath)
			properties = map[string]interface{}{"width": width, "height": height}

			tmp = Asset{
				AssetType:  AssetImageType,
				Property:   properties,
				AssetKey:   "image",
				AssetValue: img,
				CreatedAt:  time.Now(),
			}
			assets = append(assets, tmp)
		}
	}

	if data.Video != "" {
		properties = map[string]interface{}{}

		tmp = Asset{
			AssetType:  AssetVideoType,
			Property:   properties,
			AssetKey:   "video",
			AssetValue: data.Video,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Logo != "" {
		imgPath = baseUploadPath + data.Logo
		width, height = libs.GetImageDimension(imgPath)
		properties = map[string]interface{}{"width": width, "height": height}

		tmp = Asset{
			AssetType:  AssetImageType,
			Property:   properties,
			AssetKey:   "logo",
			AssetValue: data.Logo,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Rating != 0 {
		properties = map[string]interface{}{}

		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   properties,
			AssetKey:   "rating",
			AssetValue: fmt.Sprintf("%v", data.Rating),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Price != 0 {
		properties = map[string]interface{}{}

		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   properties,
			AssetKey:   "price",
			AssetValue: fmt.Sprintf("%v", data.Price),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Saleprice != 0 {
		properties = map[string]interface{}{}

		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   properties,
			AssetKey:   "saleprice",
			AssetValue: fmt.Sprintf("%v", data.Saleprice),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Downloads != 0 {
		properties = map[string]interface{}{}

		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   properties,
			AssetKey:   "downloads",
			AssetValue: fmt.Sprintf("%d", data.Downloads),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Phone != "" {
		properties = map[string]interface{}{}

		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   properties,
			AssetKey:   "phone",
			AssetValue: data.Phone,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	return assets
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

// BeautyAsset make assets of a creative beaty and key=>value for response
func BeautyAsset(assets []Asset) map[string]interface{} {
	bAssets := map[string]interface{}{}

	for _, asset := range assets {
		if _, ok := bAssets[asset.AssetKey]; ok {
			var tmp []interface{}
			if reflect.TypeOf(bAssets[asset.AssetKey]).Kind() == reflect.Slice {
				tmp = append(tmp, interfaceSlice(bAssets[asset.AssetKey])...)
			} else {
				tmp = append(tmp, bAssets[asset.AssetKey])
			}
			tmp = append(tmp, map[string]interface{}{asset.AssetKey: asset.AssetValue})

			bAssets[asset.AssetKey] = tmp
		} else {
			bAssets[asset.AssetKey] = asset.AssetValue
		}
	}

	return bAssets
}

func interfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}
