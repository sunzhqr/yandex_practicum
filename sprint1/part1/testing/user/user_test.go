package main

import (
	"testing"
)

func TestUser_Fullname(t *testing.T) {
	tests := []struct {
		name   string
		fields struct {
			FirstName string
			LastName  string
		}
		want string
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
			want: "Misha Popov",
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
			want: "Pablo Diego KHoze Frantsisko de Paula KHuan Nepomukeno" +
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
			want: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			user := User{FirstName: test.fields.FirstName, LastName: test.fields.LastName}
			if fullName := user.FullName(); fullName != test.want {
				t.Errorf("Want fullname is %s, got %s", test.want, fullName)
			}
		})
	}
}
