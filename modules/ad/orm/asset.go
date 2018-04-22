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

// FindAssetsBeautyByCreativeID find asset and beautify assets by creative id
func (m *Manager) FindAssetsBeautyByCreativeID(id int64) map[string][]interface{} {
	// find creative assets
	assets := m.ListAssetsWithFilter("creative_id=?", id)
	return beautyAsset(assets)
}
