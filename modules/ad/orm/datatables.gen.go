// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	CreativeDataTableArray []CreativeDataTable
)

func (cdta CreativeDataTableArray) Filter(u permission.Interface) CreativeDataTableArray {
	res := make(CreativeDataTableArray, len(cdta))
	for i := range cdta {
		res[i] = cdta[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (cdt CreativeDataTable) Filter(u permission.Interface) CreativeDataTable {
	action := []string{}
	res := CreativeDataTable{}

	res.Name = cdt.Name

	res.Status = cdt.Status

	res.Type = cdt.Type

	res.Impression = cdt.Impression

	res.Click = cdt.Click

	res.ECPC = cdt.ECPC

	res.ECTR = cdt.ECTR

	res.ECPM = cdt.ECPM

	res.Spend = cdt.Spend

	res.Conversion = cdt.Conversion

	res.CreatedAt = cdt.CreatedAt

	res.Actions = cdt.Actions

	if _, ok := u.HasOn("edit_creative", cdt.OwnerID, cdt.ParentIDs, cdt.DomainID, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "edit")
	}

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

	permission.Register("edit_creative", "edit_creative")

}
