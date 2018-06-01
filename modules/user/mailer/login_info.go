package mailer

import (
	"fmt"
	"net/http"
	"net/url"

	"clickyab.com/crab/modules/user/aaa"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/gettext/t9e"
)

var frontWelcomeMailPath = config.RegisterString("crab.modules.user.welcome.mail.front.path", "/user/welcome", "front login and welcome page")

// LoginInfoEmail to send welcome msg and login information to created user
func LoginInfoEmail(u *aaa.User, pass string, r *http.Request) error {
	ul := &url.URL{
		Scheme: framework.Scheme(r),
		Host:   r.Host,
		Path:   frontWelcomeMailPath.String(),
	}

	// TODO: Change email template
	template := fmt.Sprintf(
		`Hi %s %s and welcome to clickyab <br />
		You have new account to manage ads and can login with:<br /><br />
		user: %s<br />
		pass: %s<br />
		<br /><br />
		please click here: %s<br />
		`,
		u.FirstName,
		u.LastName,
		u.Email,
		pass,
		ul.String(),
	)

	SendMail(u, t9e.G("Welcome to Clickyab!").String(), template)
	return nil
}
