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
	"clickyab.com/crab/modules/user/services"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// @Validate {
// }
type editUserPayload struct {
	userPayload
	Managers      []int64 `json:"managers" validate:"omitempty"`
	owner         *aaa.User
	currentDomain *orm.Domain
	managers      []aaa.ManagerUser
	Advantage     int64 `json:"advantage"`
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
	owner, err := m.FindUserWithParentsByID(userID, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.owner = owner
	ownerScope, ok := owner.Has(permission.ScopeSelf, "can_have_account", dm.ID)
	if len(p.Managers) > 0 && ok && ownerScope == permission.ScopeSelf {
		managers := m.FindManagersByIDsDomain(p.Managers, dm.ID)
		p.managers = managers
	}

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
	// check perm
	_, ok := currentUser.HasOn("edit_user", p.owner.ID, p.currentDomain.ID, true, true, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	p.owner.CityID = intToNullInt64(p.CityID)
	p.owner.LandLine = stringToNullString(p.LandLine)
	p.owner.Cellphone = stringToNullString(p.CellPhone)
	p.owner.PostalCode = stringToNullString(p.PostalCode)
	p.owner.FirstName = p.FirstName
	p.owner.LastName = p.LastName
	p.owner.Address = stringToNullString(p.Address)
	p.owner.SSN = stringToNullString(p.SSN)

	if !p.owner.Advantage.Valid {
		p.owner.Advantage = mysql.NullInt64{Valid: p.Advantage != 0, Int64: p.Advantage}
	}

	var managerIDs []int64
	for i := range p.managers {
		managerIDs = append(managerIDs, p.managers[i].ID)
	}

	err := services.WhiteLabelEditUser(ctx, p.owner, p.userPayload.corporation, p.currentDomain.ID, managerIDs)
	if err != nil {
		return nil, err
	}

	res := c.createUserResponse(p.owner, nil, p.managers)
	return &res, nil
}
