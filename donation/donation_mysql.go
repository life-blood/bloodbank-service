package donation

import (
	"database/sql"
	"log"
)

//RepoDatabaseLayer repo
type RepoDatabaseLayer struct {
	db *sql.DB
}

//NewDonationMySQL create new repository
func NewDonationMySQL(db *sql.DB) *RepoDatabaseLayer {
	return &RepoDatabaseLayer{
		db: db,
	}
}

//Get donation by ID
func (r *RepoDatabaseLayer) Get(id int) (*Donation, error) {
	stmt, err := r.db.Prepare(`SELECT 
	donationId,
	userId,
	bloodType, 
	bloodCenter, 
	amount,
	date,
	status
	FROM donations
	WHERE donationId = ?`)

	if err != nil {
		return nil, err
	}

	var d Donation
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&d.DonationID, &d.UserID, &d.BloodType, &d.BloodCenter, &d.Amount, &d.Date, &d.Status)
	}

	return &d, nil
}

// Create a donation
func (r *RepoDatabaseLayer) Create(d *Donation) (int, error) {
	result, err := r.db.Exec(`INSERT INTO donations
	(userId, 
		bloodType, 
		bloodCenter, 
		amount, 
		date, 
		status) VALUES (?, ?, ?, ?, ?, ?)`,
		d.UserID,
		d.BloodType,
		d.BloodCenter,
		d.Amount,
		d.Date,
		d.Status)

	if err != nil {
		log.Println(err)
		return 0, nil
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}

//Update a donation
func (r *RepoDatabaseLayer) Update(d *Donation) error {
	_, err := r.db.Exec(`UPDATE donations SET
	userId=?,
	bloodType=?, 
	bloodCenter=?, 
	amount=?,
	date=?,
	status=?
	WHERE donationId = ?`,
		d.UserID,
		d.BloodType,
		d.BloodCenter,
		d.Amount,
		d.Date,
		d.Status,
		d.DonationID,
	)

	if err != nil {
		return err
	}

	return nil
}

//Delete a donation
func (r *RepoDatabaseLayer) Delete(donationID int) error {
	_, err := r.db.Query(`DELETE FROM donations WHERE donationId = ?`, donationID)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

//GetDonationList retrieve all donation for current userID
func (r *RepoDatabaseLayer) GetDonationList(userID string) ([]Donation, error) {
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
		results, err = r.db.Query(selectQuery)
	} else {
		selectQuery += `WHERE userId = ?`
		results, err = r.db.Query(selectQuery, userID)
	}

	if err != nil {
		return nil, err
	}

	donations := make([]Donation, 0)
	for results.Next() {
		var d Donation
		err = results.Scan(&d.DonationID,
			&d.UserID,
			&d.BloodType,
			&d.BloodCenter,
			&d.Amount,
			&d.Date,
			&d.Status)

		if err != nil {
			return nil, err
		}

		donations = append(donations, d)
	}

	defer results.Close()

	return donations, nil
}
