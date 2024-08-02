package dbUtil

import (
	"api/internal/configuration"
	"database/sql"
	"log"
	"os"
)

func ConnectToDb() *sql.DB {
	db, err := sql.Open("sqlite3", configuration.AppConfig.Database.Path)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func InitDb() *sql.DB {
	if _, err := os.Stat(configuration.AppConfig.Database.Path); err == nil {
		return nil
	}
	return ConnectToDb()
}
