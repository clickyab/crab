package main

import (
	_ "clickyab.com/crab/modules/ad"
	_ "clickyab.com/crab/modules/asset"
	_ "clickyab.com/crab/modules/campaign"
	_ "clickyab.com/crab/modules/dataworker"
	_ "clickyab.com/crab/modules/domain"
	_ "clickyab.com/crab/modules/financial"
	_ "clickyab.com/crab/modules/inventory"
	_ "clickyab.com/crab/modules/location"
	_ "clickyab.com/crab/modules/misc"
	_ "clickyab.com/crab/modules/time"
	_ "clickyab.com/crab/modules/upload"
	_ "clickyab.com/crab/modules/user"
	_ "github.com/clickyab/services/kv/redis"
)