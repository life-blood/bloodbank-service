package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

//Configured from .env configuration filed
const (
	dbUser = "DB_USER"
	dbPass = "DB_PASS"
	dbPort = "DB_PORT"
	dbName = "DB_NAME"
)

//GetDbConnection retrieve sql datab
func GetDbConnection() (*sql.DB, error) {
	serverName := fmt.Sprintf("localhost:%s", os.Getenv(dbPort))
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		os.Getenv(dbUser),
		os.Getenv(dbPass),
		serverName, os.Getenv(dbName))
	log.Printf("Trying to connect to database with %s", connectionString)
	return sql.Open("mysql", connectionString)
}
