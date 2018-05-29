package ucfg

import (
	"time"

	"github.com/clickyab/services/config"
)

// TokenTimeout token expire
var TokenTimeout = config.RegisterDuration("crab.modules.user.token_timeout", time.Hour*72, "token expiry time")

// DefaultRole default role
var DefaultRole = config.RegisterString("crab.modules.user.default_role", "Advertiser", "default role in system")
var DefaultAccountRole = config.RegisterString("crab.modules.user.account.default_role", "Account", "default account role in system")
