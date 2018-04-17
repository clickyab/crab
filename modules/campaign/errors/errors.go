package errors

import (
	"github.com/clickyab/services/gettext/t9e"
)

var (
	// InvalidIDErr invalid id, error
	InvalidIDErr error = t9e.G("invalid id, please check your request data.")
	// UnsupportTypeError unsupported type. error
	UnsupportTypeError error = t9e.G("unsupported type.")
	// TypeError invalid campaign type error
	TypeError error = t9e.G("invalid campaign type. you can select %s or %s or %s please check your request data and try again", "banner", "vast", "native")
	// KindError invalid campaign kind error
	KindError error = t9e.G("invalid campaign kind. you can select %s or %s please check your request data and try again", "web", "app")
	// InvalidStrategyError invalid campaign strategy error
	InvalidStrategyError error = t9e.G("invalid campaign cost strategy. you can select %s or %s or %s please check your request data and try again", "cpc", "cpa", "cpm")
	// InvalidCampaignStatusError invalid campaign status error
	InvalidCampaignStatusError error = t9e.G("invalid campaign status. you can select %s or %s or %s please check your request data and try again", "archive", "start", "pause")
	// EndTimeError campaign should end error
	EndTimeError error = t9e.G("campaign should end after start")
	// TimeScheduleError at least you error
	TimeScheduleError error = t9e.G("at least you should specify one schedule time for campaign")
	// ArchivedEditError can't error
	ArchivedEditError error = t9e.G("can't update or copy archived campaign")
	// DuplicateNameError duplicate campaign name error
	DuplicateNameError error = t9e.G("duplicate campaign name. campaign name should be a unique")
	// StartTimeError start date error
	StartTimeError error = t9e.G("start date can't be past")
	// CampaignStartTimeError start date error
	CampaignStartTimeError error = t9e.G("campaign start time should be defined")
	// ErrInventoryID there is no inventory with this id
	ErrInventoryID error = t9e.G("there is no inventory with this id")
	// AccessDenied error
	AccessDenied error = t9e.G("access denied! you don't have access for this action")
	// InventoryStateErr inventory state err
	InventoryStateErr error = t9e.G("inventory state should be either black_list or white_list")
	// UpdateCampaignErr error while update
	UpdateCampaignErr error = t9e.G("error while update campaign")
	// UpdateError when want update campaign or related data
	UpdateError error = t9e.G("can't update campaign main or related data")
	// NotFoundSchedule when can't find campaign schedule time
	NotFoundSchedule error = t9e.G("can't find campaign schedule time")
	// InventoryNotFound when can't find campaign inventory
	InventoryNotFound error = t9e.G("can't find campaign inventory")
	// NotFoundAttributes when can't find campaign attributes
	NotFoundAttributes error = t9e.G("can't find campaign attributes")
	// CreateError can't create new campaign
	CreateError error = t9e.G("db error! can't create new campaign")
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
