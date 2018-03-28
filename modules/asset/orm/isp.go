package orm

import "time"

// ISPKind is the isp type (eg,isp,both,...)
type (
	// ISPKind is the isp kind
	// @Enum{
	// }
	ISPKind string
)

const (
	// BothISPKind both
	BothISPKind ISPKind = "both"
	// CellularISPKind cellular
	CellularISPKind ISPKind = "cellular"
	// ISPISPKind just isp
	ISPISPKind ISPKind = "isp"
)

// ISP isp model in database
// @Model {
//		table = isps
//		primary = false, name
//		find_by = name
//		list = yes
// }
type ISP struct {
	Name      string      `json:"name" db:"name"`
	Kind      ISPKind     `json:"kind" db:"kind"`
	Status    AssetStatus `json:"status" db:"status"`
	CreatedAt time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}
