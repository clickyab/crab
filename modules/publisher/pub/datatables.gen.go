// Code generated build with datatable DO NOT EDIT.

package pub

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	PublisherDataTableArray []PublisherDataTable
)

func (pdta PublisherDataTableArray) Filter(u permission.Interface) PublisherDataTableArray {
	res := make(PublisherDataTableArray, len(pdta))
	for i := range pdta {
		res[i] = pdta[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (pdt PublisherDataTable) Filter(u permission.Interface) PublisherDataTable {
	action := []string{}
	res := PublisherDataTable{}

	res.Actions = pdt.Actions

	res.ID = pdt.ID

	res.UserID = pdt.UserID

	res.Name = pdt.Name

	res.Supplier = pdt.Supplier

	res.Domain = pdt.Domain

	res.PublisherType = pdt.PublisherType

	res.PubStatus = pdt.PubStatus

	res.CreatedAt = pdt.CreatedAt

	res.UpdatedAt = pdt.UpdatedAt

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}
