// Code generated build with router DO NOT EDIT.

package controllers

import (
	"sync"

	"github.com/clickyab/services/framework"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
)

var once = sync.Once{}

// Routes return the route registered with this
func (ctrl *Controller) Routes(r framework.Mux) {
	once.Do(func() {

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
