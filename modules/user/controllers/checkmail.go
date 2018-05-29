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
	m := domainOrm.NewOrmManager()
	// find user by email
	currentUser, err := aaa.NewAaaManager().FindUserByEmail(p.Email)
	if err != nil {
		return &checkMailResponse{
			CurrentDomain: false,
			Domains:       make([]domainOrm.Domain, 0),
		}, nil
	}
	// check if user is domain less
	if currentUser.DomainLess == true {
		// find all active domains
		allDomains := m.ListDomainsWithFilter("status=?", domainOrm.EnableDomainStatus)
		return &checkMailResponse{
			CurrentDomain: true,
			Domains:       allDomains,
		}, nil
	}
	// find user active domains
	domains := currentUser.FindUserActiveDomains()
	var currentDomainFound bool
	for i := range domains {
		if domains[i].DomainBase == currentDomain.DomainBase {
			currentDomainFound = true
			break
		}
	}
	return &checkMailResponse{
		CurrentDomain: currentDomainFound,
		Domains:       domains,
	}, nil
}
