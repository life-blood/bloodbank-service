package main

import (
	"bloodbankservice/donation"
	"log"
	"net/http"
)

const apiBasePath = "/api"

func main() {
	donation.SetupRoutes(apiBasePath)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatal(err)
	}
}
