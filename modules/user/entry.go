package user

import (
	_ "clickyab.com/crab/modules/user/controllers"      // import init
	_ "clickyab.com/crab/modules/user/middleware/authz" // import init
	_ "clickyab.com/crab/modules/user/models"           // import init
	_ "clickyab.com/crab/modules/user/ucfg"             // import init
)
