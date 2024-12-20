package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUser_FullName(t *testing.T) {
	tests := []struct {
		name   string
		fields struct {
			FirstName string
			LastName  string
		}
		fullName string
	}{
		{
			name: "simple test",
			fields: struct {
				FirstName string
				LastName  string
			}{
				FirstName: "Misha",
				LastName:  "Popov",
			},
			fullName: "Misha Popov",
		},
		{
			name: "long name",
			fields: struct {
				FirstName string
				LastName  string
			}{
				FirstName: "Pablo Diego KHoze Frantsisko de Paula KHuan" +
					" Nepomukeno Krispin Krispiano de la Santisima Trinidad Ruiz",
				LastName: "Picasso",
			},
			fullName: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno" +
				" Krispin Krispiano de la Santisima Trinidad Ruiz Picasso",
		},
		{
			name: "spaces",
			fields: struct {
				FirstName string
				LastName  string
			}{
				FirstName: " ",
				LastName:  " ",
			},
			fullName: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			user := User{test.fields.FirstName, test.fields.LastName}
			assert.Equal(t, test.fullName, user.FullName())
		})
	}
}
