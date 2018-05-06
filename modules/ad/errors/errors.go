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
	// AccessDenied have not access error
	AccessDenied = t9e.G("access Denied!you have not access to do this on creatives")
	// DBError to show when query has error
	DBError = t9e.G("database error! please check request data and try again")
	// EmptyValErr native asset should not have empty value
	EmptyValErr = t9e.G("native asset should not have empty value")
	// InvalideImageSize when width or height of image is not correct
	InvalideImageSize = t9e.G("invalid image size! please check request uploaded image width and height")

	// ImageRequiredErr image required in app native
	ImageRequiredErr = t9e.G("image required in app native")

	// VideoRequiredErr video required in app native
	VideoRequiredErr = t9e.G("video required in app native")

	// CtaRequiredErr cta required in app native
	CtaRequiredErr = t9e.G("cta required in app native")

	// IconRequiredErr icon required in app native
	IconRequiredErr = t9e.G("icon required in app native")

	// InvalidStatusErr invalid status error
	InvalidStatusErr = t9e.G("invalid status you can select %s or %s or %s.", "accepted", "rejected", "pending")

	// UpdateStatusDbErr error in update creative status
	UpdateStatusDbErr = t9e.G("an database error occurred when we try to update creative status ")
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

// FileNotFound error maker
func FileNotFound(name string) error {
	if name != "" {
		return t9e.G("%s file not found! please check your request data.", name)
	}

	return t9e.G("file not found! please check your request data.")

}
