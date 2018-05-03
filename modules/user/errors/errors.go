package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	//InvalidIDErr invalid id error
	InvalidIDErr = t9e.G("invalid id, please check your request data.")
	// InalidAuditAction when audit log action is invalid
	InalidAuditAction = t9e.G("invalid audit action, please check your data.")
	// InalidAuditDomainID when audit log action is invalid
	InalidAuditDomainID = t9e.G("invalid domain id for audit data")
	// InalidAuditPerm when audit log user perm is invalid
	InalidAuditPerm = t9e.G("invalid user perm for audit data")
	// InalidAuditPermScope when audit log user perm scope is invalid
	InalidAuditPermScope = t9e.G("invalid user perm scope for audit data")
	// InalidAuditUserID when audit log action is invalid
	InalidAuditUserID = t9e.G("invalid user id for audit data")
	// InalidAuditOwnerID when audit log action is invalid
	InalidAuditOwnerID = t9e.G("invalid owner id for audit data")
	// InalidAuditTargetID when audit log target id is invalid
	InalidAuditTargetID = t9e.G("invalid target id for audit data")
	// InalidAuditTargetModel when audit log target model is invalid
	InalidAuditTargetModel = t9e.G("invalid target model for audit data")
	// InalidAuditDetail when audit log action is invalid
	InalidAuditDetail = t9e.G("audit detail data can not be empty")
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
	//AlreadyVerifiedErr user is verified before
	AlreadyVerifiedErr = t9e.G("your user is verified before")
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

// NotFoundWithEmail maker
func NotFoundWithEmail(email string) error {
	if email != "" {
		return t9e.G("can't find user with email %s", email)
	}

	return t9e.G("user not found, please check your request data.")
}
