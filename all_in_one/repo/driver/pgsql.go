package driver

import (
	"database/sql"
	"fmt"
	"myproject/config"

	_ "github.com/lib/pq"
)

func Connect(config config.Config) (*sql.DB, error) {
	fmt.Println(config)
	db, _ := sql.Open(config.Database.Driver, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.Name))

	defer db.Close()
	err := db.Ping()
	if err != nil {
		return nil, err

	}
	return db, nil
}
