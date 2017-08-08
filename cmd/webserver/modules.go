package main

import (
	_ "clickyab.com/crab/modules/domain"
	_ "clickyab.com/crab/modules/location"
	_ "clickyab.com/crab/modules/misc"
	_ "clickyab.com/crab/modules/user"
	_ "github.com/clickyab/services/kv/redis"
	_ "clickyab.com/crab/modules/upload"
)
