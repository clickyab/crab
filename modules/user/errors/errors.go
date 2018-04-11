package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	//InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	//InvalidEmailError invalid email error
	InvalidEmailError = t9e.G("invalid email, please check your request data.")
	//InvalidEmailPassError invalid email or password for login error
	InvalidEmailPassError = t9e.G("invalid email or password, please check your request data.")
	//InvalidVerifyCodeError verify code is invalid error
	InvalidVerifyCodeError = t9e.G("verify code is invalid.")
	//UserNotVerifiedError when user is not verified
	UserNotVerifiedError = t9e.G("user is not verified yet. please check your email and verify it.")
	//UserBlockedError when user is block
	UserBlockedError = t9e.G("your user is blocked! please try to connect with our support team.")
)

// NotFoundError maker
func NotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("user with identifier %d not found, please check your request data.", id)
	}

	return t9e.G("user not found, please check your request data.")
}

// NotFoundWithDomainError maker
func NotFoundWithDomainError(dName string) error {
	if dName != "" {
		return t9e.G("can't find user with related domain", dName)
	}

	return t9e.G("can't find user with related domain")
}
