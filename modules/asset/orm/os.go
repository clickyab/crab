package orm

import (
	"fmt"
	"time"

	"strings"

	"github.com/clickyab/services/assert"
)

// OS os model in database
// @Model {
//		table = oses
//		primary = true, id
//		find_by = name,id
//		list = yes
// }
type OS struct {
	ID        int64       `json:"id" db:"id"`
	Name      string      `json:"name" db:"name"`
	Status    AssetStatus `json:"status" db:"status"`
	Mobile    bool        `json:"mobile" db:"mobile"`
	Desktop   bool        `json:"desktop" db:"desktop"`
	Tablet    bool        `json:"tablet" db:"tablet"`
	Other     bool        `json:"other" db:"other"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}

// OSBrowser OSBrowser model in database
// @Model {
//		table = os_browser
//		primary = false, browser_id,os_id
//		list = yes
// }
type OSBrowser struct {
	BrowserID int64 `json:"browser_id" db:"browser_id"`
	OSID      int64 `json:"os_id" db:"os_id"`
}

// OSList with valid browsers
type OSList struct {
	OS
	ValidBrowsers      string   `json:"-" db:"valid_browsers"`
	ValidBrowsersArray []string `json:"valid_browsers" db:"-"`
}

// ListOSWithBrowser try to list all OS with browsers
func (m *Manager) ListOSWithBrowser() []OSList {

	var res []OSList

	q := fmt.Sprintf(`SELECT %s,GROUP_CONCAT(b.name) AS valid_browsers FROM %s AS o
	INNER JOIN %s ob on o.id = ob.os_id
	INNER JOIN %s AS b on ob.browser_id = b.id
	WHERE o.status=? GROUP BY o.id`,
		GetSelectFields(OSTableFull, "o"),
		OSTableFull,
		OSBrowserTableFull,
		BrowserTableFull,
	)

	_, err := m.GetRDbMap().Select(&res, q, EnableAssetStatus)
	assert.Nil(err)

	for i := range res {
		res[i].ValidBrowsersArray = strings.Split(res[i].ValidBrowsers, ",")
	}

	return res
}
