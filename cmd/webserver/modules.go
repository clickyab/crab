package main

import (
	_ "clickyab.com/crab/modules/domain"
	_ "clickyab.com/crab/modules/misc"
	_ "clickyab.com/crab/modules/user"
	_ "github.com/clickyab/services/store/redis"
	_ "github.com/clickyab/services/eav/redis"
)
