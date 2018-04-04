package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// ISPKidError invalid isp kind error
	ISPKidError = t9e.G("invalid isp kind. you can just select BothISPKind or CellularISPKind orISPISPKind")
)

// NotFoundError maker
func NotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("campaign with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("campaign not found, please check your request data.")
}
