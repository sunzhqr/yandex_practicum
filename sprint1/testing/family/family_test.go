package main

import (
	"fmt"
	"testing"
)

func TestFamily_AddNew(t *testing.T) {
	testPeople := []struct {
		name           string
		existedMembers map[Relationship]Person
		newMember      struct {
			r Relationship
			p Person
		}
		wantErr error
	}{
		{
			name: "add father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Marie",
					LastName:  "Curie",
				},
			},
			newMember: struct {
				r Relationship
				p Person
			}{
				r: Father,
				p: Person{
					FirstName: "Albert",
					LastName:  "Einstein",
					Age:       70,
				},
			},
			wantErr: nil,
		},
		{
			name: "add another father",
			existedMembers: map[Relationship]Person{
				Mother: {
					FirstName: "Marie",
					LastName:  "Curie",
				},
				Father: {
					FirstName: "Albert",
					LastName:  "Einstein",
					Age:       70,
				},
			},
			newMember: struct {
				r Relationship
				p Person
			}{
				r: Father,
				p: Person{
					FirstName: "Isaak",
					LastName:  "Newton",
				},
			},
			wantErr: nil,
		},
	}

	for _, testPerson := range testPeople {
		t.Run(testPerson.name, func(t *testing.T) {
			f := Family{
				Members: testPerson.existedMembers,
			}
			if err := f.AddNew(testPerson.newMember.r, testPerson.newMember.p); err != testPerson.wantErr {
				t.Errorf("AddNew() err = %v, wanErr = %v", err, testPerson.wantErr)
				fmt.Printf("Your family consists %v\n", f.Members)
				fmt.Println("And you tried to add", testPerson.newMember.p.FirstName, "with role", testPerson.newMember.r)
			}
		})
	}
}
