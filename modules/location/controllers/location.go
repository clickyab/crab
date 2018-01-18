package location

import (
	"context"
	"errors"
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

// list of countries
// @Rest {
// 		url = /countries
//		method = get
// }
func (ctrl *Controller) countries(ctx context.Context, r *http.Request) (countries, error) {
	res := location.NewLocationManager().ListCountries()
	return countries(res), nil
}

type provinces []location.Province

// list of provinces
// @Rest {
// 		url = /provinces/:country_id
//		method = get
// }
func (ctrl *Controller) provinces(ctx context.Context, r *http.Request) (provinces, error) {
	d, e := strconv.ParseInt(xmux.Param(ctx, "country_id"), 10, 64)
	if e != nil || d == 0 {
		return nil, errors.New("country id is not correct")
	}
	m := location.NewLocationManager()
	c := &location.Country{ID: d}
	res := m.GetCountryProvinces(c)
	if len(res) == 0 {
		return nil, errors.New("country id is not correct")
	}
	return provinces(res), nil
}

type cities []location.City

// list of cities
// @Rest {
// 		url = /cities/:provinces_id
//		method = get
// }
func (ctrl *Controller) cities(ctx context.Context, r *http.Request) (cities, error) {
	d, e := strconv.ParseInt(xmux.Param(ctx, "provinces_id"), 10, 64)
	if e != nil || d == 0 {
		return nil, errors.New("city id is not correct")
	}
	m := location.NewLocationManager()
	c := &location.Province{ID: d}
	res := m.GetProvinceCities(c)
	if len(res) == 0 {
		return nil, errors.New("city id is not correct")

	}
	return cities(res), nil
}
