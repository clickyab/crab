package errors

import "github.com/clickyab/services/gettext/t9e"

var (
	// AccessDenied have not access error
	AccessDenied = t9e.G("access Denied!you have not access to do this on domains")
	// InvalidIDErr invalid id, error
	InvalidIDErr error = t9e.G("invalid id, please check your request data.")
	// InvalidDomainStatus domain status error
	InvalidDomainStatus error = t9e.G("domain status is invalid,you can select enable or disable")
	// CreateDomainErr create domain error
	CreateDomainErr error = t9e.G("create white label error")
	// AccessDeniedErr create domain access error
	AccessDeniedErr error = t9e.G("you can't create new domain")
	// AlreadyExistDomainErr already domain exist error
	AlreadyExistDomainErr error = t9e.G("a domain with this name is already exist.")
	// AlreadyExistUserErr already exist user error
	AlreadyExistUserErr error = t9e.G("a user already exist.")
	// UpdateDomainErr error in updating domain
	UpdateDomainErr error = t9e.G("update domain error.")
	// LogoNotFound logo not found
	LogoNotFound error = t9e.G("logo not found")
	// FindAdminPermErr error while finding perms
	FindAdminPermErr error = t9e.G("error while finding perms")
	// CreateRolePermErr error while creating role permission
	CreateRolePermErr error = t9e.G("error while creating role permission")
	// RegisterUserErr error while registering user
	RegisterUserErr error = t9e.G("error while registering user")
	// CreateAdminRoleERR error while creating admin role
	CreateAdminRoleERR error = t9e.G("error while creating admin role")
	// UpdateStatusDbErr error in update domain status
	UpdateStatusDbErr = t9e.G("an database error occurred when we try to update domain status ")
)

// DomainNotFoundError maker
func DomainNotFoundError(id int64) error {
	if id > 0 {
		return t9e.G("domain with identifier %d not found, please check your request data.", id)
	}

	return t9e.G("domain not found, please check your request data.")
}
