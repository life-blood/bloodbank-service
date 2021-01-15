package donation

import (
	"bloodbankservice/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"
)

// used to hold donation list in memory
var donationMap = struct {
	sync.RWMutex
	m map[int]Donation
}{m: make(map[int]Donation)}

func init() {
	fmt.Println("Loading available donations")
	donMap, err := loadDonationMap()
	donationMap.m = donMap
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d donations laoded... \n", len(donationMap.m))
}

func loadDonationMap() (map[int]Donation, error) {
	fileName := "donations.json"
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("file [%s] does not exist", fileName)
	}

	file, _ := ioutil.ReadFile(fileName)
	donationList := make([]Donation, 0)
	err = json.Unmarshal([]byte(file), &donationList)
	if err != nil {
		log.Fatal(err)
	}
	donMap := make(map[int]Donation)
	for i := 0; i < len(donationList); i++ {
		donMap[donationList[i].DonationID] = donationList[i]
	}
	return donMap, nil
}

func getDonation(donationID int) *Donation {
	donationMap.RLock()
	defer donationMap.RUnlock()
	if donation, ok := donationMap.m[donationID]; ok {
		return &donation
	}
	return nil
}

func removeDonation(donationID int) {
	donationMap.Lock()
	defer donationMap.Unlock()
	delete(donationMap.m, donationID)
}

func getDonationList(userID string) ([]Donation, error) {
	var selectQuery string = `SELECT 
	donationId,
	userId,
	bloodType, 
	bloodCenter, 
	amount,
	date,
	status
	FROM donations
	`

	var results *sql.Rows
	var err error
	if userID == "" {
		results, err = database.DbConnection.Query(selectQuery)
	} else {
		selectQuery += `WHERE userId = ?`
		results, err = database.DbConnection.Query(selectQuery, userID)
	}

	if err != nil {
		return nil, err
	}

	donations := make([]Donation, 0)
	for results.Next() {
		var donation Donation
		err = results.Scan(&donation.DonationID,
			&donation.UserID,
			&donation.BloodType,
			&donation.BloodCenter,
			&donation.Amount,
			&donation.Date,
			&donation.Status)

		if err != nil {
			return nil, err
		}

		donations = append(donations, donation)
	}

	defer results.Close()

	return donations, nil
}

func getDonationListByUserID(userID string) ([]Donation, error) {
	results, err := database.DbConnection.Query(`SELECT 
	donationId,
	userId,
	bloodType, 
	bloodCenter, 
	amount,
	date,
	status
	FROM donations
	WHERE userId= ?`, userID)

	if err != nil {
		return nil, err
	}

	donations := make([]Donation, 0)
	for results.Next() {
		var donation Donation
		err = results.Scan(&donation.DonationID,
			&donation.UserID,
			&donation.BloodType,
			&donation.BloodCenter,
			&donation.Amount,
			&donation.Date,
			&donation.Status)

		if err != nil {
			return nil, err
		}

		donations = append(donations, donation)
	}

	defer results.Close()

	return donations, nil
}

func getDonationIds() []int {
	donationMap.RLock()
	donationIds := []int{}
	for key := range donationMap.m {
		donationIds = append(donationIds, key)
	}
	donationMap.RUnlock()
	sort.Ints(donationIds)
	return donationIds
}

func getNextDonationID() int {
	donationIds := getDonationIds()
	return donationIds[len(donationIds)-1] + 1
}

func addOrUpdateDonation(donation Donation) (int, error) {
	// if the donation id is set, update, otherwise add
	addOrUpdateID := -1
	if donation.DonationID > 0 {
		oldDonation := getDonation(donation.DonationID)
		// if it exists, replace it, otherwise return error
		if oldDonation == nil {
			return 0, fmt.Errorf("donation id [%d] doesn't exist", donation.DonationID)
		}
		addOrUpdateID = donation.DonationID
	} else {
		addOrUpdateID = getNextDonationID()
		donation.DonationID = addOrUpdateID
	}
	donationMap.Lock()
	donationMap.m[addOrUpdateID] = donation
	donationMap.Unlock()
	return addOrUpdateID, nil
}
