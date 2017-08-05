package mail

import (
	"net/smtp"

	"fmt"

	"github.com/clickyab/services/assert"
)

// todo: shohuld subtitude to with client.mail
// Send sends Email to client
func Send(subject string, msg string, to string) {
	msg = fmt.Sprintf("Subject: %s\r\n"+
		"From: %s\r\n"+
		"To: %s\r\n"+
		"%s", subject, from.String(), to, msg)

	err := smtp.SendMail(
		smtpAddress+":"+smtpAddressPort,
		auth,
		from.String(),
		[]string{to},
		[]byte(msg),
	)

	assert.Nil(err)
}
