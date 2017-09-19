package main

import (
	"database/sql"
	"fmt"
	"os"

	"clickyab.com/crab/cmd"
	"github.com/clickyab/services/config"
	"github.com/clickyab/services/initializer"
	"github.com/clickyab/services/migration"
	"github.com/clickyab/services/mysql"
	"github.com/ogier/pflag"
	"github.com/sirupsen/logrus"
)

var (
	action = pflag.String("action", "up", "up/down is supported, default is up")
	n      int
)

type migrationManager struct {
}

func (migrationManager) GetSQLDB() *sql.DB {
	m := &mysql.Manager{}
	return m.GetWSQLDB()
}

func (migrationManager) GetDialect() string {
	return "mysql"
}

func main() {
	config.Initialize(cmd.Organization, cmd.Application, cmd.Prefix)
	defer initializer.Initialize()()

	pflag.Parse()
	var err error
	m := &migrationManager{}
	if *action == "up" {
		n, err = migration.Do(m, migration.Up, 0)
		fmt.Printf("\n\n%d migration is applied\n", n)
	} else if *action == "down" {
		n, err = migration.Do(m, migration.Down, 1)
		fmt.Printf("\n\n%d migration is applied\n", n)
	} else if *action == "down-all" {
		n, err = migration.Do(m, migration.Down, 0)
		fmt.Printf("\n\n%d migration is applied\n", n)
	} else if *action == "redo" {
		n, err = migration.Do(m, migration.Down, 1)
		if err == nil {
			n, err = migration.Do(m, migration.Up, 1)
		}
		fmt.Printf("\n\n%d migration is applied\n", n)

	} else if *action == "list" {
		migration.List(m, os.Stdout)
	}

	if err != nil {
		logrus.Fatal(err)
	}
}
