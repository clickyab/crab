package main

import (
	"clickyab.com/crab/cmd"
	"github.com/clickyab/services/config"
	_ "github.com/clickyab/services/fluentd"
	"github.com/clickyab/services/initializer"
	_ "github.com/clickyab/services/kv/redis"
	_ "github.com/clickyab/services/redmine"
	"github.com/clickyab/services/shell"
	_ "github.com/clickyab/services/slack"
	"github.com/sirupsen/logrus"
)

func main() {
	config.Initialize(cmd.Organization, cmd.Application, cmd.Prefix)
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Infof("Sig %s received, exiting...", sig)
}
