package controllers

import "github.com/clickyab/services/framework/controller"

// Controller is the controller for the location package
// @Route {
// 		middleware = domain.Access
//		group = /time
// }
type Controller struct {
	controller.Base
}
