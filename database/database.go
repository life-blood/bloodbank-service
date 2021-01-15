package database

import (
	"database/sql"
	"log"
)

var DbConnection *sql.DB

func SetupDatabase() {
	var err error
	DbConnection, err = sql.Open("mysql", "root:passwordb@tcp(127.0.0.1:3306)/bloodbankdb")
	if err != nil {
		log.Fatal(err)
	}
}
