package database

import (
	"database/sql"
	"log"
)

var DbConnection *sql.DB

func SetupDatabase() {
	var err error
	DbConnection, err = sql.Open("mysql", "b2d8489953993a:3f02cb75@tcp(eu-cdbr-west-03.cleardb.net)/heroku_4fa9ce9da159b8c")
	if err != nil {
		log.Fatal(err)
	}
}
