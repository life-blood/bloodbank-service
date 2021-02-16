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
