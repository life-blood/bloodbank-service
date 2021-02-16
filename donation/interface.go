package donation

//Reader interface
type Reader interface {
	Get(id int) (*Donation, error)
	GetDonationList(userID string) ([]Donation, error)
}

//Writer donation writer
type Writer interface {
	Create(d *Donation) (int, error)
	Update(d *Donation) error
	Delete(id int) error
}

//Repository interface
type Repository interface {
	Reader
	Writer
}
