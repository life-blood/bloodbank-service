package donation_test

import (
	"bloodbankservice/donation"
	"bloodbankservice/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDonation(t *testing.T) {
	d, err := donation.NewDonation(1, "a123z", "A", "SofiaMed", "300 mil", "16.2.2021", "Completed")
	assert.Nil(t, err)
	assert.Equal(t, d.BloodCenter, "SofiaMed")
}

func TestDonationValidate(t *testing.T) {
	type test struct {
		id          int
		userId      string
		bloodType   string
		bloodCenter string
		amount      string
		date        string
		status      string
		want        error
	}

	tests := []test{
		{
			id:          1,
			userId:      "user-123",
			bloodType:   "A",
			bloodCenter: "SofiaMed",
			amount:      "300 mil",
			date:        "16.2.2021",
			status:      "In progress",
			want:        nil,
		},
		{
			id:          1,
			userId:      "",
			bloodType:   "A",
			bloodCenter: "SofiaMed",
			amount:      "300 mil",
			date:        "16.2.2021",
			status:      "In progress",
			want:        entity.ErrInvalidEntity,
		},
		{
			id:          1,
			userId:      "user-123",
			bloodType:   "",
			bloodCenter: "SofiaMed",
			amount:      "300 mil",
			date:        "16.2.2021",
			status:      "In progress",
			want:        entity.ErrInvalidEntity,
		},
		{
			id:          1,
			userId:      "",
			bloodType:   "A",
			bloodCenter: "SofiaMed",
			amount:      "300 mil",
			date:        "16.2.2021",
			status:      "In progress",
			want:        entity.ErrInvalidEntity,
		},
		{
			id:          1,
			userId:      "user-123",
			bloodType:   "A",
			bloodCenter: "SofiaMed",
			amount:      "",
			date:        "16.2.2021",
			status:      "In progress",
			want:        entity.ErrInvalidEntity,
		},
		{
			id:          1,
			userId:      "user-123",
			bloodType:   "A",
			bloodCenter: "SofiaMed",
			amount:      "300 mil",
			date:        "16.2.2021",
			status:      "",
			want:        entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := donation.NewDonation(
			tc.id, tc.userId, tc.bloodType, tc.bloodCenter, tc.amount, tc.date, tc.status,
		)
		assert.Equal(t, err, tc.want)
	}
}
