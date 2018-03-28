package errors

import (
	"github.com/clickyab/services/gettext/t9e"
)

var (
	// InvalidIDErr invalid id, error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// UnsupportTypeError unsupported type. error
	UnsupportTypeError = t9e.G("unsupported type.")
	// TypeError invalid campaign type error
	TypeError = t9e.G("invalid campaign type. you can select %s or %s or %s please check your request data and try again", "banner", "vast", "native")
	// KindError invalid campaign kind error
	KindError = t9e.G("invalid campaign kind. you can select %s or %s or %s please check your request data and try again", "web", "app")
	// InvalidCampaignStatusError invalid campaign status error
	InvalidCampaignStatusError = t9e.G("invalid campaign status. you can select %s or %s or %s please check your request data and try again", "archive", "start", "pause")
	// EndTimeError campaign should end error
	EndTimeError = t9e.G("campaign should end after start")
	// TimeScheduleError at least you error
	TimeScheduleError = t9e.G("at least you should specify one schedule time for campaign")
	// ArchivedEditError can't error
	ArchivedEditError = t9e.G("can't update or copy archived campaign")
	// DuplicateNameError duplicate campaign name error
	DuplicateNameError = t9e.G("duplicate campaign name. campaign name should be a unique")
	// StartTimeError start date error
	StartTimeError error = t9e.G("start date can't be past")
	// ErrInventoryID there is no inventory with this id
	ErrInventoryID error = t9e.G("there is no inventory with this id")
	// AccessDenied error
	AccessDenied error = t9e.G("access denied! you don't have access for this action")
)

// NotFoundError maker
func NotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("campaign with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("campaign not found, please check your request data.")
}

// InvalidError maker
func InvalidError(dataName string) error {
	return t9e.G("Invalid %s. please check your request data and try again", dataName)
}
