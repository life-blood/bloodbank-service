package donation

import (
	"bloodbankservice/database"
	"database/sql"
	"log"
)

func getDonation(donationID int) (*Donation, error) {
	var selectQuery string = `SELECT 
	donationId,
	userId,
	bloodType, 
	bloodCenter, 
	amount,
	date,
	status
	FROM donations
	WHERE donationId = ?
	`
	row := database.DbConnection.QueryRow(selectQuery, donationID)
	donation := &Donation{}

	err := row.Scan(
		&donation.DonationID,
		&donation.UserID,
		&donation.BloodType,
		&donation.BloodCenter,
		&donation.Amount,
		&donation.Date,
		&donation.Status,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return donation, nil
}

func removeDonation(donationID int) error {
	_, err := database.DbConnection.Query(`DELETE FROM donations where WHERE donationId = ?`, donationID)
	if err != nil {
		return err
	}

	return nil
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

func updateDonation(donation Donation) error {
	_, err := database.DbConnection.Exec(`UPDATE donations SET
	userId=?,
	bloodType=?, 
	bloodCenter=?, 
	amount=?,
	date=?,
	status=?
	WHERE donationId = ?`,
		donation.UserID,
		donation.BloodType,
		donation.BloodCenter,
		donation.Amount,
		donation.Date,
		donation.Status,
		donation.DonationID,
	)

	if err != nil {
		return err
	}

	return nil
}

func insertDonation(donation Donation) (int, error) {
	result, err := database.DbConnection.Exec(`INSERT INTO donations
	(userId, 
		bloodType, 
		bloodCenter, 
		amount, 
		date, 
		status) VALUES (?, ?, ?, ?, ?, ?)`,
		donation.UserID,
		donation.BloodType,
		donation.BloodCenter,
		donation.Amount,
		donation.Date,
		donation.Status)

	if err != nil {
		log.Println(err)
		return 0, nil
	}

	log.Println("Inserting donation")

	log.Println(donation)

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}
