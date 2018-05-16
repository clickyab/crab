package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// InvalidIDErr invalid id, error
	InvalidIDErr error = t9e.G("invalid id, please check your request data.")
	// InvalidDomainStatus domain status error
	InvalidDomainStatus error = t9e.G("domain status is invalid,you can select enable or disable")
	// CreateDomainErr create domain error
	CreateDomainErr error = t9e.G("create domain error")
	// AccessDeniedErr create domain access error
	AccessDeniedErr error = t9e.G("you can't create new domain")
	// AlreadyExistErr already exist error
	AlreadyExistErr error = t9e.G("a domain with this name is already exist.")
)

// NotFoundError maker
func NotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("domain with identifier %s not found, please check your request data.", id)
	}

	return t9e.G("domain not found, please check your request data.")
}
