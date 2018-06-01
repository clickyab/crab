package orm

import (
	"time"

	"fmt"

	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
)

// DomainStatus is the user domain status
type (
	// DomainStatus is the user domain status
	// @Enum{
	// }
	DomainStatus string
)

const (
	// EnableDomainStatus enable domain
	EnableDomainStatus DomainStatus = "enable"
	// DisableDomainStatus disable domain
	DisableDomainStatus DomainStatus = "disable"
)

// Domain domain model in database
// @Model {
//		table = domains
//		primary = true, id
//		find_by = id,domain_base
//		list = yes
// }
type Domain struct {
	ID          int64                  `json:"id" db:"id"`
	DomainBase  string                 `json:"domain_base" db:"domain_base"`
	Title       string                 `json:"title" db:"title"`
	Logo        mysql.NullString       `json:"logo" db:"logo"`
	Theme       string                 `json:"theme" db:"theme"`
	Description mysql.NullString       `json:"description" db:"description"`
	Attributes  mysql.GenericJSONField `json:"attributes" db:"attributes"`
	Status      DomainStatus           `json:"status" db:"status"`

	MinTotalBudget  int64 `json:"min_total_budget" db:"min_total_budget"`
	MinDailyBudget  int64 `json:"min_daily_budget" db:"min_daily_budget"`
	MinWebNativeCPC int64 `json:"min_web_native_cpc" db:"min_web_native_cpc"`
	MinWebBannerCPC int64 `json:"min_web_banner_cpc" db:"min_web_banner_cpc"`
	MinWebVastCPC   int64 `json:"min_web_vast_cpc" db:"min_web_vast_cpc"`
	MinAppNativeCPC int64 `json:"min_app_native_cpc" db:"min_app_native_cpc"`
	MinAppBannerCPC int64 `json:"min_app_banner_cpc" db:"min_app_banner_cpc"`
	MinAppVastCPC   int64 `json:"min_app_vast_cpc" db:"min_app_vast_cpc"`
	MinWebCPC       int64 `json:"min_web_cpc" db:"min_web_cpc"`
	MinAppCPC       int64 `json:"min_app_cpc" db:"min_app_cpc"`
	MinWebNativeCPM int64 `json:"min_web_native_cpm" db:"min_web_native_cpm"`
	MinWebBannerCPM int64 `json:"min_web_banner_cpm" db:"min_web_banner_cpm"`
	MinWebVastCPM   int64 `json:"min_web_vast_cpm" db:"min_web_vast_cpm"`
	MinAppNativeCPM int64 `json:"min_app_native_cpm" db:"min_app_native_cpm"`
	MinAppBannerCPM int64 `json:"min_app_banner_cpm" db:"min_app_banner_cpm"`
	MinAppVastCPM   int64 `json:"min_app_vast_cpm" db:"min_app_vast_cpm"`
	MinWebCPM       int64 `json:"min_web_cpm" db:"min_web_cpm"`
	MinAppCPM       int64 `json:"min_app_cpm" db:"min_app_cpm"`

	Advantage int `json:"advantage" db:"advantage"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// DomainUser domain_user model in database
// @Model {
//		table = users_domains
//		primary = true,id
//		list = true
// }
type DomainUser struct {
	ID       int64           `json:"id" db:"id"`
	RoleID   int64           `json:"role_id" db:"role_id"`
	DomainID mysql.NullInt64 `json:"domain_id" db:"domain_id"`
	Status   DomainStatus    `json:"status" db:"status"`
	UserID   int64           `json:"user_id" db:"user_id"`
}

// FindActiveDomainByName find active domain by name
func (m *Manager) FindActiveDomainByName(name string) (*Domain, error) {
	var res Domain
	q := fmt.Sprintf("SELECT %s FROM %s WHERE domain_base=? AND status=?", GetSelectFields(DomainTableFull, ""), DomainTableFull)
	err := m.GetRDbMap().SelectOne(&res, q, name, EnableDomainStatus)
	if err != nil {
		return nil, err
	}
	return &res, nil

}

// FindUserDomainsByEmail find active user domain based on its email
func (m *Manager) FindUserDomainsByEmail(e string) []Domain {
	var res []Domain
	q := fmt.Sprintf("SELECT %s FROM %s AS d "+
		"INNER JOIN %s AS dm ON dm.domain_id=d.id "+
		"INNER JOIN %s AS u ON u.id=dm.user_id "+
		"WHERE u.email=? AND d.status=?", GetSelectFields(DomainTableFull, "d"), DomainTableFull, DomainUserTableFull, "users")
	_, err := m.GetRDbMap().Select(&res, q, e, EnableDomainStatus)
	assert.Nil(err)
	return res
}

// FindActiveUserDomainByUserDomain FindActiveUserDomainByUserDomain
func (m *Manager) FindActiveUserDomainByUserDomain(userID, domainID int64) (*DomainUser, error) {
	var res *DomainUser
	q := fmt.Sprintf(`SELECT %s FROM %s AS du 
	INNER JOIN %s AS d ON (du.domain_id=d.id)
	WHERE du.user_id=? AND du.domain_id=? AND d.status=? AND du.status=?`,
		GetSelectFields(DomainUserTableFull, "du"),
		DomainUserTableFull,
		DomainTableFull,
	)
	err := m.GetRDbMap().SelectOne(&res, q, userID, domainID, EnableDomainStatus, EnableDomainStatus)
	if err != nil {
		return nil, err
	}
	return res, nil
}
