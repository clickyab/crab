package controllers

import (
	"context"
	"errors"
	"net/http"

	"time"

	"reflect"

	"strings"

	"clickyab.com/crab/modules/inventory/orm"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

var (
	macRange = config.RegisterDuration("srv.graph.max_range", 24*90*time.Hour, "maximum possible for graph date range")
	epoch    time.Time
	layout   = "20060102"
)

type graphResponse struct {
	Format string      `json:"format"`
	From   time.Time   `json:"from"`
	To     time.Time   `json:"to"`
	Type   string      `json:"type"`
	Data   []graphData `json:"data"`
}

type graphData struct {
	Title  string  `json:"title"`
	Name   string  `json:"name"`
	Hidden bool    `json:"hidden"`
	Type   string  `json:"type"`
	Data   []int64 `json:"data"`
	//       ^^^ type can be int64 or float64 and should be set dynamically
}

func graph(
	pc permission.InterfaceComplete,
	filters map[string]string,
	search map[string]string,
	contextparams map[string]string, from, to time.Time) graphResponse {
	l, fn := dateRange(from, to)

	tr := orm.NewOrmManager().GraphFiller(pc, filters, search, contextparams, from, to)
	tm := make(map[string]graphData)

	for _, v := range tr {
		key, vals := getMeta(v)
		m, err := fn(key.(time.Time))
		assert.Nil(err)

		for _, k := range vals {
			var d graphData
			if rs, ok := tm[k.Name]; ok {
				d = rs
			} else {
				d = graphData{
					Name:   k.Name,
					Title:  k.Title,
					Hidden: k.Hidden,
					Type:   k.Type,
					Data:   make([]int64, l),
				}
			}

			d.Data[m] = k.Data[0]
			tm[k.Name] = d

		}
	}

	res := graphResponse{
		From:   from,
		To:     to,
		Format: "daily", // other options: weekly|monthly|yearly
		//       ^^^ set in template
		Type: "number", // or percent
		//     ^^^ set in template
		Data: make([]graphData, 0),
	}
	for _, v := range tm {
		res.Data = append(res.Data, v)
	}
	return res
}

// GraphRoute return all user inventories
// @Route {
// 		url = /graph
//		method = get
//		200 = graphResponse
// }
func (ctrl *Controller) GraphRoute(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	//usr := authz.MustGetUser(ctx)
	//domain := domain.MustGetDomain(ctx)

	filter := make(map[string]string)

	if e := r.URL.Query().Get("kind"); e != "" {
		filter["inventories.kind"] = e
	}
	search := make(map[string]string)

	if e := r.URL.Query().Get("name"); e != "" {
		search["inventories.name"] = e
	}
	params := make(map[string]string)

	for _, i := range xmux.Params(ctx) {
		params[i.Name] = xmux.Param(ctx, i.Name)
	}

	fp, tp := r.URL.Query().Get("from"), r.URL.Query().Get("to")

	from, to, err := dateParam(fp, tp)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//pc := permission.NewInterfaceComplete(usr, usr.ID, "inventory_graph", "self", domain.ID)

	ctrl.OKResponse(w, graph(nil, filter, search, params, from, to))
}

func getMeta(c orm.CampaignReport) (interface{}, []*graphData) {
	res := make([]*graphData, 0)
	var key interface{}
	t := reflect.TypeOf(c)
	v := reflect.ValueOf(c)
	for i := 0; i < t.NumField(); i++ {
		if _, ok := t.Field(i).Tag.Lookup("key"); ok {
			key = v.Field(i).Interface()
		}
		if c := t.Field(i).Tag.Get("graph"); c != "" {
			s := strings.Split(c, ",")
			if len(s) < 4 {
				panic("graph tag is wrong")
			}

			for i := range s {
				s[i] = strings.TrimSpace(s[i])
			}

			res = append(res, &graphData{
				Name:   s[0],
				Title:  s[1],
				Type:   s[2],
				Hidden: s[3] == "true",
				Data:   []int64{v.Field(i).Int()},
			})
		}
	}

	// find field tag graph and extract name
	// name, title, hidden, type
	return key, res
}

func dateParam(f, t string) (time.Time, time.Time, error) {
	from, err := time.Parse(time.RFC3339Nano, f)
	from = from.Truncate(time.Hour * 24)
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("wrong date format")
	}
	to, err := time.Parse(time.RFC3339Nano, t)
	to = to.Truncate(time.Hour * 24)

	if err != nil {
		return time.Time{}, time.Time{}, errors.New("wrong date format")
	}

	if to.Before(from) {
		from, to = to, from
	}

	if from.IsZero() && to.IsZero() {
		to = time.Now()
		from = to.AddDate(0, 0, -90)
	} else if from.IsZero() {
		from = to.AddDate(0, 0, -90)
	} else if to.IsZero() {
		to = from.AddDate(0, 0, 90)
	}
	if from.Before(epoch) {
		from = epoch
	}

	return from, to, nil
}

func timeToID(d time.Time) int64 {
	d, _ = time.Parse(layout, d.Format(layout))
	h := int64(d.Sub(epoch).Hours())
	return (h / 24) + 1
}

func timeTableRange(f, t time.Time) (int, func(int64) (int, error)) {
	from, _ := time.Parse(layout, f.Format(layout))
	to, _ := time.Parse(layout, t.Format(layout))
	res := make(map[int64]int)
	for i := 0; ; i++ {
		x := from.AddDate(0, 0, i)
		if x.After(to) {
			break
		}
		res[timeToID(x)] = i
	}
	return len(res), func(m int64) (int, error) {
		if v, ok := res[m]; ok {
			return v, nil
		}
		return 0, errors.New("out of range")
	}
}

func dateRange(f, t time.Time) (int, func(time.Time) (int, error)) {
	from := f.Truncate(time.Hour * 24)
	to := t.Truncate(time.Hour * 24)
	res := make(map[string]int)
	for i := 0; ; i++ {
		x := from.AddDate(0, 0, i)
		if x.After(to) {
			break
		}
		res[x.Format(layout)] = i
	}
	return len(res), func(m time.Time) (int, error) {
		if v, ok := res[m.Format(layout)]; ok {
			return v, nil
		}
		return 0, errors.New("out of range")
	}
}

func init() {
	epoch, _ = time.Parse(layout, "20180101")
}
