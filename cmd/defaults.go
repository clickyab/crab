package cmd

import (
	_ "github.com/clickyab/services/mysql/connection/mysql" // import mysql
	_ "github.com/clickyab/services/redis"                  // the redis must be there
)

const (
	// Organization is the organization name
	Organization = "clickyab"
	// Application name
	Application = "crab"
	// Prefix is the config prefix
	Prefix = "CRB"
)
