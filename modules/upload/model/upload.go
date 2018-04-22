package model

import (
	"encoding/json"
	"time"

	"database/sql/driver"

	"fmt"

	"clickyab.com/crab/modules/upload/errors"
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
	// VideoMime VideoMime
	VideoMime Mime = "video/mp4"
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
	Attr      FileAttr  `json:"attr" db:"attr"`

	Label string `json:"-" db:"-"`
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
	default:
		return errors.InvalidFileTypeError
	}

	return json.Unmarshal(c, b)
}

// Value for ad attr
func (b FileAttr) Value() (driver.Value, error) {
	return json.Marshal(b)
}

// FindSectionUploadByID return the Upload base on its id and section
func (m *Manager) FindSectionUploadByID(id, section string) (*Upload, error) {
	var res Upload
	err := m.GetRDbMap().SelectOne(
		&res,
		fmt.Sprintf("SELECT %s FROM %s WHERE id=? AND section=?", getSelectFields(UploadTableFull, ""), UploadTableFull),
		id,
		section,
	)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
