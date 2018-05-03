// Code generated build with graph DO NOT EDIT.

package orm

import "github.com/clickyab/services/permission"

type (
	BillingSpendGraphArray []BillingSpendGraph
)

func (bsga BillingSpendGraphArray) Filter(u permission.Interface) BillingSpendGraphArray {
	res := make(BillingSpendGraphArray, len(bsga))
	for i := range bsga {
		res[i] = bsga[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (bsg BillingSpendGraph) Filter(u permission.Interface) BillingSpendGraph {
	res := BillingSpendGraph{}

	res.ID = bsg.ID

	res.Spend = bsg.Spend

	return res
}

func init() {

}
