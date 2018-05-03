// Code generated build with graph DO NOT EDIT.

package controllers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/financial/orm"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

var (
	maxRangeBillinggraphreport = config.RegisterDuration("srv.graph.max_range", 24*90*time.Hour, "maximum possible for graph date range")
	epochBillinggraphreport    time.Time
	layoutBillinggraphreport                 = "2006010215"
	factorBillinggraphreport   time.Duration = 24
)

type graphBillinggraphreportResponse struct {
	Format string                        `json:"format"`
	From   time.Time                     `json:"from"`
	To     time.Time                     `json:"to"`
	Type   string                        `json:"type"`
	Data   []graphBillinggraphreportData `json:"data"`
}

type graphBillinggraphreportData struct {
	Title     string    `json:"title"`
	Name      string    `json:"name"`
	Hidden    bool      `json:"hidden"`
	Type      string    `json:"type"`
	Order     int64     `json:"order"`
	Data      []float64 `json:"data"`
	Sum       float64   `json:"sum"`
	Avg       float64   `json:"avg"`
	Min       float64   `json:"min"`
	Max       float64   `json:"max"`
	OmitEmpty bool      `json:"-"`
}

// @Route {
//		url = /graph/spend
//		method = get
//		resource = get_billing:self
//		_from_ = string , from date rfc3339 ex:2002-10-02T15:00:00.05Z
//		_to_ = string , to date rfc3339 ex:2002-10-02T15:00:00.05Z
//		200 = graphBillinggraphreportResponse
// }
func (ctrl *Controller) graphBillinggraphreport(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	usr := authz.MustGetUser(ctx)
	domain := domain.MustGetDomain(ctx)
	filter := make(map[string]string)

	search := make(map[string]string)

	params := make(map[string]string)
	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}
	from, to, err := dateParamBillinggraphreport(r.URL.Query().Get("from"), r.URL.Query().Get("to"))
	if err != nil {
		ctrl.BadResponse(w, err)
		return
	}
	l, fn := dateRangeBillinggraphreport(from, to)
	tm := make(map[string]graphBillinggraphreportData)
	pc := permission.NewInterfaceComplete(usr, usr.ID, "get_billing", "self", domain.ID)

	tm["spend"] = graphBillinggraphreportData{
		Name:      "spend",
		Title:     "Spend",
		Type:      "line",
		Hidden:    false,
		Order:     1,
		OmitEmpty: false,
		Data:      make([]float64, l),
	}
	for i, v := range orm.NewOrmManager().FillGraph(pc, filter, search, params, from, to) {
		m, err := fn(v.ID)
		assert.Nil(err)

		txspend := tm["spend"]
		cvspend := float64(v.Spend)
		tm["spend"].Data[m] += cvspend
		txspend.Sum += cvspend
		if i == 0 {
			txspend.Min = cvspend
			txspend.Max = cvspend
		} else {
			if cvspend > txspend.Max {
				txspend.Max = cvspend
			}
			if txspend.Min > cvspend {
				txspend.Min = cvspend
			}
		}
		tm["spend"] = txspend
	}
	res := graphBillinggraphreportResponse{
		From:   from,
		To:     to,
		Format: "daily",  // hourly|daily|weekly|monthly|yearly
		Type:   "number", //  number|percent
		Data:   make([]graphBillinggraphreportData, 0),
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

func dateParamBillinggraphreport(f, t string) (time.Time, time.Time, error) {
	from, err := time.Parse(time.RFC3339Nano, f)
	from = from.Truncate(time.Hour * factorBillinggraphreport)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("wrong date format")
	}
	to, err := time.Parse(time.RFC3339Nano, t)
	to = to.Truncate(time.Hour * factorBillinggraphreport)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("wrong date format")
	}
	if to.Before(from) {
		from, to = to, from
	}
	if from.IsZero() && to.IsZero() {
		to = time.Now()
		from = to.AddDate(0, 0, -maxRangeBillinggraphreport.Int())
	} else if from.IsZero() {
		from = to.AddDate(0, 0, -maxRangeBillinggraphreport.Int())
	} else if to.IsZero() {
		to = from.AddDate(0, 0, maxRangeBillinggraphreport.Int())
	}

	if to.After(time.Now()) {
		to = time.Now()
	}

	if from.Before(to.AddDate(0, 0, -maxRangeBillinggraphreport.Int())) {
		from = to.AddDate(0, 0, -maxRangeBillinggraphreport.Int())
	}

	if from.Before(epochBillinggraphreport) {
		from = epochBillinggraphreport
	}

	return from, to, nil
}

func timeToIDBillinggraphreport(d time.Time) int64 {
	h := int64(d.Truncate(time.Hour * factorBillinggraphreport).Sub(epochBillinggraphreport).Hours())
	return (h / 24) + 1
}

func dateRangeBillinggraphreport(f, t time.Time) (int, func(int64) (int, error)) {
	from := f.Truncate(time.Hour * factorBillinggraphreport)
	to := t.Truncate(time.Hour * factorBillinggraphreport)
	res := make(map[string]int)
	for i := 0; ; i++ {
		x := from.AddDate(0, 0, i)
		if x.After(to) {
			break
		}
		res[fmt.Sprint(timeToIDBillinggraphreport(x))] = i
	}
	return len(res), func(m int64) (int, error) {

		if v, ok := res[fmt.Sprint(m)]; ok {

			return v, nil
		}
		return 0, errors.New("out of range. probably mismatch key type. check FillGraph annotation (e.g. daily to hourly or vice versa)")
	}
}

func init() {
	epochBillinggraphreport, _ = time.Parse(layoutBillinggraphreport, "2018010100")
}
