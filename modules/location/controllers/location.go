package location

import (
	"context"
	"net/http"

	"strconv"

	"clickyab.com/crab/modules/location/location"
	"github.com/clickyab/services/framework/controller"
	"github.com/rs/xmux"
)

// Controller is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /location
// }
type Controller struct {
	controller.Base
}

type countries []location.Country

// @Route {
// 		url = /countries
//		method = get
//		200 = countries
// }
func (ctrl *Controller) Countries(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	m := location.NewLocationManager()
	ctrl.JSON(w, http.StatusOK, m.ListCountries())
}

type provinces []location.Province

// @Route {
// 		url = /provinces/:country_id
//		method = get
//		200 = provinces
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) Provinces(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	d, e := strconv.ParseInt(xmux.Param(ctx, "country_id"), 10, 64)
	if e != nil || d == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := location.NewLocationManager()
	c := &location.Country{ID: d}
	res := m.GetCountryProvinces(c)
	if len(res) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctrl.JSON(w, http.StatusOK, res)
}

type cities []location.City

// @Route {
// 		url = /cities/:provinces_id
//		method = get
//		200 = cities
//		400 = controller.ErrorResponseSimple
//		404 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) Cities(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	d, e := strconv.ParseInt(xmux.Param(ctx, "provinces_id"), 10, 64)
	if e != nil || d == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	m := location.NewLocationManager()
	c := &location.Province{ID: d}
	res := m.GetProvinceCities(c)
	if len(res) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctrl.JSON(w, http.StatusOK, res)
}

