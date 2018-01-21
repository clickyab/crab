// Code generated build with graph DO NOT EDIT.

package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

var (
	maxRangeChart = config.RegisterDuration("srv.graph.max_range", 24*90*time.Hour, "maximum possible for graph date range")
	epochChart    time.Time
	layoutChart                 = "2006010215"
	factorChart   time.Duration = 24
)

type graphChartResponse struct {
	Format string           `json:"format"`
	From   time.Time        `json:"from"`
	To     time.Time        `json:"to"`
	Type   string           `json:"type"`
	Data   []graphChartData `json:"data"`
}

type graphChartData struct {
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	Hidden    bool      `json:"hidden"`
	Type      string    `json:"type"`
	Data      []float64 `json:"data"`
	Sum       float64   `json:"sum"`
	Avg       float64   `json:"avg"`
	Min       float64   `json:"min"`
	Max       float64   `json:"max"`
	OmitEmpty bool      `json:"-"`
}

// @Route {
//		url = /graph/all
//		method = get
//		resource = campaign_graph:self
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		200 = graphChartResponse
//		_kind_ = string , filter the kind field valid values are "web","app"
//		_type_ = string , filter the type field valid values are "banner","vast","native"
//		_owner_email_ = string , search the owner_email field
//		_title_ = string , search the title field
// }
func (ctrl *Controller) graphChart(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	filter := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" && orm.CampaignKind(e).IsValid() {
		filter["cp.kind"] = e
	}
	if e := r.URL.Query().Get("type"); e != "" && orm.CampaignType(e).IsValid() {
		filter["cp.type"] = e
	}

	search := make(map[string]string)

	if e := r.URL.Query().Get("owner_email"); e != "" {
		search["owner.email"] = e
	}

	if e := r.URL.Query().Get("title"); e != "" {
		search["cp.title"] = e
	}

	params := make(map[string]string)
	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}
	from, to, err := dateParamChart(r.URL.Query().Get("from"), r.URL.Query().Get("to"))
	if err != nil {
		ctrl.BadResponse(w, err)
		return
	}
	l, fn := dateRangeChart(from, to)
	tm := make(map[string]graphChartData)
	pc := permission.NewInterfaceComplete(usr, usr.ID, "campaign_graph", "self", domain.ID)

	tm["avg_cpc"] = graphChartData{
		Name:      "avg_cpc",
		Title:     "Avg. CPC",
		Type:      "line",
		Hidden:    false,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}

	tm["avg_cpm"] = graphChartData{
		Name:      "avg_cpm",
		Title:     "Avg. CPM",
		Type:      "line",
		Hidden:    false,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}

	tm["ctr"] = graphChartData{
		Name:      "ctr",
		Title:     "CTR",
		Type:      "line",
		Hidden:    false,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}

	tm["imp"] = graphChartData{
		Name:      "imp",
		Title:     "Total Impression",
		Type:      "bar",
		Hidden:    true,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}

	tm["click"] = graphChartData{
		Name:      "click",
		Title:     "Click",
		Type:      "line",
		Hidden:    true,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}

	tm["total_spent"] = graphChartData{
		Name:      "total_spent",
		Title:     "Total spent",
		Type:      "line",
		Hidden:    false,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}
	for i, v := range orm.NewOrmManager().FillCampaignGraph(pc, filter, search, params, from, to) {
		m, err := fn(v.ID)
		assert.Nil(err)

		txavg_cpc := tm["avg_cpc"]
		cvavg_cpc := v.AvgCPC
		tm["avg_cpc"].Data[m] += cvavg_cpc
		txavg_cpc.Sum += cvavg_cpc
		if i == 0 {
			txavg_cpc.Min = cvavg_cpc
			txavg_cpc.Max = cvavg_cpc
		} else {
			if cvavg_cpc > txavg_cpc.Max {
				txavg_cpc.Max = cvavg_cpc
			}
			if txavg_cpc.Min > cvavg_cpc {
				txavg_cpc.Min = cvavg_cpc
			}
		}
		tm["avg_cpc"] = txavg_cpc

		txavg_cpm := tm["avg_cpm"]
		cvavg_cpm := v.AvgCPM
		tm["avg_cpm"].Data[m] += cvavg_cpm
		txavg_cpm.Sum += cvavg_cpm
		if i == 0 {
			txavg_cpm.Min = cvavg_cpm
			txavg_cpm.Max = cvavg_cpm
		} else {
			if cvavg_cpm > txavg_cpm.Max {
				txavg_cpm.Max = cvavg_cpm
			}
			if txavg_cpm.Min > cvavg_cpm {
				txavg_cpm.Min = cvavg_cpm
			}
		}
		tm["avg_cpm"] = txavg_cpm

		txctr := tm["ctr"]
		cvctr := v.Ctr
		tm["ctr"].Data[m] += cvctr
		txctr.Sum += cvctr
		if i == 0 {
			txctr.Min = cvctr
			txctr.Max = cvctr
		} else {
			if cvctr > txctr.Max {
				txctr.Max = cvctr
			}
			if txctr.Min > cvctr {
				txctr.Min = cvctr
			}
		}
		tm["ctr"] = txctr

		tximp := tm["imp"]
		cvimp := float64(v.TotalImp)
		tm["imp"].Data[m] += cvimp
		tximp.Sum += cvimp
		if i == 0 {
			tximp.Min = cvimp
			tximp.Max = cvimp
		} else {
			if cvimp > tximp.Max {
				tximp.Max = cvimp
			}
			if tximp.Min > cvimp {
				tximp.Min = cvimp
			}
		}
		tm["imp"] = tximp

		txclick := tm["click"]
		cvclick := float64(v.TotalClick)
		tm["click"].Data[m] += cvclick
		txclick.Sum += cvclick
		if i == 0 {
			txclick.Min = cvclick
			txclick.Max = cvclick
		} else {
			if cvclick > txclick.Max {
				txclick.Max = cvclick
			}
			if txclick.Min > cvclick {
				txclick.Min = cvclick
			}
		}
		tm["click"] = txclick

		txtotal_spent := tm["total_spent"]
		cvtotal_spent := float64(v.TotalSpent)
		tm["total_spent"].Data[m] += cvtotal_spent
		txtotal_spent.Sum += cvtotal_spent
		if i == 0 {
			txtotal_spent.Min = cvtotal_spent
			txtotal_spent.Max = cvtotal_spent
		} else {
			if cvtotal_spent > txtotal_spent.Max {
				txtotal_spent.Max = cvtotal_spent
			}
			if txtotal_spent.Min > cvtotal_spent {
				txtotal_spent.Min = cvtotal_spent
			}
		}
		tm["total_spent"] = txtotal_spent
	}
	res := graphChartResponse{
		From:   from,
		To:     to,
		Format: "daily",  // hourly|daily|weekly|monthly|yearly
		Type:   "number", //  number|percent
		Data:   make([]graphChartData, 0),
	}
	for _, v := range tm {
		if v.Sum == 0 && v.OmitEmpty {
			continue
		}
		if l != 0 {
			v.Avg = v.Sum / float64(l)
		}
		res.Data = append(res.Data, v)
	}
	ctrl.OKResponse(w, res)
}

