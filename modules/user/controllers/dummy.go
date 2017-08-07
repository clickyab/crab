package user

import (
	"context"
	"net/http"
)

// dummy route. Remove it when the real routes are ready
// @Route {
// 		url = /dummy
//		method = get
//      200 = controller.NormalResponse
//      400 = controller.ErrorResponseSimple
// }
func (u Controller) dummy(_ context.Context, w http.ResponseWriter, r *http.Request) {
	u.OKResponse(w, nil)
}
