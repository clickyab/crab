package user

import (
	"context"
	"net/http"
	"strconv"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/errors"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/rs/xmux"
)

// @Validate {
// }
type editUserPayload struct {
	userPayload
	Managers      []int64 `json:"managers" validate:"required"`
	owner         *aaa.User
	currentDomain *orm.Domain
	corporation   *aaa.Corporation
}

func (p *editUserPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm

	id := xmux.Param(ctx, "id")
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.InvalidIDErr
	}

	if len(p.Managers) == 0 {
		return errors.EmptyManagerErr
	}

	m := aaa.NewAaaManager()

	managerIDs := m.FindManagersByIDsDomain(p.Managers, dm.ID)

	if len(managerIDs) != len(p.Managers) {
		return errors.ManagerMismatchErr
	}

	p.Managers = managerIDs

	owner, err := m.FindUserWithParentsByID(userID, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.owner = owner
	cc, err := m.FindCorporationByUserID(p.owner.ID)
	if err != nil {
		if !p.Gender.IsValid() || p.Gender == aaa.NotSpecifiedGender {
			return errors.GenderInvalid
		}
	} else {
		if p.LegalName == "" {
			return errors.LegalEmptyErr
		}
		p.corporation = cc
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
func (c *Controller) adminEdit(ctx context.Context, r *http.Request, p *editUserPayload) (*editAdminResp, error) {
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
	if p.corporation == nil {
		p.owner.Gender = p.Gender
	}
	p.owner.SSN = stringToNullString(p.SSN)

	err := m.UpdateUser(p.owner)
	if err != nil {
		return nil, errors.UserUpdateErr
	}

	if p.corporation != nil {
		p.corporation.LegalName = p.LegalName
		p.corporation.EconomicCode = stringToNullString(p.EconomicCode)
		p.corporation.LegalRegister = stringToNullString(p.LegalRegister)

		err = m.UpdateCorporation(p.corporation)
		if err != nil {
			return nil, errors.CorporationUpdateErr
		}
	}

	//assign mangers
	advisors, err := m.AssignManagers(p.owner.ID, p.currentDomain.ID, p.Managers)
	if err != nil {
		return nil, errors.AssignManagersErr
	}

	return &editAdminResp{
		Account:  c.createUserResponse(p.owner, nil),
		Managers: advisors,
	}, nil
}

// editAdminResp login or ping or other response
type editAdminResp struct {
	Account  userResponse   `json:"account"`
	Managers []*aaa.Advisor `json:"managers"`
}
