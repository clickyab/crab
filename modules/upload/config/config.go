package config

import "github.com/clickyab/services/config"

var Path = config.RegisterString("crab.modules.upload.path", "/upload", "a path to the location that uploaded file should save")
