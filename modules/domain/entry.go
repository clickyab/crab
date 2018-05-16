package domain

import (
	_ "clickyab.com/crab/modules/domain/config"            //import init
	_ "clickyab.com/crab/modules/domain/controllers"       //import init
	_ "clickyab.com/crab/modules/domain/middleware/domain" //import init
	_ "clickyab.com/crab/modules/domain/orm"               //import init
)
