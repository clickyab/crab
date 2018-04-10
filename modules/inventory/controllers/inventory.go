package controllers

import (
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/framework/controller"
)

var maxPubInventoryCount = config.RegisterInt("crab.modules.inventory.pub.count", 250, "max publisher in single inventory")

// Controller is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /inventory
// }
type Controller struct {
	controller.Base
}
