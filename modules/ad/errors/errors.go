package errors

import (
	"github.com/clickyab/services/gettext/t9e"
)

var (
	// InvalidIDErr for all invalid id errors
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// UnsupportTypeError for unsupported type errors
	UnsupportTypeError = t9e.G("unsupported type.")
	// NoBannerError error when no banner selected
	NoBannerError = t9e.G("no banner selected.")
	// TypeError for banner type errors
	TypeError = t9e.G("invalid ad banner type. you can select %s or %s or %s please check your request data and try again", "banner", "video", "native")
	// BannerInvalidCampaignError banner campaign is not your selected campaign error
	BannerInvalidCampaignError = t9e.G("banner campaign is not your selected campaign.")
	// InvalidUploadedFile invalid banner uploaded file error
	InvalidUploadedFile = t9e.G("invalid banner uploaded file.")
	// InvalidDimension invalid file dimension error
	InvalidDimension = t9e.G("invalid file dimension.")
	// DuplicateAdSrc duplicate src error
	DuplicateAdSrc = t9e.G("duplicate src. add banner file is duplicate")
)

// AdNotFound maker
func AdNotFound(id int64) error {
	if id > 0 {
		return t9e.G("ad with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("ad not found, please check your request data.")
}

// BannerNotFound maker
func BannerNotFound(id int64) error {
	if id > 0 {
		return t9e.G("ad with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("ad not found, please check your request data.")
}