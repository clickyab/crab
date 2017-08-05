package mail

import (
	"net/smtp"

	"github.com/clickyab/services/config"
)

var (
	auth smtp.Auth
	from = config.RegisterString("crab.mail_service.from", "clickyab.com", "from part of mail")

	smtpUsername = config.GetString("crab.cmtp.username")
	smtpPassword = config.GetString("crab.cmtp.password")
	smtpHost     = config.GetString("crab.smtp.host")

	smtpAddress     = config.GetString("crab.smtp.address")
	smtpAddressPort = config.GetString("crab.smtp.address_port")
)

func init() {
	auth = smtp.PlainAuth(
		"",
		smtpUsername,
		smtpPassword,
		smtpHost,
	)

}
