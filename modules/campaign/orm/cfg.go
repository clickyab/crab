package orm

import "github.com/clickyab/services/config"

var (
	//defaultWebBannerCPM = config.RegisterInt64("crab.modules.campaigns.defaults.web.banner.cpm", 0, "default web banner cpm")
	defaultWebBannerCPC = config.RegisterInt64("crab.modules.campaigns.defaults.web.banner.cpc", 250, "default web banner cpc")
	//defaultWebBannerCPA = config.RegisterInt64("crab.modules.campaigns.defaults.web.banner.cpa", 0, "default web banner cpa")

	//defaultWebVastCPM = config.RegisterInt64("crab.modules.campaigns.defaults.web.vast.cpm", 0, "default web vast cpm")
	defaultWebVastCPC = config.RegisterInt64("crab.modules.campaigns.defaults.web.vast.cpc", 200, "default web vast cpc")
	//defaultWebVastCPA = config.RegisterInt64("crab.modules.campaigns.defaults.web.vast.cpa", 0, "default web vast cpa")

	//defaultWebNativeCPM = config.RegisterInt64("crab.modules.campaigns.defaults.web.native.cpm", 0, "default web native cpm")
	defaultWebNativeCPC = config.RegisterInt64("crab.modules.campaigns.defaults.web.native.cpc", 150, "default web native cpc")
	//defaultWebNativeCPA = config.RegisterInt64("crab.modules.campaigns.defaults.web.native.cpa", 0, "default web native cpa")

	//defaultAppBannerCPM = config.RegisterInt64("crab.modules.campaigns.defaults.app.banner.cpm", 0, "default app banner cpm")
	defaultAppBannerCPC = config.RegisterInt64("crab.modules.campaigns.defaults.app.banner.cpc", 700, "default app banner cpc")
	//defaultAppBannerCPA = config.RegisterInt64("crab.modules.campaigns.defaults.app.banner.cpa", 0, "default app banner cpa")

	//defaultAppVastCPM = config.RegisterInt64("crab.modules.campaigns.defaults.app.vast.cpm", 0, "default app vast cpm")
	defaultAppVastCPC = config.RegisterInt64("crab.modules.campaigns.defaults.app.vast.cpc", 700, "default app vast cpc")
	//defaultAppVastCPA = config.RegisterInt64("crab.modules.campaigns.defaults.app.vast.cpa", 0, "default app vast cpa")

	//defaultAppNativeCPM = config.RegisterInt64("crab.modules.campaigns.defaults.app.native.cpm", 0, "default app native cpm")
	defaultAppNativeCPC = config.RegisterInt64("crab.modules.campaigns.defaults.app.native.cpc", 700, "default app native cpc")
	//defaultAppNativeCPA = config.RegisterInt64("crab.modules.campaigns.defaults.app.native.cpa", 0, "default app native cpa")
)
