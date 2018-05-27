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

	TotalBudget  int64 `json:"total_budget" db:"total_budget"`
	DailyBudget  int64 `json:"daily_budget" db:"daily_budget"`
	WebNativeCPC int64 `json:"web_native_cpc" db:"web_native_cpc"`
	WebBannerCPC int64 `json:"web_banner_cpc" db:"web_banner_cpc"`
	WebVastCPC   int64 `json:"web_vast_cpc" db:"web_vast_cpc"`
	AppNativeCPC int64 `json:"app_native_cpc" db:"app_native_cpc"`
	AppBannerCPC int64 `json:"app_banner_cpc" db:"app_banner_cpc"`
	AppVastCPC   int64 `json:"app_vast_cpc" db:"app_vast_cpc"`
	WebCPC       int64 `json:"web_cpc" db:"web_cpc"`
	AppCPC       int64 `json:"app_cpc" db:"app_cpc"`
	WebNativeCPM int64 `json:"web_native_cpm" db:"web_native_cpm"`
	WebBannerCPM int64 `json:"web_banner_cpm" db:"web_banner_cpm"`
	WebVastCPM   int64 `json:"web_vast_cpm" db:"web_vast_cpm"`
	AppNativeCPM int64 `json:"app_native_cpm" db:"app_native_cpm"`
	AppBannerCPM int64 `json:"app_banner_cpm" db:"app_banner_cpm"`
	AppVastCPM   int64 `json:"app_vast_cpm" db:"app_vast_cpm"`
	WebCPM       int64 `json:"web_cpm" db:"web_cpm"`
	AppCPM       int64 `json:"app_cpm" db:"app_cpm"`

	Advantage int `json:"advantage" db:"advantage"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// DomainUser domain_user model in database
// @Model {
//		table = users_domains
//		primary = false, user_id, domain_id
// }
type DomainUser struct {
	DomainID int64        `json:"domain_id" db:"domain_id"`
	Status   DomainStatus `json:"status" db:"status"`
	UserID   int64        `json:"user_id" db:"user_id"`
}

// FindActiveDomainByName find active domain by name
func (m *Manager) FindActiveDomainByName(name string) (*Domain, error) {
	var res Domain
	q := fmt.Sprintf("SELECT %s FROM %s WHERE domain_base=? AND status=?", getSelectFields(DomainTableFull, ""), DomainTableFull)
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
		"WHERE u.email=? AND d.status=?", getSelectFields(DomainTableFull, "d"), DomainTableFull, DomainUserTableFull, "users")
	_, err := m.GetRDbMap().Select(&res, q, e, EnableDomainStatus)
	assert.Nil(err)
	return res
}
