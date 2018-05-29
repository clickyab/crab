package user

import (
	"context"
	"net/http"
	"strconv"

	"database/sql"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate {
// }
type editUserPayload struct {
	userPayload
	Managers      []int64 `json:"managers" validate:"omitempty"`
	RolesID       []int64 `json:"roles_id" validate:"required" error:"roles id is required"`
	owner         *aaa.User
	currentDomain *orm.Domain
	roles         []*aaa.Role
	managers      []*aaa.ManagerUser
}

func (p *editUserPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm

	id := xmux.Param(ctx, "id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}

	m := aaa.NewAaaManager()

	if len(p.Managers) > 0 {
		managers := m.FindManagersByIDsDomain(p.Managers, dm.ID)
		p.managers = managers
	}

	// find roles
	// validate if role ids is valid and not in forbidden keys
	roles, err := m.FindRolesByDomainExclude(p.RolesID, dm.ID, getForbiddenRoles()...)
	if err != nil {
		xlog.GetWithError(ctx, err).Debug("database error, can't find role id, or role is forbidden")
		return errors.InvalidOrForbiddenRoleErr
	}
	if len(roles) == 0 {
		return errors.InvalidOrForbiddenRoleErr
	}
	p.roles = roles
	owner, err := m.FindUserWithParentsByID(userID, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.owner = owner
	cc, err := m.FindCorporationByUserID(p.owner.ID)
	if err != nil {
		if err != sql.ErrNoRows {
			xlog.GetWithError(ctx, errors.DBError)
			return errors.DBError
		}
		// user is personal
		if !p.Gender.IsValid() || p.Gender == aaa.NotSpecifiedGender {
			return errors.GenderInvalid
		}
		p.owner.Gender = p.Gender
	} else { //user is corporation
		if p.LegalName == "" {
			return errors.LegalEmptyErr
		}
		p.userPayload.corporation = cc
		p.userPayload.corporation.LegalName = p.LegalName
		p.userPayload.corporation.EconomicCode = stringToNullString(p.EconomicCode)
		p.userPayload.corporation.LegalRegister = stringToNullString(p.LegalRegister)
	}

	return nil
}

// adminEdit route for edit user profile
// @Rest {
// 		url = /update/:id
//		protected = true
//		method = put
//		resource = edit_user:global
// }
func (c *Controller) adminEdit(ctx context.Context, r *http.Request, p *editUserPayload) (*userResponse, error) {
	currentUser := authz.MustGetUser(ctx)
	_, ok := aaa.CheckPermOn(p.owner, currentUser, "edit_user", p.currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	m := aaa.NewAaaManager()
	p.owner.CityID = intToNullInt64(p.CityID)
	p.owner.LandLine = stringToNullString(p.LandLine)
	p.owner.Cellphone = stringToNullString(p.CellPhone)
	p.owner.PostalCode = stringToNullString(p.PostalCode)
	p.owner.FirstName = p.FirstName
	p.owner.LastName = p.LastName
	p.owner.Address = stringToNullString(p.Address)
	p.owner.SSN = stringToNullString(p.SSN)

	var managerIDs []int64
	for i := range p.managers {
		managerIDs = append(managerIDs, p.managers[i].ID)
	}

	err := m.WhiteLabelEditUserRoles(ctx, p.owner, p.userPayload.corporation, p.currentDomain.ID, p.roles, managerIDs)
	if err != nil {
		return nil, err
	}

	p.owner.Roles = p.roles

	res := c.createUserResponse(p.owner, nil, p.managers)
	return &res, nil
}
