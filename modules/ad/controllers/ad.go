package controllers

import "github.com/clickyab/services/framework/controller"

// Controller is the controller for the ad package
// @Route {
//		group = /ad
//		middleware = domain.Access
// }
type Controller struct {
	controller.Base
}
