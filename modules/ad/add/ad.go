package add

import "time"

// AdActiveStatus is the ad active status
type (
	// AdActiveStatus is the ad active status
	// @Enum{
	// }
	AdActiveStatus string
)

const (
	// RejectedAdStatus ad rejected status
	RejectedAdStatus AdActiveStatus = "rejected"
	// AcceptedAdStatus ad accepted status
	AcceptedAdStatus AdActiveStatus = "accepted"
	// PendingAdStatus ad pending status
	PendingAdStatus AdActiveStatus = "pending"
)

// AdType is the ad type
type (
	// AdType is the ad type
	// @Enum{
	// }
	AdType string
)

const (
	// BannerAdType banner
	BannerAdType AdType = "banner"
	// VideoAdType video
	VideoAdType AdType = "video"
	// NativeAdType native
	NativeAdType AdType = "native"
)

// Ad Ad model in database
// @Model {
//		table = ads
//		primary = true, id
//		find_by = id
//		list = yes
// }
type Ad struct {
	ID         int64          `json:"id" db:"id"`
	CampaignID int64          `json:"campaign_id" db:"campaign_id"`
	Src        string         `json:"src" db:"src"`
	Mime       string         `json:"mime" db:"mime"`
	Target     string         `json:"target" db:"target"`
	Width      int64          `json:"width" db:"width"`
	Height     int64          `json:"height" db:"height"`
	Status     AdActiveStatus `json:"status" db:"status"`
	Type       AdType         `json:"type" db:"type"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at" db:"updated_at"`
}
