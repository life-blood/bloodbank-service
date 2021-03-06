package main

import (
	"bloodbankservice/donation"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"bloodbankservice/database"
)

//bloodbank service base URL path
const apiBasePath = "/api"

//Get the Port from the environment so we the service can be run on Heroku
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "5001"
		log.Println("No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func main() {
	//load .env file from given path
	//we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Create database connection
	db, err := database.GetDbConnection()
	if err != nil {
		log.Fatal(err)
	}

	//Initialize donation service
	donationRepo := donation.NewDonationMySQL(db)
	donation.DonationService = donation.NewService(donationRepo)
	donation.SetupRoutes(apiBasePath)

	//Setup the backend service
	port := getPort()
	err = http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
