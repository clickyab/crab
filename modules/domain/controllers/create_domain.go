package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/errors"
	"clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/mysql"
	gom "github.com/go-sql-driver/mysql"

	"clickyab.com/crab/modules/upload/model"
	"clickyab.com/crab/modules/user/mailer"
	"clickyab.com/crab/modules/user/ucfg"
	"github.com/clickyab/services/assert"
)

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

	MinTotalBudget  int64 `json:"min_total_budget" validate:"required"`
	MinDailyBudget  int64 `json:"min_daily_budget" validate:"required"`
	MinWebNativeCPC int64 `json:"min_web_native_cpc" validate:"required"`
	MinWebBannerCPC int64 `json:"min_web_banner_cpc" validate:"required"`
	MinWebVastCPC   int64 `json:"min_web_vast_cpc" validate:"required"`
	MinAppNativeCPC int64 `json:"min_app_native_cpc" validate:"required"`
	MinAppBannerCPC int64 `json:"min_app_banner_cpc" validate:"required"`
	MinAppVastCPC   int64 `json:"min_app_vast_cpc" validate:"required"`
	MinWebCPC       int64 `json:"min_web_cpc" validate:"required"`
	MinAppCPC       int64 `json:"min_app_cpc" validate:"required"`

	MinWebNativeCPM int64 `json:"min_web_native_cpm" validate:"required"`
	MinWebBannerCPM int64 `json:"min_web_banner_cpm" validate:"required"`
	MinWebVastCPM   int64 `json:"min_web_vast_cpm" validate:"required"`
	MinAppNativeCPM int64 `json:"min_app_native_cpm" validate:"required"`
	MinAppBannerCPM int64 `json:"min_app_banner_cpm" validate:"required"`
	MinAppVastCPM   int64 `json:"min_app_vast_cpm" validate:"required"`
	MinWebCPM       int64 `json:"min_web_cpm" validate:"required"`
	MinAppCPM       int64 `json:"min_app_cpm" validate:"required"`

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
//		resource = create_domain:superGlobal
// }
func (c *Controller) createDomain(ctx context.Context, r *http.Request, p *createDomainPayload) (*orm.Domain, error) {
	// create domain object
	newDomain := &orm.Domain{
		Title:       p.Title,
		DomainBase:  p.DomainBase,
		Theme:       p.Theme,
		Logo:        mysql.NullString{Valid: p.logo.ID != "", String: p.logo.ID},
		Description: mysql.NullString{Valid: p.Description != "", String: p.Description},
		Attributes:  p.Attributes,
		Status:      p.Status,
		Advantage:   p.Advantage,
		UserConfig: orm.UserConfig{
			MinAppBannerCPC: p.MinAppBannerCPC,
			MinAppNativeCPC: p.MinAppNativeCPC,
			MinAppVastCPC:   p.MinAppVastCPC,
			MinWebBannerCPC: p.MinWebBannerCPC,
			MinWebNativeCPC: p.MinWebNativeCPC,
			MinWebVastCPC:   p.MinWebVastCPC,
			MinAppBannerCPM: p.MinAppBannerCPM,
			MinAppNativeCPM: p.MinAppNativeCPM,
			MinAppVastCPM:   p.MinAppVastCPM,
			MinWebBannerCPM: p.MinWebBannerCPM,
			MinWebNativeCPM: p.MinWebNativeCPM,
			MinWebVastCPM:   p.MinWebVastCPM,
			MinWebCPC:       p.MinWebCPC,
			MinWebCPM:       p.MinWebCPM,
			MinAppCPM:       p.MinAppCPM,
			MinAppCPC:       p.MinAppCPC,
			MinTotalBudget:  p.MinTotalBudget,
			MinDailyBudget:  p.MinDailyBudget,
		},
	}

	user := &aaa.User{
		Email:      p.Email,
		Password:   p.Password,
		FirstName:  p.FirstName,
		LastName:   p.LastName,
		DomainLess: false,
		Status:     aaa.ActiveUserStatus,
	}

	corp := &aaa.Corporation{}
	corp.LegalName = p.Company

	role, err := aaa.NewAaaManager().FindRoleByName(ucfg.DefaultOwnerRole.String())
	if err != nil {
		return nil, errors.OwnerRoleNotFound
	}

	err = createWhiteLabel(user, corp, newDomain, role)
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
