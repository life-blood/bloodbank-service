package donation

//Reader interface
type Reader interface {
	//Retrieve donation by ID
	Get(id int) (*Donation, error)

	//Retrieve donations filtered by userID if present
	GetDonationList(userID string) ([]Donation, error)
}

//Writer donation writer
type Writer interface {
	//Create new Donation entity
	Create(d *Donation) (int, error)

	//Update existing Donation entity in the system
	Update(d *Donation) error

	//Delete Donation entity if present in the system
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}
