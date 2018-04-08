package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	domainOrm "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
)

// @Validate {
// }
type checkMailPayload struct {
	Email string `json:"email" validate:"email" error:"email is invalid"`
}

type checkMailResponse struct {
	Domains       []domainOrm.Domain `json:"domains"`
	CurrentDomain bool               `json:"current_domain"`
}

// checkMail check mail in system
// @Rest {
// 		url = /mail/check
// 		method = post
// }
func (c *Controller) checkMail(ctx context.Context, r *http.Request, p *checkMailPayload) (*checkMailResponse, error) {
	currentDomain := domain.MustGetDomain(ctx)
	m := aaa.NewAaaManager()
	// find userPayload domains
	domains := m.FindUserDomainsByEmail(p.Email)
	var currentDomainFound bool
	for i := range domains {
		if domains[i].Name == currentDomain.Name {
			currentDomainFound = true
			break
		}
	}
	return &checkMailResponse{
		CurrentDomain: currentDomainFound,
		Domains:       domains,
	}, nil
}
