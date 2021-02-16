package donation

import (
	"bloodbankservice/entity"
)

//Service for donations
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//GetDonation get a donation by donation id
func (s *Service) GetDonation(id int) (*Donation, error) {
	d, err := s.repo.Get(id)
	if d == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return d, nil
}

//GetDonationList get all donations for given user
func (s *Service) GetDonationList(userID string) ([]Donation, error) {
	return s.repo.GetDonationList(userID)
}

// CreateDonation create a donation
func (s *Service) CreateDonation(d *Donation) (int, error) {
	return s.repo.Create(d)
}

//UpdateDonation update a donation identified by id
func (s *Service) UpdateDonation(d *Donation) error {
	err := d.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(d)
}

//DeleteDonation delete a donation identified by id
func (s *Service) DeleteDonation(id int) error {
	_, err := s.repo.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
