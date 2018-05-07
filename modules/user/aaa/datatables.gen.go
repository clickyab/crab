// Code generated build with datatable DO NOT EDIT.

package aaa

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	UserListArray []UserList
)

func (ula UserListArray) Filter(u permission.Interface) UserListArray {
	res := make(UserListArray, len(ula))
	for i := range ula {
		res[i] = ula[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (ul UserList) Filter(u permission.Interface) UserList {
	action := []string{}
	res := UserList{}

	res.ID = ul.ID

	res.FullName = ul.FullName

	res.Status = ul.Status

	res.Balance = ul.Balance

	res.Email = ul.Email

	res.CellPhone = ul.CellPhone

	res.LandLine = ul.LandLine

	res.AccountType = ul.AccountType

	res.SSN = ul.SSN

	res.Avatar = ul.Avatar

	res.CreatedAt = ul.CreatedAt

	res.Actions = ul.Actions

	if _, ok := u.HasOn("edit_user", ul.OwnerID, ul.ParentIDs, ul.DomainID, permission.ScopeGlobal); ok {
		action = append(action, "edit")
	}

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

	permission.Register("edit_user", "edit_user")

}
