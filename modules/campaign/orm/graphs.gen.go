// Code generated build with graph DO NOT EDIT.

package orm

import "github.com/clickyab/services/permission"

type (
	CampaignGraphArray []CampaignGraph
)

func (cga CampaignGraphArray) Filter(u permission.Interface) CampaignGraphArray {
	res := make(CampaignGraphArray, len(cga))
	for i := range cga {
		res[i] = cga[i].Filter(u)
	}

	return res
}

// Filter is for filtering base on permission
func (cg CampaignGraph) Filter(u permission.Interface) CampaignGraph {
	res := CampaignGraph{}

	res.OwnerEmail = cg.OwnerEmail

	res.Kind = cg.Kind

	res.Type = cg.Type

	res.Title = cg.Title

	res.ID = cg.ID

	res.AvgCPC = cg.AvgCPC

	res.AvgCPM = cg.AvgCPM

	res.Ctr = cg.Ctr

	res.TotalImp = cg.TotalImp

	res.TotalClick = cg.TotalClick

	res.TotalSpent = cg.TotalSpent

	return res
}

func init() {

}
