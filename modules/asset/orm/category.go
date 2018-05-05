package orm

import "github.com/go-sql-driver/mysql"

// Category category model in database
// @Model {
//		table = categories
//		primary = false, name
//		find_by = name
//		list = yes
// }
type Category struct {
	Name        string         `json:"name" db:"name"`
	Description string         `json:"description" db:"description"`
	DeletedAt   mysql.NullTime `json:"deleted_at" db:"deleted_at"`
	Status      AssetStatus    `json:"status" db:"status"`
}

// ModelKind is the ,odel type (eg,campaign,publisher,...)
type (
	// ModelKind is the model kind
	// @Enum{
	// }
	ModelKind string
)

const (
	// CampaignModel campaign model
	CampaignModel ModelKind = "campaign"
	// PublisherModel publisher model
	PublisherModel ModelKind = "publisher"
)

// CategoryModel category model table in database
// @Model {
//		table = category_model
//		primary = false, model_id,category,model
//		list = yes
// }
type CategoryModel struct {
	ModelID  int64     `json:"model_id" db:"model_id"`
	Category string    `json:"category" db:"category"`
	Model    ModelKind `json:"model" db:"model"`
}
