// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	InventoryDataTableArray []InventoryDataTable
)

func (idta InventoryDataTableArray) Filter(u permission.Interface) InventoryDataTableArray {
	res := make(InventoryDataTableArray, len(idta))
	for i := range idta {
		res[i] = idta[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (idt InventoryDataTable) Filter(u permission.Interface) InventoryDataTable {
	action := []string{}
	res := InventoryDataTable{}

	res.Actions = idt.Actions

	res.ID = idt.ID

	res.CreatedAt = idt.CreatedAt

	res.UpdatedAt = idt.UpdatedAt

	res.Active = idt.Active

	res.Name = idt.Name

	res.Domain = idt.Domain

	res.Cat = idt.Cat

	res.Publisher = idt.Publisher

	res.Kind = idt.Kind

	res.Status = idt.Status

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}
