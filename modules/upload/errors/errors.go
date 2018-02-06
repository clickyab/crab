package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// InvalidFileTypeError invalid file type error
	InvalidFileTypeError = t9e.G("invalid file type.")
	// FileDimensionError can't get file dimensions error
	FileDimensionError = t9e.G("can't get file dimensions")
)

// NotFoundError maker
func NotFoundError(id int) error {
	if id > 0 {
		return t9e.G("file with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("file not found, please check your request data.")
}

// InvalidError maker
func InvalidError(dataName string) error {
	return t9e.G("Invalid %s. please check your request data and try again", dataName)
}
