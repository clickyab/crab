package main

import (
	"clickyab.com/crab/cmd"
	"github.com/Sirupsen/logrus"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/shell"
)

func main() {
	config.Initialize(cmd.Organization, cmd.Application, cmd.Prefix)
	defer initializer.Initialize()()

	sig := shell.WaitExitSignal()
	logrus.Infof("Sig %s received, exiting...", sig)
}
