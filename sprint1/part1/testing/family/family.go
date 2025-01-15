package main

import (
	"errors"
	"fmt"
)

type Relationship string

const (
	Father      = Relationship("father")
	Mother      = Relationship("mother")
	Child       = Relationship("child")
	GrandMother = Relationship("grandMother")
	GrandFather = Relationship("grandFather")
)

var (
	ErrRelationAlreadyExists = errors.New("relationship already exists")
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Family struct {
	Members map[Relationship]Person
}

func (f *Family) AddNew(r Relationship, p Person) error {
	if f.Members == nil {
		f.Members = map[Relationship]Person{}
	}
	if _, ok := f.Members[r]; ok {
		return ErrRelationAlreadyExists
	}
	f.Members[r] = p
	return nil
}

func main() {
	fatherOfCode := Person{"Sanzhar", "Myrzash", 19}
	// fatherOfArch := Person{"Robert", "Martin", 70}
	f := Family{}
	err := f.AddNew(Father, fatherOfCode)
	fmt.Println(f, err)
	err = f.AddNew(Father, Person{"Robert", "Martin", 70})
	fmt.Println(f, err)
}
