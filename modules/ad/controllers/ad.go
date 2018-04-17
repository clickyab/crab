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

// Importatn: only add here shared func that use in other package routes
