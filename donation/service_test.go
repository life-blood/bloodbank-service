package donation

import (
	"bloodbankservice/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newFixtureDonation() *Donation {
	return &Donation{
		DonationID:  1,
		UserID:      "user-123",
		BloodType:   "A",
		BloodCenter: "SofiaMed",
		Amount:      "400 mil",
		Date:        "15.2.2021",
		Status:      "In progress",
	}
}

func Test_Create(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)

	donation := newFixtureDonation()

	_, err := service.CreateDonation(donation)
	assert.Nil(t, err)
}

func Test_GetDonation(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	donation := newFixtureDonation()
	donationID, _ := service.CreateDonation(donation)

	retrievedDonation, err := service.GetDonation(donationID)

	assert.Nil(t, err)
	assert.Equal(t, donation, retrievedDonation)
}

func Test_UpdateDonation(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	donation := newFixtureDonation()
	_, _ = service.CreateDonation(donation)

	t.Run("when valid donation", func(t *testing.T) {
		donation.Status = "Completed"
		err := service.UpdateDonation(donation)
		assert.Nil(t, err)
	})

	t.Run("when invalid", func(t *testing.T) {
		donation.Status = ""
		err := service.UpdateDonation(donation)

		assert.Equal(t, err, entity.ErrInvalidEntity)
	})
}

func Test_DeleteDonation(t *testing.T) {
	repo := newInmem()
	service := NewService(repo)
	donation := newFixtureDonation()
	donationID, _ := service.CreateDonation(donation)

	t.Run("when valid donation", func(t *testing.T) {
		err := service.DeleteDonation(donationID)
		assert.Nil(t, err)
	})

	t.Run("when invalid", func(t *testing.T) {
		err := service.DeleteDonation(-1)

		assert.Equal(t, err, entity.ErrNotFound)
	})
}
