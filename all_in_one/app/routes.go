package app

import (
	"database/sql"
	"fmt"
	"myproject/config"
	"myproject/utils"

	"github.com/labstack/echo/v4"
)

func Router(config config.Config, db *sql.DB) {
	e := echo.New()
	api := e.Group("/api")
	fmt.Println(api)
	err := e.Start(config.Server.Port)
	if err != nil {
		utils.Logging("Failed to start the server", err)
	}
}
