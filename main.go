package main

import (
	"bloodbankservice/donation"
	"log"
	"net/http"
	"os"
)

const apiBasePath = "/api"

func main() {
	port := os.Getenv("PORT")
	donation.SetupRoutes(apiBasePath)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
