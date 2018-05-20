package aaa

import (
	"time"

	"clickyab.com/crab/modules/user/errors"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"gopkg.in/gorp.v2"
)

// AuditActionType is audit action type
// @Enum{
// }
type AuditActionType string

const (
	// InsertNewData action
	InsertNewData AuditActionType = "insert"
	// UpdateData data action
	UpdateData AuditActionType = "update"
	// DeleteData data
	DeleteData AuditActionType = "delete"
)

// AuditLog audit logs model in database
// @Model {
//		table = audit_logs
//		primary = true, id
//		find_by = id
//		list = yes
// }
type AuditLog struct {
	ID             int64                `json:"id" db:"id"`
	DomainID       int64                `json:"domain_id" db:"domain_id"`
	UserID         int64                `json:"user_id" db:"user_id"`
	UserPerm       string               `json:"user_perm" db:"user_perm"`
	PermScope      permission.UserScope `json:"perm_scope" db:"perm_scope"`
	Action         AuditActionType      `json:"action" db:"action"`
	TargetModel    string               `json:"target_model" db:"target_model"`
	TargetID       int64                `json:"target_id" db:"target_id"`
	OwnerID        int64                `json:"owner_id" db:"owner_id"`
	Impersonate    bool                 `json:"impersonate" db:"impersonate"`
	ImpersonatorID mysql.NullInt64      `json:"impersonator_id" db:"impersonator_id"`
	Description    mysql.NullString     `json:"description" db:"description"`
	CreatedAt      time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt      mysql.NullTime       `json:"updated_at" db:"updated_at"`
}

// AuditExtraData to embed in models and set all needed data
type AuditExtraData struct {
	domainID       int64                  `json:"-" db:"-"`
	userPerm       string                 `json:"-" db:"-"`
	permScope      permission.UserScope   `json:"-" db:"-"`
	userID         int64                  `json:"-" db:"-"`
	targetModel    string                 `json:"-" db:"-"`
	targetID       int64                  `json:"-" db:"-"`
	ownerID        int64                  `json:"-" db:"-"`
	impersonate    bool                   `json:"-" db:"-"`
	impersonatorID int64                  `json:"-" db:"-"`
	description    string                 `json:"-" db:"-"`
	data           map[string]interface{} `json:"-" db:"-"`
}

// ValidateAuditData validate data
func ValidateAuditData(data *AuditExtraData) error {
	err := preInsertValidate(data)
	if err != nil {
		return err
	}

	if data.targetID < 1 {
		return errors.InalidAuditTargetID
	}
	if len(data.targetModel) < 2 {
		return errors.InalidAuditTargetModel
	}

	return nil
}

func preInsertValidate(data *AuditExtraData) error {
	if data.domainID < 1 {
		return errors.InalidAuditDomainID
	}
	if data.userPerm == "" {
		return errors.InalidAuditPerm
	}
	if !data.permScope.IsValid() {
		return errors.InalidAuditPermScope
	}
	if data.userID < 1 {
		return errors.InalidAuditUserID
	}
	if data.ownerID < 1 {
		return errors.InalidAuditOwnerID
	}
	if len(data.data) == 0 {
		return errors.InalidAuditDetail
	}

	return nil
}

// PreInsert to validate audit data
func (data *AuditExtraData) PreInsert(s gorp.SqlExecutor) error {
	return preInsertValidate(data)
}

// PreUpdate to validate audit data
func (data *AuditExtraData) PreUpdate(s gorp.SqlExecutor) error {
	return ValidateAuditData(data)
}

// PreDelete to validate audit data
func (data *AuditExtraData) PreDelete(s gorp.SqlExecutor) error {
	return ValidateAuditData(data)
}

// PostInsert do insert audit log
func (data *AuditExtraData) PostInsert(s gorp.SqlExecutor) error {
	return AddAuditLog(data, "insert")
}

// PostUpdate post update funcs of gorp to log in audit
func (data *AuditExtraData) PostUpdate(s gorp.SqlExecutor) error {
	return AddAuditLog(data, "update")
}

// PostDelete post delete funcs of gorp to log in audit
func (data *AuditExtraData) PostDelete(s gorp.SqlExecutor) error {
	return AddAuditLog(data, "delete")
}

// SetAuditDomainID to set domain id of entity
func (data *AuditExtraData) SetAuditDomainID(dID int64) error {
	if dID < 1 {
		return errors.InalidAuditDomainID
	}

	data.domainID = dID

	return nil
}

// SetAuditUserData to set user data that are edit/insert/delete entity
func (data *AuditExtraData) SetAuditUserData(userID int64, token string, domainID int64, uPerm string, uScope permission.UserScope) error {
	if userID < 1 {
		return errors.InalidAuditUserID
	}
	if uPerm == "" {
		return errors.InalidAuditPerm
	}
	if !uScope.IsValid() {
		return errors.InalidAuditPermScope
	}

	impersonated := false
	var impersonatorID int64
	impersonatorToken := ImpersonatorToken(token)

	if impersonatorToken != "" {
		impersonated = true
		id, err := ExtractUserID(impersonatorToken)
		if err != nil {
			return err
		}
		impersonatorID = id
	}

	data.userID = userID
	data.impersonate = impersonated
	data.impersonatorID = impersonatorID
	data.userPerm = uPerm
	data.permScope = uScope
	return nil
}

// SetAuditDescribe to set describe of action
func (data *AuditExtraData) SetAuditDescribe(d map[string]interface{}, description string) error {
	if len(d) == 0 {
		return errors.InalidAuditDetail
	}

	data.data = d
	data.description = description

	return nil
}

// SetAuditOwnerID to set owner id of entity
func (data *AuditExtraData) SetAuditOwnerID(oID int64) error {
	if oID < 1 {
		return errors.InalidAuditOwnerID
	}

	data.ownerID = oID

	return nil
}

// SetAuditEntity to set data of entity
func (data *AuditExtraData) SetAuditEntity(model string, id int64) error {
	if id < 1 {
		return errors.InalidAuditTargetID
	}
	if len(model) < 2 {
		return errors.InalidAuditTargetModel
	}

	data.targetModel = model
	data.targetID = id

	return nil
}

// AddAuditLog add new audit log
func AddAuditLog(data *AuditExtraData, action string) error {
	var actionType AuditActionType
	switch action {
	case "insert":
		actionType = InsertNewData
	case "update":
		actionType = UpdateData
	case "delete":
		actionType = DeleteData
	default:
		return errors.InalidAuditAction
	}

	impersonatorID := mysql.NullInt64{Int64: 0, Valid: false}
	if data.impersonatorID > 0 {
		impersonatorID.Int64 = data.impersonatorID
		impersonatorID.Valid = true
	}

	audit := AuditLog{
		DomainID:       data.domainID,
		UserID:         data.userID,
		UserPerm:       data.userPerm,
		PermScope:      data.permScope,
		Action:         actionType,
		TargetModel:    data.targetModel,
		TargetID:       data.targetID,
		OwnerID:        data.ownerID,
		Impersonate:    data.impersonate,
		ImpersonatorID: impersonatorID,
		Description:    mysql.NullString{String: data.description, Valid: true},
	}

	db := NewAaaManager()
	err := db.CreateAuditLog(&audit)
	if err != nil {
		return err
	}

	detail := AuditLogDetail{
		AuditLogID: audit.ID,
		Data:       data.data,
	}

	return db.CreateAuditLogDetail(&detail)
}
