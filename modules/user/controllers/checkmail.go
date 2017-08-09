package user

import (
	"context"
	"errors"
	"net/http"

	"clickyab.com/crab/modules/domain/dmn"
	middleware2 "clickyab.com/crab/modules/domain/middleware"
	"clickyab.com/crab/modules/user/aaa"
)

// @Validate {
// }
type checkMailPayload struct {
	Email string `json:"email" validate:"email" error:"email is invalid"`
}

type checkMailResponse struct {
	Domains       []dmn.Domain `json:"domains"`
	CurrentDomain bool         `json:"current_domain"`
}

// checkMail check mail in system
// @Route {
// 		url = /mail/check
//		method = post
//      payload = checkMailPayload
//		200 = checkMailResponse
//		404 = controller.ErrorResponseSimple
// }
func (u Controller) checkMail(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*checkMailPayload)
	currentDomain := middleware2.MustGetDomain(ctx)
	m := aaa.NewAaaManager()
	// find user domains
	domains := m.FindUserDomainsByEmail(pl.Email)
	if len(domains) == 0 {
		// no active domain found
		u.NotFoundResponse(w, errors.New("domain not found"))
		return
	}
	var currentDomainFound bool
	for i := range domains {
		if domains[i].Name == currentDomain.Name {
			currentDomainFound = true
			break
		}
	}
	u.OKResponse(w, checkMailResponse{
		CurrentDomain: currentDomainFound,
		Domains:       domains,
	})
}
