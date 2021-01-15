package donation

// Donation ...
type Donation struct {
	DonationID  int    `json:"donationId"`
	UserID      string `json:"userId"`
	BloodType   string `json:"bloodType"`
	BloodCenter string `json:"bloodCenter"`
	Amount      string `json:"amount"`
	Date        string `json:"date"`
	Status      string `json:"status"`
}
