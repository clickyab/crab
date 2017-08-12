package ucfg

import (
	"time"

	"github.com/clickyab/services/config"
)

var TokenTimeout = config.RegisterDuration("crab.modules.user.token_timeout", time.Hour*72, "token expiry time")
var DefaultRole = config.RegisterString("crab.modules.user.default_role", "Advertiser", "default role in system")
