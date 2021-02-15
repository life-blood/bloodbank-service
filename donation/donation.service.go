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

const donationsPath = "donations"

// SetupRoutes :
func SetupRoutes(apiBasePath string) {
	donationsHandler := http.HandlerFunc(handleDonations)
	donationHandler := http.HandlerFunc(handleDonation)
	http.Handle(fmt.Sprintf("%s/%s", apiBasePath, donationsPath), cors.Middleware(donationsHandler))
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, donationsPath), cors.Middleware(donationHandler))
}

func handleDonations(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// Filter by usedId if /donations?userId='ID'
		param1 := req.URL.Query().Get("userId")
		donationList, err := getDonationList(param1)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		j, err := json.Marshal(donationList)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(j)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var donation Donation
		err := json.NewDecoder(req.Body).Decode(&donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = insertDonation(donation)
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
		donation, err := getDonation(donationID)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		if donation == nil {
			w.WriteHeader(http.StatusNotFound)
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
			log.Fatal(err)
		}

	case http.MethodPost:
		var donation Donation
		err := json.NewDecoder(r.Body).Decode(&donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if donation.DonationID != donationID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = updateDonation(donation)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		log.Printf("Removing donation with ID %d", donationID)
		removeDonation(donationID)

	case http.MethodOptions:
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
