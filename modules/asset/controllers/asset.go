package controllers

import "github.com/clickyab/services/framework/controller"

// Controller is the controller for the asset package
// @Route {
// 		middleware = domain.Access
//		group = /asset
// }
type Controller struct {
	controller.Base
}
