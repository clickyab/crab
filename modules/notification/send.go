package notification

import "clickyab.com/crab/modules/notification/internal/mail"

type Notiftype int

const (
	SMSType Notiftype = iota
	MailType
)

func Send(subject string, msg string, to string, t Notiftype) {
	switch t {
	case MailType:
		mail.Send(subject, msg, to)
	case SMSType:

	}
}
