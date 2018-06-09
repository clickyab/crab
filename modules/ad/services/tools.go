package services

import (
	"fmt"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/mailer"
	"github.com/clickyab/services/assert"
	"github.com/clickyab/services/mysql"
	"github.com/clickyab/services/permission"
	"github.com/fatih/structs"
)

// ChangeCreativesStatus to change creatives status
func ChangeCreativesStatus(req []orm.ChangeStatusReq,
	userID, domainID int64, userToken string, scope permission.UserScope) error {

	m := orm.NewOrmManager()
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	for _, re := range req {
		err := re.Creative.SetAuditUserData(userID, userToken, domainID, "change_creative_status", scope)
		if err != nil {
			return err
		}
		d := structs.Map(re.Creative)
		err = re.Creative.SetAuditDescribe(d, "change creative status")
		if err != nil {
			return err
		}
		err = re.Creative.SetAuditOwnerID(userID)
		if err != nil {
			return err
		}
		err = re.Creative.SetAuditEntity("creative", re.CreativeID)
		if err != nil {
			return err
		}
		if re.Status == orm.AcceptedCreativeStatus {
			re.Creative.Status = re.Status
		} else if re.Status == orm.RejectedCreativeStatus {
			re.Creative.Status = re.Status
			re.Creative.RejectReasonsID = mysql.NullInt64{Valid: re.RejectReasonID != 0, Int64: re.RejectReasonID}
		}
		re.Creative.Status = re.Status
		err = m.UpdateCreative(re.Creative)
		if err != nil {
			return err
		}
	}
	return nil
}

// AcceptCreatives accept creative
func AcceptCreatives(creatives []*orm.Creative, userID, domainID int64, userToken string, scope permission.UserScope) error {
	m := orm.NewOrmManager()
	err := m.Begin()
	assert.Nil(err)
	defer func() {
		if err != nil {
			assert.Nil(m.Rollback())
		} else {
			assert.Nil(m.Commit())
		}
	}()
	for _, creative := range creatives {
		err = creative.SetAuditDomainID(domainID)
		if err != nil {
			return err
		}
		err = creative.SetAuditOwnerID(creative.UserID)
		if err != nil {
			return err
		}
		d := structs.Map(creative)
		err = creative.SetAuditDescribe(d, "accept campaign creative status")
		if err != nil {
			return err
		}
		err = creative.SetAuditUserData(userID, userToken, domainID, "change_creative_status", scope)
		if err != nil {
			return err
		}
		err = creative.SetAuditEntity("creative", creative.ID)
		if err != nil {
			return err
		}
		creative.Status = orm.AcceptedCreativeStatus
		err = m.UpdateCreative(creative)
		if err != nil {
			return err
		}
	}
	return nil
}

// SendChangeStatusMessage send change status email
func SendChangeStatusMessage(ids []int64, campaignID int64) error {
	m := orm.NewOrmManager()
	creatives, err := m.GetCreativeWithIDRange(ids, campaignID)
	if err != nil {
		return err
	}
	err = sendChangeStatusMessage(creatives)
	if err != nil {
		return errors.SendNotifyEmailErr
	}
	return nil
}

func sendChangeStatusMessage(creatives []*orm.CreativeWithRelation) error {
	var msg string
	var userMessage string
	// TODO change email messages
	for _, creative := range creatives {
		if creative.CreativeStatus == orm.AcceptedCreativeStatus {
			msg = fmt.Sprintf("Your creative with name:'%s' approved", creative.Name)
		} else if creative.CreativeStatus == orm.RejectedCreativeStatus {
			msg = fmt.Sprintf("Your creative with name '%s' rejected because of '%s'", creative.Name, creative.RejectReasonReason.String)
		}
		userMessage += msg + "<br>"
	}
	userObj := creatives[0]
	// send emails
	tempUser := &aaa.User{
		Email:     userObj.UserEmail,
		FirstName: userObj.UserFirstName,
		LastName:  userObj.UserLastName,
	}
	mailer.SendMail(tempUser, "Your Creative Check Result", userMessage)
	return nil
}
