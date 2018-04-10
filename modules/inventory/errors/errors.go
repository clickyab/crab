package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// AccessDeniedErr access denied
	AccessDeniedErr = t9e.G("access denied. can't edit inventory")
	// EmptyPublisherSelectedErr
	EmptyPublisherSelectedErr = t9e.G("inventory cant be empty")
)

// NotFoundError maker
func NotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("inventory with identifier %d not found, please check your request data.", id)
	}

	return t9e.G("inventory not found, please check your request data.")
}

// MaxPubLimit reached
func MaxPubLimit(limit int) error {
	return t9e.G("the publishers should not exceeds %d", limit)
}
