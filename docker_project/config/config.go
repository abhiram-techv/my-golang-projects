package config

import (
	"database/sql"
	"docker_project/model"
	"encoding/json"
	"io/ioutil"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func LoadApiConfig(filename string) (model.ApiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return model.ApiConfigData{}, err
	}
	var c model.ApiConfigData

	err = json.Unmarshal(bytes, &c)

	if err != nil {
		return model.ApiConfigData{}, err
	}

	return c, nil
}

func Connect() (*sql.DB, error) {

	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		return nil, err
	}
	return db, nil
}
