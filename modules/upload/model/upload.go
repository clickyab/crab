package model

import (
	"encoding/json"
	"errors"
	"time"

	"database/sql/driver"
)

// Mime all mime type
type Mime string

const (
	// JPGMime jpeg image
	JPGMime Mime = "image/jpeg"
	// PJPGMime pjpeg image
	PJPGMime Mime = "image/pjpeg"
	// PNGMime PNGMime
	PNGMime Mime = "image/png"
	// GifMime GifMime
	GifMime Mime = "image/gif"
)

// Upload model in database
// @Model {
//		table = uploads
//		primary = false, id
//		find_by = id
//		list = yes
// }
type Upload struct {
	ID        string    `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at"  db:"created_at"`
	MIME      string    `json:"mime"  db:"mime"`
	Size      int64     `json:"size"  db:"size"`
	UserID    int64     `json:"user_id"  db:"user_id"`
	Section   string    `json:"section" db:"section"`
	Attr      FileAttr  `json:"attr"`
}

// BannerAttr banner ad type attr
type BannerAttr struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// AvatarAttr avatar  type attr
type AvatarAttr struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// VideoAttr video ad type attr
type VideoAttr struct {
	Duration int `json:"duration"`
}

// NativeAttr native ad type attr
type NativeAttr struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// FileAttr ad attribute
type FileAttr struct {
	Avatar *AvatarAttr `json:"avatar,omitempty"`
	Banner *BannerAttr `json:"banner,omitempty"`
	Video  *VideoAttr  `json:"video,omitempty"`
	Native *NativeAttr `json:"native,omitempty"`
}

// Scan for add attr
func (b *FileAttr) Scan(src interface{}) error {
	var c []byte
	switch src.(type) {
	case []byte:
		c = src.([]byte)
	case string:
		c = []byte(src.(string))
	case nil:
		c = make([]byte, 0)
		return nil
	default:
		return errors.New("unsupported type")
	}

	return json.Unmarshal(c, b)
}

// Value for ad attr
func (b FileAttr) Value() (driver.Value, error) {
	if b.Banner != nil {
		b.Native = nil
		b.Video = nil
		b.Avatar = nil
	} else if b.Native != nil {
		b.Banner = nil
		b.Video = nil
		b.Avatar = nil
	} else if b.Video != nil {
		b.Banner = nil
		b.Native = nil
		b.Avatar = nil
	} else if b.Avatar != nil {
		b.Native = nil
		b.Video = nil
		b.Banner = nil
	} else {
		b.Native = nil
		b.Video = nil
		b.Banner = nil
		b.Avatar = nil
	}
	return json.Marshal(b)
}
