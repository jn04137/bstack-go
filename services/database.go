package services

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func createConnection() (*sql.DB, error) {
	cfg := mysql.Config {
		User: os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net: "tcp",
		Addr: os.Getenv("DB_HOST_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	return db, err
}
