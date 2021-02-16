package donation

import "bloodbankservice/entity"

// Donation is а struct used to represent Donation entity in the LifeBlood system
type Donation struct {
	DonationID  int    `json:"donationId"`
	UserID      string `json:"userId"`
	BloodType   string `json:"bloodType"`
	BloodCenter string `json:"bloodCenter"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}

//NewDonation Create a new donation entity with provided data
func NewDonation(
	id int,
	userID string,
	bloodType string,
	bloodCenter string,
	amount string,
	date string,
	status string,
) (*Donation, error) {

	d := &Donation{
		DonationID:  id,
		UserID:      userID,
		BloodType:   bloodType,
		BloodCenter: bloodCenter,
		Amount:      amount,
		Date:        date,
		Status:      status,
	}

	err := d.Validate()
	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	return d, nil
}

//Validate donation entry
func (d *Donation) Validate() error {
	if d.DonationID <= 0 ||
		d.UserID == "" ||
		d.BloodType == "" ||
		d.BloodCenter == "" ||
		d.Amount == "" ||
		d.Status == "" {
		return entity.ErrInvalidEntity
	}
	return nil
}
