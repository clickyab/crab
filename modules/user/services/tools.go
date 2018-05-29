package services

import (
	"context"

	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/xlog"
	gom "github.com/go-sql-driver/mysql"
)

// WhiteLabelEditUserRoles edit user to whitelabel in a transaction
func WhiteLabelEditUserRoles(ctx context.Context, user *aaa.User, corp *aaa.Corporation, domainID int64, roles []*aaa.Role, managers []int64) error {
	m := aaa.NewAaaManager()
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	err = m.UpdateUser(user)
	if err != nil {
		return errors.UserUpdateErr
	}

	//update corporation
	if corp != nil {
		err = m.UpdateCorporation(corp)
		if err != nil {
			return errors.CorporationUpdateErr
		}
	}

	//assign roles
	err = m.AssignRoles(user.ID, domainID, roles...)
	if err != nil {
		xlog.GetWithError(ctx, errors.AssignManagersErr)
		return errors.DbAddUserRoleErr
	}

	//assign managers
	_, err = m.AssignManagers(user.ID, domainID, managers)
	if err != nil {
		xlog.GetWithError(ctx, errors.AssignManagersErr)
		return errors.AssignManagersErr
	}

	return nil
}

// WhiteLabelAddUserRoles register user to whitelabel in a transaction
func WhiteLabelAddUserRoles(ctx context.Context, user *aaa.User, corp *aaa.Corporation, domainID int64, roles []*aaa.Role) error {
	m := aaa.NewAaaManager()
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	// register user with first role item
	registerRole := roles[0] // added to make code readable
	err = m.RegisterUser(user, corp, domainID, registerRole.ID)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return errors.RegisterUserErr
		}
		if mysqlError.Number == 1062 {
			return errors.DuplicateEmailErr
		}
		xlog.GetWithError(ctx, err).Debug("error in add new user to whitelabel route")
		return errors.DBError
	}
	if len(roles) > 1 {
		roles = roles[1:]
		err = m.AddRolesToUser(user.ID, roles...)
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("database error when add role to user")
			return errors.DbAddUserRoleErr
		}
	}
	return nil
}
