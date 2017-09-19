package controllers

import "github.com/clickyab/services/framework/controller"

// Controller is the controller for the user package
// @Route {
//		group = /publisher
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}
