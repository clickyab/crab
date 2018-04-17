package orm

import (
	"fmt"
	"reflect"
	"time"

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

// BeautyAsset make assets of a creative beaty and key=>value for response
func BeautyAsset(assets []Asset) map[string]interface{} {
	var bAssets = make(map[string]interface{})

	for _, asset := range assets {
		if _, ok := bAssets[asset.AssetKey]; ok {
			var tmp = make([]interface{}, 0)
			if reflect.TypeOf(bAssets[asset.AssetKey]).Kind() == reflect.Slice {
				tmp = append(tmp, interfaceSlice(bAssets[asset.AssetKey])...)
			} else {
				tmp = append(tmp, bAssets[asset.AssetKey])
			}
			tmp = append(tmp, asset.AssetValue)

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
