// Code generated build with restful DO NOT EDIT.

package location

import (
	"context"
	"net/http"

	"github.com/clickyab/services/framework"
)

// list of countries
// @Route {
// 		url = /countries
//		method = get
//		200 = countries
//		400 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) countriesGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := ctrl.countries(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// list of provinces
// @Route {
// 		url = /provinces/:country_id
//		method = get
//		200 = provinces
//		400 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) provincesGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := ctrl.provinces(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}

// list of cities
// @Route {
// 		url = /cities/:province
//		method = get
//		200 = cities
//		400 = controller.ErrorResponseSimple
// }
func (ctrl *Controller) citiesGet(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	res, err := ctrl.cities(ctx, r)
	if err != nil {
		framework.Write(w, err, http.StatusBadRequest)
		return
	}
	framework.Write(w, res, http.StatusOK)
}
