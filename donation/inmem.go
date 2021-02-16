package donation

import "bloodbankservice/entity"

//inmem in memory repo
type inmem struct {
	m map[int]*Donation
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[int]*Donation{}
	return &inmem{
		m: m,
	}
}

//Get donation by id
func (r *inmem) Get(id int) (*Donation, error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

//Create a donation
func (r *inmem) Create(d *Donation) (int, error) {
	r.m[d.DonationID] = d
	return d.DonationID, nil
}

//Delete a donation
func (r *inmem) Delete(id int) error {
	if r.m[id] == nil {
		return entity.ErrNotFound
	}
	r.m[id] = nil
	return nil
}

//Update a donation
func (r *inmem) Update(d *Donation) error {
	if r.m[d.DonationID] == nil {
		return entity.ErrNotFound
	}

	r.m[d.DonationID] = d
	return nil
}

//GetDonationList by userID
func (r *inmem) GetDonationList(userID string) ([]Donation, error) {
	var donations []Donation
	for _, element := range r.m {
		if element.UserID == userID {
			donations = append(donations, *element)
		}
	}
	return donations, nil
}
