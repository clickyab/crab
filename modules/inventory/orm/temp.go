package orm

import (
	"time"

	"math/rand"

	"github.com/clickyab/services/permission"
)

// @Graph {
// 	url = /graph
//  method = get
//  resource = inventory_graph:self
//	type = percent
//  200 = GraphResponse
// 	type = percent
//  format = daily
// }
type CampaignReport struct {
	Campaign string    `db:"name" search:"true"  `
	Kind     string    `db:"kind" filter:"true"`
	ID       time.Time `db:"time_id" key:""`
	CPC      int64     `db:"cpc" graph:"cpc,campaign cpc,line,true"`
	CPM      int64     `db:"cpm" graph:"cpm,campaign cpm,bar,false"`
}

func (m *Manager) GraphFiller(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string,
	from, to time.Time) []CampaignReport {

	res := make([]CampaignReport, 0)

	for _, v := range data {
		if v.ID.Before(from) || v.ID.After(to) {
			continue
		}
		if len(filters) == 0 {
			res = append(res, v)
			continue
		}
		for _, m := range filters {
			if v.Campaign == m {
				res = append(res, v)
			}
		}
	}
	return res
}

func init() {
	temp()
}

var data = make([]CampaignReport, 0)

func temp() {
	a := []string{"Niyosha", "Entekhab", "Digikala"}
	b := []string{"web", "app"}
	t, _ := time.Parse("20060102", "20180114")
	for i := 0; i < 10; i++ {
		data = append(data, CampaignReport{
			Campaign: a[rand.Int31n(3)],
			Kind:     b[rand.Int31n(2)],
			CPC:      rand.Int63n(10) + 1,
			CPM:      rand.Int63n(10) + 1,
			ID:       t.AddDate(0, 0, i),
		})
	}

}
