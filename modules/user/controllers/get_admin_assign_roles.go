package user

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
)

type roleResp []aaa.Role

// getWhitelabelAssignRole get assign role for admin/owner
// @Rest {
// 		url = /whitelabel/roles
//		protected = true
// 		method = get
//		resource = get_assign_admin_roles:global
// }
func (c *Controller) getWhitelabelAssignRole(ctx context.Context, r *http.Request) (roleResp, error) {
	currentUser := authz.MustGetUser(ctx)
	currentDomain := domain.MustGetDomain(ctx)
	currentUser.SetUserRole(currentDomain.ID)
	roles := aaa.NewAaaManager().ListRolesWithFilter("level<?", currentUser.Role.Level)
	return roleResp(roles), nil
}
