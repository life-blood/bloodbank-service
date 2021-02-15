package donation

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