func dateParamChart(f, t string) (time.Time, time.Time, error) {
	from, err := time.Parse(time.RFC3339Nano, f)
	from = from.Truncate(time.Hour * factorChart)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("wrong date format")
	}
	to, err := time.Parse(time.RFC3339Nano, t)
	to = to.Truncate(time.Hour * factorChart)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("wrong date format")
	}
	if to.Before(from) {
		from, to = to, from
	}
	if from.IsZero() && to.IsZero() {
		to = time.Now()
		from = to.AddDate(0, 0, -maxRangeChart.Int())
	} else if from.IsZero() {
		from = to.AddDate(0, 0, -maxRangeChart.Int())
	} else if to.IsZero() {
		to = from.AddDate(0, 0, maxRangeChart.Int())
	}

	if to.After(time.Now()) {
		to = time.Now()
	}

	if from.Before(to.AddDate(0, 0, -maxRangeChart.Int())) {
		from = to.AddDate(0, 0, -maxRangeChart.Int())
	}

	if from.Before(epochChart) {
		from = epochChart
	}

	return from, to, nil
}

func timeToIDChart(d time.Time) int64 {
	h := int64(d.Truncate(time.Hour * factorChart).Sub(epochChart).Hours())
	return (h / 24) + 1
}

func dateRangeChart(f, t time.Time) (int, func(int64) (int, error)) {
	from := f.Truncate(time.Hour * factorChart)
	to := t.Truncate(time.Hour * factorChart)
	res := make(map[string]int)
	for i := 0; ; i++ {
		x := from.AddDate(0, 0, i)
		if x.After(to) {
			break
		}
		res[fmt.Sprint(timeToIDChart(x))] = i
	}
	return len(res), func(m int64) (int, error) {

		if v, ok := res[fmt.Sprint(m)]; ok {

			return v, nil
		}
		return 0, errors.New("out of range. probably mismatch key type. check FillCampaignGraph annotation (e.g. daily to hourly or vice versa)")
	}
}

func init() {
	epochChart, _ = time.Parse(layoutChart, "2018010100")
}
