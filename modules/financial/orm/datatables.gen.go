// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	BillingDataTableArray []BillingDataTable
)

func (bdta BillingDataTableArray) Filter(u permission.Interface) BillingDataTableArray {
	res := make(BillingDataTableArray, len(bdta))
	for i := range bdta {
		res[i] = bdta[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (bdt BillingDataTable) Filter(u permission.Interface) BillingDataTable {
	action := []string{}
	res := BillingDataTable{}

	res.ID = bdt.ID

	res.UserID = bdt.UserID

	res.PayModel = bdt.PayModel

	res.FirstName = bdt.FirstName

	res.LastName = bdt.LastName

	res.Email = bdt.Email

	res.Amount = bdt.Amount

	res.Balance = bdt.Balance

	res.CreatedAt = bdt.CreatedAt

	res.Actions = bdt.Actions

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}
