package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
)

// NotFoundError maker
func NotFoundError(id int) error {
	if id > 0 {
		return t9e.G("location with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("location not found, please check your request data.")
}
