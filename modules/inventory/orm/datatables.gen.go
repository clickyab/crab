// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	InventoryDataTableArray []InventoryDataTable

	PublisherDataTableArray []PublisherDataTable
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

	res.OwnerID = idt.OwnerID

	res.DomainID = idt.DomainID

	res.ParentIDs = idt.ParentIDs

	res.Actions = idt.Actions

	res.ID = idt.ID

	res.CreatedAt = idt.CreatedAt

	res.UpdatedAt = idt.UpdatedAt

	res.UserID = idt.UserID

	res.DomainID = idt.DomainID

	res.Label = idt.Label

	res.Status = idt.Status

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}

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

	res.OwnerID = pdt.OwnerID

	res.DomainID = pdt.DomainID

	res.ParentIDs = pdt.ParentIDs

	res.Actions = pdt.Actions

	res.ID = pdt.ID

	res.Name = pdt.Name

	res.Domain = pdt.Domain

	res.Categories = pdt.Categories

	res.Supplier = pdt.Supplier

	res.Kind = pdt.Kind

	res.Status = pdt.Status

	res.CreatedAt = pdt.CreatedAt

	res.UpdatedAt = pdt.UpdatedAt

	res.DeletedAt = pdt.DeletedAt

	if _, ok := u.HasOn("none", pdt.OwnerID, pdt.ParentIDs, pdt.DomainID, permission.ScopeGlobal); ok {
		action = append(action, "edit")
	}

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

	permission.Register("none", "none")

}
