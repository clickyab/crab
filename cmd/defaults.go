package cmd

import (
	_ "github.com/clickyab/services/kv/redis"               // the redis must be there
	_ "github.com/clickyab/services/mysql/connection/mysql" // import mysql
)

const (
	// Organization is the organization name
	Organization = "clickyab"
	// Application name
	Application = "crab"
	// Prefix is the config prefix
	Prefix = "CRB"
)
