package main

import (
	"strings"
	"time"

	"clickyab.com/crab/cmd"
	"clickyab.com/crab/modules/dataworker/workers"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/ogier/pflag"
	"github.com/sirupsen/logrus"
)

var (
	keyFields    = pflag.String("keyfields", "campaign_id", "key field is a field that we group data wit it. comma seprated")
	targetFields = pflag.String("targetfields", "cpc,cpm", "target fields should update comma seprated")
	findQuery    = pflag.String("query", "-", "main query that find result")
	targetTable  = pflag.String("targetTable", "-", "target table for insert or update")
)

func main() {
	config.Initialize(cmd.Organization, cmd.Application, cmd.Prefix)
	defer initializer.Initialize()()

	pflag.Parse()
	now := time.Now()

	logrus.Debug("#######################################################################")
	logrus.Debugf("run dataworker command to update data at %s \n", now.Format("2006/01/02 00:00:00"))
	logrus.Debugf("find query is: %s \n", *findQuery)
	logrus.Debugf("key fields are: %s \n", *keyFields)
	logrus.Debugf("and target fields to update are: %s \n", *targetFields)
	logrus.Debugf("and target table: %s \n", *targetTable)

	tf := strings.Split(*targetFields, ",")
	kys := strings.Split(*keyFields, ",")
	workers.Run(*findQuery, *targetTable, kys, tf)

	logrus.Debug("#######################################################################")
}
