package mailer

import (
	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/notification"
)

var (
	defMail = config.RegisterString("crab.misc.mailer.sender.mail", "hi@clickyab.com", "mail to send data from")
	defFrom = config.RegisterString("crab.misc.mailer.sender.name", "clickyab", "mail to send data from")
)

// SendMail is the wrapper over the send mail
func SendMail(usr *aaa.User, subject, msg string) {
	to := notification.Duet{
		Name:    usr.FirstName + " " + usr.LastName,
		Contact: usr.Email,
	}

	from := notification.Duet{
		Name:    defFrom.String(),
		Contact: defMail.String(),
	}

	notification.Send(notification.MailType, subject, msg, from, to)
}
