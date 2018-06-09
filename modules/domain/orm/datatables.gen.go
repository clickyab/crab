// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	DomainDetailsArray []DomainDetails
)

func (dda DomainDetailsArray) Filter(u permission.Interface) DomainDetailsArray {
	res := make(DomainDetailsArray, len(dda))
	for i := range dda {
		res[i] = dda[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (dd DomainDetails) Filter(u permission.Interface) DomainDetails {
	action := []string{}
	res := DomainDetails{}

	res.ID = dd.ID

	res.Title = dd.Title

	res.Status = dd.Status

	res.CorporationName = dd.CorporationName

	res.DomainBase = dd.DomainBase

	res.OwnerEmail = dd.OwnerEmail

	res.Balance = dd.Balance

	res.Actions = dd.Actions

	if _, ok := u.HasOn("get_domain", dd.OwnerID, dd.DomainID, false, false); ok {
		action = append(action, "detail")
	}

	if _, ok := u.HasOn("edit_domain", dd.OwnerID, dd.DomainID, false, false); ok {
		action = append(action, "edit")
	}

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

	permission.Register("get_domain", "get_domain")

	permission.Register("edit_domain", "edit_domain")

}
