package orm

import (
	"fmt"
	"time"
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
	ID         int64      `json:"id" db:"id"`
	CreativeID int64      `json:"creative_id" db:"creative_id"`
	AssetType  AssetTypes `json:"asset_type" db:"asset_type"`
	Property   string     `json:"property" db:"property"`
	AssetKey   string     `json:"asset_key" db:"asset_key"`
	AssetValue string     `json:"asset_value" db:"asset_value"`
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at" db:"updated_at"`
}

// NativeAssets list of valid assets for native
type NativeAssets struct {
	//Required assets
	Title       string `json:"title" validation:"required"`
	Description string `json:"description" validation:"required"`
	CTA         string `json:"cta" validation:"required"`
	//Optional assets
	Icon      string  `json:"icon" validation:"omitempty"`
	Image     string  `json:"image" validation:"omitempty"`
	Video     string  `json:"video" validation:"omitempty"`
	Logo      string  `json:"logo" validation:"omitempty"`
	Rating    float64 `json:"rating" validation:"omitempty"`
	Price     float64 `json:"price" validation:"omitempty"`
	Saleprice float64 `json:"saleprice" validation:"omitempty"`
	Downloads int64   `json:"downloads" validation:"omitempty"`
	Phone     string  `json:"phone" validation:"omitempty"`
}

// GenerateNativeAssets generate a slice of native assets
func GenerateNativeAssets(data NativeAssets) []Asset {
	var assets []Asset

	tmp := Asset{
		AssetType:  AssetTextType,
		Property:   "",
		AssetKey:   "title",
		AssetValue: data.Title,
		CreatedAt:  time.Now(),
	}
	assets = append(assets, tmp)

	tmp = Asset{
		AssetType:  AssetTextType,
		Property:   "",
		AssetKey:   "description",
		AssetValue: data.Description,
		CreatedAt:  time.Now(),
	}
	assets = append(assets, tmp)

	tmp = Asset{
		AssetType:  AssetTextType,
		Property:   "",
		AssetKey:   "cta",
		AssetValue: data.CTA,
		CreatedAt:  time.Now(),
	}
	assets = append(assets, tmp)

	if data.Icon != "" {
		tmp = Asset{
			AssetType:  AssetImageType,
			Property:   "",
			AssetKey:   "icon",
			AssetValue: data.Icon,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Image != "" {
		tmp = Asset{
			AssetType:  AssetImageType,
			Property:   "",
			AssetKey:   "image",
			AssetValue: data.Image,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Video != "" {
		tmp = Asset{
			AssetType:  AssetVideoType,
			Property:   "",
			AssetKey:   "video",
			AssetValue: data.Video,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Logo != "" {
		tmp = Asset{
			AssetType:  AssetImageType,
			Property:   "",
			AssetKey:   "logo",
			AssetValue: data.Logo,
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Rating != 0 {
		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   "",
			AssetKey:   "rating",
			AssetValue: fmt.Sprintf("%v", data.Rating),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Price != 0 {
		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   "",
			AssetKey:   "price",
			AssetValue: fmt.Sprintf("%v", data.Price),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Saleprice != 0 {
		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   "",
			AssetKey:   "saleprice",
			AssetValue: fmt.Sprintf("%v", data.Saleprice),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Downloads != 0 {
		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   "",
			AssetKey:   "downloads",
			AssetValue: fmt.Sprintf("%d", data.Downloads),
			CreatedAt:  time.Now(),
		}
		assets = append(assets, tmp)
	}

	if data.Phone != "" {
		tmp = Asset{
			AssetType:  AssetNumberType,
			Property:   "",
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
