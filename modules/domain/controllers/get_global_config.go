package controllers

import (
	"context"
	"net/http"

	"clickyab.com/crab/modules/domain/orm"
)

// getGlobalConfig get global config
// @Rest {
// 		url = /super-global-config
//		protected = true
// 		method = get
//		resource = get_global_config:superGlobal
// }
func (c *Controller) getGlobalConfig(ctx context.Context, r *http.Request) (*orm.UserConfig, error) {
	sg := &orm.UserConfig{
		MinTotalBudget:  3000,
		MinDailyBudget:  3000,
		MinWebNativeCPC: 120,
		MinWebBannerCPC: 250,
		MinWebVastCPC:   200,
		MinAppNativeCPC: 70,
		MinAppBannerCPC: 0,
		MinAppVastCPC:   0,
		MinWebCPC:       300,
		MinAppCPC:       100,
		MinWebNativeCPM: 2000,
		MinWebBannerCPM: 2000,
		MinWebVastCPM:   2000,
		MinAppNativeCPM: 2000,
		MinAppBannerCPM: 0,
		MinAppVastCPM:   0,
		MinWebCPM:       0,
		MinAppCPM:       0,
	}
	return sg, nil
}
