package main

import (
	"github.com/aceberg/WMService/conf"
	"github.com/aceberg/WMService/db"
	"github.com/aceberg/WMService/web"
)

func main() {
	appConfig := conf.GetConfig()

	db.CreateDB(appConfig.DbPath)

	web.Webgui(appConfig)
}
