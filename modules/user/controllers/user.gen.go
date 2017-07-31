// Code generated build with router DO NOT EDIT.

package user

import (
	"sync"

	"github.com/clickyab/services/framework/middleware"
	"github.com/clickyab/services/framework/router"
	"github.com/clickyab/services/initializer"
	"github.com/rs/xmux"
)

var once = sync.Once{}

// Routes return the route registered with this
func (ctrl *Controller) Routes(r *xmux.Mux, mountPoint string) {
	once.Do(func() {

		initializer.DoInitialize(ctrl)
	})
}

func init() {
	router.Register(&Controller{})
}
