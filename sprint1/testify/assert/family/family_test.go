package family

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFamily_AddNew(t *testing.T) {
	type newPerson struct {
		r Relationship
		p Person
	}
	tests := []struct {
		name      string
		members   map[Relationship]Person
		newMember newPerson
		wantErr   bool
	}{
		{
			name: "add father",
			members: map[Relationship]Person{
				Mother: {
					FirstName: "Marie",
					LastName:  "Curie",
				},
			},
			newMember: newPerson{
				r: Father,
				p: Person{
					FirstName: "Albert",
					LastName:  "Einstein",
					Age:       70,
				},
			},
			wantErr: false,
		},
		{
			name: "add another father",
			members: map[Relationship]Person{
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
			wantErr: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			family := &Family{
				Members: test.members,
			}
			err := family.AddNew(test.newMember.r, test.newMember.p)
			if !test.wantErr {
				require.NoError(t, err)
				assert.Contains(t, family.Members, test.newMember.r)
				return
			}
			assert.Error(t, err)
		})
	}
}
