package donation

// Donation
type Donation struct {
	DonationID  int    `json:"donationID"`
	UserID      string `json:"userID"`
	BloodCenter string `json:"bloodcenter"`
	Amount      string `json:amount`
	Date        string `json:"date"`
	Status      string `json:"status"`
}
