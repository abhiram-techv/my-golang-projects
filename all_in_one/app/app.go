package app

import (
	"myproject/config"
	"myproject/repo/driver"
	"myproject/utils"
)

func Run(config config.Config) {
	db, err := driver.Connect(config)
	if err != nil {
		utils.Logging("failed to connect db", err)
	}
	Router(config, db)
}
