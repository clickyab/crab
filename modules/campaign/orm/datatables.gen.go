// Code generated build with datatable DO NOT EDIT.

package orm

import (
	"strings"

	"github.com/clickyab/services/permission"
)

type (
	CampaignDetailsArray []CampaignDetails

	CampaignLogArray []CampaignLog

	PublisherDetailsArray []PublisherDetails

	CampaignDailyArray []CampaignDaily
)

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

	res.Actions = cd.Actions

	if _, ok := u.HasOn("archive_campaign", cd.OwnerID, cd.DomainID, false, false, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "archive")
	}

	if _, ok := u.HasOn("copy_campaign", cd.OwnerID, cd.DomainID, false, false, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "copy")
	}

	if _, ok := u.HasOn("get_campaign", cd.OwnerID, cd.DomainID, false, false, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "detail")
	}

	if _, ok := u.HasOn("edit_campaign", cd.OwnerID, cd.DomainID, false, false, permission.ScopeSelf, permission.ScopeGlobal); ok {
		action = append(action, "edit")
	}

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

	permission.Register("archive_campaign", "archive_campaign")

	permission.Register("copy_campaign", "copy_campaign")

	permission.Register("get_campaign", "get_campaign")

	permission.Register("edit_campaign", "edit_campaign")

}

func (cla CampaignLogArray) Filter(u permission.Interface) CampaignLogArray {
	res := make(CampaignLogArray, len(cla))
	for i := range cla {
		res[i] = cla[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (cl CampaignLog) Filter(u permission.Interface) CampaignLog {
	action := []string{}
	res := CampaignLog{}

	res.CreatedAt = cl.CreatedAt

	res.Action = cl.Action

	res.ImpersonatorEmail = cl.ImpersonatorEmail

	res.ManipulatorEmail = cl.ManipulatorEmail

	res.OwnerEmail = cl.OwnerEmail

	res.Data = cl.Data

	res.CampaignName = cl.CampaignName

	res.Kind = cl.Kind

	res.Strategy = cl.Strategy

	res.StartAt = cl.StartAt

	res.EndAt = cl.EndAt

	res.TotalBudget = cl.TotalBudget

	res.DailyBudget = cl.DailyBudget

	res.MaxBid = cl.MaxBid

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}

func (pda PublisherDetailsArray) Filter(u permission.Interface) PublisherDetailsArray {
	res := make(PublisherDetailsArray, len(pda))
	for i := range pda {
		res[i] = pda[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (pd PublisherDetails) Filter(u permission.Interface) PublisherDetails {
	action := []string{}
	res := PublisherDetails{}

	res.Domain = pd.Domain

	res.Impression = pd.Impression

	res.Click = pd.Click

	res.ECPC = pd.ECPC

	res.ECTR = pd.ECTR

	res.ECPM = pd.ECPM

	res.Spend = pd.Spend

	res.Conversion = pd.Conversion

	res.ConversionRate = pd.ConversionRate

	res.CPA = pd.CPA

	res.Actions = pd.Actions

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}

func (cda CampaignDailyArray) Filter(u permission.Interface) CampaignDailyArray {
	res := make(CampaignDailyArray, len(cda))
	for i := range cda {
		res[i] = cda[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (cd CampaignDaily) Filter(u permission.Interface) CampaignDaily {
	action := []string{}
	res := CampaignDaily{}

	res.Date = cd.Date

	res.Impression = cd.Impression

	res.Click = cd.Click

	res.ECTR = cd.ECTR

	res.ECPC = cd.ECPC

	res.ECPM = cd.ECPM

	res.Spend = cd.Spend

	res.Conversion = cd.Conversion

	res.ConversionRate = cd.ConversionRate

	res.CPA = cd.CPA

	res.Actions = strings.Join(action, ",")
	return res
}

func init() {

}
