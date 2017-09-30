package user

import (
	"context"
	"net/http"

	dm "clickyab.com/crab/modules/domain/middleware/domain"
	domain "clickyab.com/crab/modules/domain/models"
	"clickyab.com/crab/modules/user/models"
)

// @Validate {
// }
type checkMailPayload struct {
	Email string `json:"email" validate:"email" error:"email is invalid"`
}

type checkMailResponse struct {
	Domains       domain.Domains `json:"domains"`
	CurrentDomain bool           `json:"current_domain"`
}

// checkMail check mail in system
// @Route {
// 		url = /mail/check
//		method = post
//		payload = checkMailPayload
//		200 = checkMailResponse
// }
func (u Controller) checkMail(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	pl := u.MustGetPayload(ctx).(*checkMailPayload)
	currentDomain := dm.MustGetDomain(ctx)
	m := models.NewModelsManager()
	// find userPayload domains
	domains := m.FindUserDomainsByEmail(pl.Email)
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
