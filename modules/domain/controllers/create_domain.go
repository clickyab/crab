package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	gom "github.com/go-sql-driver/mysql"

	"clickyab.com/crab/modules/domain/middleware/domain"
	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/config"
)

var defaultOwnerRole = config.RegisterString("crab.modules.domain.default.owner.role", "Owner", "default domain role name")
var defaultAdminDomain = config.RegisterString("crab.modules.domain.default.admin.domain", "staging.crab.clickyab.ae", "default admin domain name")

// @Validate{
//}
type createDomainPayload struct {
	DomainBase string `json:"domain_base" validate:"required"`
	Title      string `json:"title" validate:"required"`

	Email     string `json:"email" validate:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password" validate:"gt=5"`
	Company   string `json:"company" validate:"gt=3"`

	Theme       string                 `json:"theme" validate:"required"`
	Logo        string                 `json:"logo" validate:"required"`
	SendMail    bool                   `json:"send_mail"`
	Description string                 `json:"description" validate:"omitempty"`
	Attributes  map[string]interface{} `json:"attributes" validate:"omitempty"`
	Status      orm.DomainStatus       `json:"status" validate:"required"`

	TotalBudget  int64 `json:"total_budget" validate:"required"`
	DailyBudget  int64 `json:"daily_budget" validate:"required"`
	WebNativeCPC int64 `json:"web_native_cpc" validate:"required"`
	WebBannerCPC int64 `json:"web_banner_cpc" validate:"required"`
	WebVastCPC   int64 `json:"web_vast_cpc" validate:"required"`
	AppNativeCPC int64 `json:"app_native_cpc" validate:"required"`
	AppBannerCPC int64 `json:"app_banner_cpc" validate:"required"`
	AppVastCPC   int64 `json:"app_vast_cpc" validate:"required"`
	WebCPC       int64 `json:"web_cpc" validate:"required"`
	AppCPC       int64 `json:"app_cpc" validate:"required"`

	WebNativeCPM int64 `json:"web_native_cpm" validate:"required"`
	WebBannerCPM int64 `json:"web_banner_cpm" validate:"required"`
	WebVastCPM   int64 `json:"web_vast_cpm" validate:"required"`
	AppNativeCPM int64 `json:"app_native_cpm" validate:"required"`
	AppBannerCPM int64 `json:"app_banner_cpm" validate:"required"`
	AppVastCPM   int64 `json:"app_vast_cpm" validate:"required"`
	WebCPM       int64 `json:"web_cpm" validate:"required"`
	AppCPM       int64 `json:"app_cpm" validate:"required"`

	Advantage int `json:"advantage" validate:"max=99,min=0"`

	logo *model.Upload
}

func (p *createDomainPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.Status.IsValid() {
		return errors.InvalidDomainStatus
	}

	//find logo by id
	uploadFile, err := model.NewModelManager().FindSectionUploadByID(p.Logo, "domain-image")
	if err != nil {
		return errors.LogoNotFound
	}
	p.logo = uploadFile

	return nil
}

// createDomain to domain
// @Rest {
// 		url = /create
//		protected = true
// 		method = post
//		resource = god:global
// }
func (c *Controller) createDomain(ctx context.Context, r *http.Request, p *createDomainPayload) (*orm.Domain, error) {
	currentUser := authz.MustGetUser(ctx)
	currentDomain := domain.MustGetDomain(ctx)
	// check permission
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "god", currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDeniedErr
	}

	// create domain object
	newDomain := &orm.Domain{
		Title:        p.Title,
		DomainBase:   p.DomainBase,
		Theme:        p.Theme,
		Logo:         mysql.NullString{Valid: p.logo.ID != "", String: p.logo.ID},
		Description:  mysql.NullString{Valid: p.Description != "", String: p.Description},
		Attributes:   p.Attributes,
		Status:       p.Status,
		Advantage:    p.Advantage,
		AppBannerCPC: p.AppBannerCPC,
		AppNativeCPC: p.AppNativeCPC,
		AppVastCPC:   p.AppVastCPC,
		WebBannerCPC: p.WebBannerCPC,
		WebNativeCPC: p.WebNativeCPC,
		WebVastCPC:   p.WebVastCPC,
		AppBannerCPM: p.AppBannerCPM,
		AppNativeCPM: p.AppNativeCPM,
		AppVastCPM:   p.AppVastCPM,
		WebBannerCPM: p.WebBannerCPM,
		WebNativeCPM: p.WebNativeCPM,
		WebVastCPM:   p.WebVastCPM,
		WebCPC:       p.WebCPC,
		WebCPM:       p.WebCPM,
		AppCPM:       p.AppCPM,
		AppCPC:       p.AppCPC,
		TotalBudget:  p.TotalBudget,
		DailyBudget:  p.DailyBudget,
	}

	user := &aaa.User{
		Email:     p.Email,
		Password:  p.Password,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Status:    aaa.ActiveUserStatus,
	}

	corp := &aaa.Corporation{}
	corp.LegalName = p.Company

	role := &aaa.Role{
		Name:     defaultOwnerRole.String(),
		DomainID: newDomain.ID,
	}

	err := createWhiteLabel(user, corp, newDomain, role)
	if err != nil {
		return nil, err
	}

	// send email to user its email and pass

	if p.SendMail {
		assert.Nil(mailer.LoginInfoEmail(user, p.Password, r))
	}
	return newDomain, nil
}

// createWhiteLabel try to create white label
func createWhiteLabel(user *aaa.User, corp *aaa.Corporation, domain *orm.Domain, role *aaa.Role) error {
	m := orm.NewOrmManager()
	err := m.Begin()
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()

	// create domain
	err = m.CreateDomain(domain)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return errors.CreateDomainErr
		}
		if mysqlError.Number == 1062 {
			return errors.AlreadyExistDomainErr
		}
	}
	// create role
	aManger, err := aaa.NewAaaManagerFromTransaction(m.GetWDbMap())
	if err != nil {
		return err
	}
	role.DomainID = domain.ID
	err = aManger.CreateRole(role)
	if err != nil {
		return errors.CreateAdminRoleERR
	}

	//create role perms, first get perm of admin
	//then insert the perm in role perm table
	perms, err := aManger.FindRolePermByName(defaultOwnerRole.String(), defaultAdminDomain.String())
	if err != nil {
		return errors.FindAdminPermErr
	}

	for i := range perms {
		rolePerm := &aaa.RolePermission{
			Perm:   perms[i].Perm,
			RoleID: role.ID,
			Scope:  perms[i].Scope,
		}
		err = aManger.CreateRolePermission(rolePerm)
		if err != nil {
			return errors.CreateRolePermErr
		}
	}

	// register user
	err = aManger.RegisterUser(user, corp, domain.ID, role.ID)
	if err != nil {
		mysqlError, ok := err.(*gom.MySQLError)
		if !ok {
			return errors.RegisterUserErr
		}
		if mysqlError.Number == 1062 {
			return errors.AlreadyExistUserErr
		}
	}
	if err != nil {
		return errors.RegisterUserErr
	}
	return nil
}
