package main

import (
	"myproject/app"
	"myproject/config"
)

func main() {
	config := config.LoadConfig()
	app.Run(config)
}
