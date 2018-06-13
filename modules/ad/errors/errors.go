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
	// InvalidImageSize when width or height of image is not correct
	InvalidImageSize = t9e.G("invalid image dimension db! please check request uploaded image width and height")
	// InvalidVideoSize when width or height of image is not correct
	InvalidVideoSize = t9e.G("invalid video size! please check request uploaded video width and height")

	// VerticalImageRequiredErr image required in app native
	VerticalImageRequiredErr = t9e.G("vertical image required in app native")

	// HorizontalImageRequiredErr image required in app native
	HorizontalImageRequiredErr = t9e.G("horizontal image required in app native")

	// VideoRequiredErr video required in app native
	VideoRequiredErr = t9e.G("video required in app native")

	// CtaRequiredErr cta required in app native
	CtaRequiredErr = t9e.G("cta required in app native")

	// IconRequiredErr icon required in app native
	IconRequiredErr = t9e.G("icon required in app native")

	// InvalidStatusErr invalid status error
	InvalidStatusErr = t9e.G("invalid status you can select accepted or rejected or pending.")

	// UpdateStatusDbErr error in update creative status
	UpdateStatusDbErr = t9e.G("an database error occurred when we try to update creative status ")

	// InvalidNotifyUser error in approve, reject creatives if Notify user value is invalid
	InvalidNotifyUser = t9e.G("notify user value is invalid, you can select yes or no")

	// AssetsPermErr permission error assets
	AssetsPermErr = t9e.G("Access denied! you don't have permission to read assets")

	//SendNotifyEmailErr error when notify user for creative status change
	SendNotifyEmailErr = t9e.G("error in sending notify email")

	// CreativeNotFoundErr creative not found error
	CreativeNotFoundErr = t9e.G("creative not found")
	// InvalidBannerTypeErr InvalidBannerTypeErr
	InvalidBannerTypeErr = t9e.G("banner type should be image/video")
	// InvalidBannerImage InvalidBannerImage
	InvalidBannerImage = t9e.G("banner image required")
	// InvalidBannerVideo InvalidBannerVideo
	InvalidBannerVideo = t9e.G("banner video required")
	// InvalidCTA InvalidCTA
	InvalidCTA = t9e.G("cta required")
	//NotImplementedVastAppErr vast app not implemented error
	NotImplementedVastAppErr = t9e.G("vast app not implemented error")

	//InvalidUploadSectionErr InvalidUploadSectionErr
	InvalidUploadSectionErr = t9e.G("upload section is not valid for this creative type")
	// ArchiveCreativeDbErr error when archive creative from db
	ArchiveCreativeDbErr = t9e.G("database error when archive creative")
	// AlReadyArchivedErr creative that you try to archive is already archived
	AlReadyArchivedErr = t9e.G("creative that you try to archive is already archived")
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
