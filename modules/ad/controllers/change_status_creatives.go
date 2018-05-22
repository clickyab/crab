package controllers

import (
	"context"
	"net/http"

	"fmt"

	"strconv"

	"clickyab.com/crab/modules/ad/errors"
	"clickyab.com/crab/modules/ad/orm"
	orm3 "clickyab.com/crab/modules/campaign/orm"
	"clickyab.com/crab/modules/domain/middleware/domain"
	orm2 "clickyab.com/crab/modules/domain/orm"
	"clickyab.com/crab/modules/user/aaa"
	"clickyab.com/crab/modules/user/mailer"
	"clickyab.com/crab/modules/user/middleware/authz"
	"github.com/clickyab/services/permission"
	"github.com/clickyab/services/xlog"
	"github.com/rs/xmux"
)

// NotifyUserTypes is the notify type to get user this ability to notify user on creative approve or reject
type (
	// NotifyUserTypes is the notify type to get user this ability to notify user on creative approve or reject
	// @Enum{
	// }
	NotifyUserTypes string
)

const (
	// DontNotifyUser no
	DontNotifyUser NotifyUserTypes = "no"
	// DoNotifyUser yes
	DoNotifyUser NotifyUserTypes = "yes"
)

// @Validate{
//}
type changeStatusPayload struct {
	CreativesStatus []orm.ChangeStatusReq `json:"creatives_status" validation:"required"`
	NotifyUser      NotifyUserTypes       `json:"notify_users" validation:"omitempty"`
	currentDomain   *orm2.Domain          `json:"-"`
	currentCampaign *orm3.Campaign        `json:"-"`
}

// ChangeStatusResult return result of bulk changing status of creatives
type ChangeStatusResult struct {
	CreativesStatus []orm.ChangeStatusReq `json:"creatives_status"`
}

func (p *changeStatusPayload) ValidateExtra(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if !p.NotifyUser.IsValid() {
		return errors.InvalidStatusErr
	}
	idInt, err := strconv.ParseInt(xmux.Param(ctx, "id"), 10, 0)
	if err != nil {
		return errors.InvalidIDErr
	}
	dm := domain.MustGetDomain(ctx)
	p.currentDomain = dm
	// find campaign by id domain
	currentCampaign, err := orm3.NewOrmManager().FindCampaignByIDDomain(idInt, dm.ID)
	if err != nil {
		return errors.InvalidIDErr
	}
	p.currentCampaign = currentCampaign
	return nil
}

func sendChangeStatusMessage(sq []orm.ChangeStatusReq, campaignID int64) error {
	var ids []int64
	for _, r := range sq {
		ids = append(ids, r.CreativeID)
	}
	m := orm.NewOrmManager()
	creatives, err := m.GetCreativeWithIDRange(ids, campaignID)
	if err != nil {
		return err
	}
	var msg string
	var userMessage string
	// TODO change email messages
	for _, creative := range creatives {
		if creative.CreativeStatus == orm.AcceptedCreativeStatus {
			msg = fmt.Sprintf("Your creative with name:'%s' approved", creative.Name)
		} else if creative.CreativeStatus == orm.RejectedCreativeStatus {
			msg = fmt.Sprintf("Your creative with name '%s' rejected because of '%s'", creative.Name, creative.RejectReasonReason)
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

// changeCreativesStatus bulk approve reject creatives status of a campaign
// @Rest {
// 		url = /change-creatives-status/:id
//		protected = true
// 		method = put
// 		resource = change_creatives_status:global
// }
func (c *Controller) changeCreativesStatus(ctx context.Context, r *http.Request, p *changeStatusPayload) (*ChangeStatusResult, error) {
	currentUser := authz.MustGetUser(ctx)
	//check permission
	_, ok := aaa.CheckPermOn(currentUser, currentUser, "change_creatives_status", p.currentDomain.ID, permission.ScopeGlobal)
	if !ok {
		return nil, errors.AccessDenied
	}
	m := orm.NewOrmManager()
	// apply approve or reject
	prob := m.ChangeCreativesStatus(p.CreativesStatus, p.currentCampaign.ID)
	if prob != nil {
		xlog.GetWithError(ctx, prob).Debug("database error when change creatives status")
		return nil, errors.UpdateStatusDbErr
	}
	result := ChangeStatusResult{
		CreativesStatus: p.CreativesStatus,
	}
	if p.NotifyUser == DoNotifyUser {
		err := sendChangeStatusMessage(p.CreativesStatus, p.currentCampaign.ID)
		if err != nil {
			xlog.GetWithError(ctx, err).Debug("send notify email for creative status change failed")
			return nil, errors.SendNotifyEmailErr
		}
	}
	return &result, nil
}
