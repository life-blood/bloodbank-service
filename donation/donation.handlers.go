package donation

import (
	"bloodbankservice/cors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Donation url path used to retrieve and edit data
const donationsPath = "donations"

//DonationService provide donation functionalities
var DonationService *Service

//SetupRoutes the basic supported routes are
// - api/donations/ - Operations for multiple donations
// - api/donation/ - Operations for single a donation
// all routes are with eanbled CORS
func SetupRoutes(apiBasePath string) {
	donationsHandler := http.HandlerFunc(handleDonations)
	donationHandler := http.HandlerFunc(handleDonation)

	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, donationsPath), cors.Middleware(donationsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, donationsPath), cors.Middleware(donationHandler))
}

//Handle operations for multiple Donation entities
func handleDonations(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// Filter by userId if /donations?userId='ID'
		param1 := req.URL.Query().Get("userId")
		donationList, err := DonationService.GetDonationList(param1)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Print(err)
			return
		}
		j, err := json.Marshal(donationList)
		if err != nil {
			log.Print(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Print(err)
		}
	case http.MethodPost:
		var donation Donation
		err := json.NewDecoder(req.Body).Decode(&donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = DonationService.CreateDonation(&donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

//Handle operations for single Donation entity
func handleDonation(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", donationsPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	donationID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		log.Printf("Retrieving donation by id %d", donationID)
		donation, err := DonationService.GetDonation(donationID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		j, err := json.Marshal(donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(j)
		if err != nil {
			log.Print(err)
		}
	case http.MethodPost:
		var donation Donation
		err := json.NewDecoder(r.Body).Decode(&donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = DonationService.UpdateDonation(&donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		log.Printf("Removing donation with ID %d", donationID)
		err = DonationService.DeleteDonation(donationID)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
