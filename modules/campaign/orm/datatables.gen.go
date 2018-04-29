// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	CampaignDailyDataTableArray []CampaignDailyDataTable

	CampaignDetailsArray []CampaignDetails
)

func (cddta CampaignDailyDataTableArray) Filter(u permission.Interface) CampaignDailyDataTableArray {
	res := make(CampaignDailyDataTableArray, len(cddta))
	for i := range cddta {
		res[i] = cddta[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (cddt CampaignDailyDataTable) Filter(u permission.Interface) CampaignDailyDataTable {
	action := []string{}
	res := CampaignDailyDataTable{}

	res.CreatedAt = cddt.CreatedAt

	res.Imp = cddt.Imp

	res.Click = cddt.Click

	res.Conv = cddt.Conv

	res.Cpm = cddt.Cpm

	res.Cpc = cddt.Cpc

	res.Spent = cddt.Spent

	res.Cpa = cddt.Cpa

	res.Ctr = cddt.Ctr

	res.Actions = cddt.Actions

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}

func (cda CampaignDetailsArray) Filter(u permission.Interface) CampaignDetailsArray {
	res := make(CampaignDetailsArray, len(cda))
	for i := range cda {
		res[i] = cda[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (cd CampaignDetails) Filter(u permission.Interface) CampaignDetails {
	action := []string{}
	res := CampaignDetails{}

	res.ID = cd.ID

	res.Title = cd.Title

	res.Status = cd.Status

	res.Kind = cd.Kind

	res.TotalImp = cd.TotalImp

	res.TotalClick = cd.TotalClick

	res.ECTR = cd.ECTR

	res.ECPC = cd.ECPC

	res.ECPM = cd.ECPM

	res.TotalSpend = cd.TotalSpend

	res.MaxBid = cd.MaxBid

	res.Conversion = cd.Conversion

	res.TotalBudget = cd.TotalBudget

	res.TodaySpend = cd.TodaySpend

	res.CreatedAt = cd.CreatedAt

	res.StartAt = cd.StartAt

	res.EndAt = cd.EndAt

	res.TodayCTR = cd.TodayCTR

	res.TodayImp = cd.TodayImp

	res.TodayClick = cd.TodayClick

	res.Creative = cd.Creative

	res.OwnerEmail = cd.OwnerEmail

	res.ConversionRate = cd.ConversionRate

	res.CPA = cd.CPA

	res.Strategy = cd.Strategy

	res.Exchange = cd.Exchange

	res.OwnerID = cd.OwnerID

	res.DomainID = cd.DomainID

	res.Actions = cd.Actions

	if _, ok := u.HasOn("campaign_archive", cd.OwnerID, cd.ParentIDs, cd.DomainID, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "archive")
	}

	if _, ok := u.HasOn("campaign_copy", cd.OwnerID, cd.ParentIDs, cd.DomainID, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "copy")
	}

	if _, ok := u.HasOn("campaign_detail", cd.OwnerID, cd.ParentIDs, cd.DomainID, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "detail")
	}

	if _, ok := u.HasOn("campaign_edit", cd.OwnerID, cd.ParentIDs, cd.DomainID, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "edit")
	}

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

	permission.Register("campaign_archive", "campaign_archive")

	permission.Register("campaign_copy", "campaign_copy")

	permission.Register("campaign_detail", "campaign_detail")

	permission.Register("campaign_edit", "campaign_edit")

}
