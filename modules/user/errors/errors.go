package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	//InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	//InvalidEmailError invalid email error
	InvalidEmailError = t9e.G("invalid email, please check your request data.")
	//InvalidVerifyCodeError verify code is invalid error
	InvalidVerifyCodeError = t9e.G("verify code is invalid.")
)

// NotFoundError maker
func NotFoundError(id int) error {
	if id > 0 {
		return t9e.G("user with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("user not found, please check your request data.")
}
