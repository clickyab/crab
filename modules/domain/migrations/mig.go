package migrations

import (
	"github.com/clickyab/services/migration"
)

func init() {
	// Ensure the migrations are included in the final migration
	migration.Register(Asset, AssetDir, "db")
}
