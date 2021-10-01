package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	BirthDate time.Time
}

func (ch *Person) ChangeName(string, string) {
	fmt.Printf("%s %s", ch.FirstName, ch.LastName)

}

func (bd *Person) Birthday(time.Time) {
	start := time.Date(1994, 3, 2, 0, 0, 0, 0, time.UTC)
	oneYearLater := start.AddDate(1, 0, 0)
	fmt.Printf("oneYearLater: start.AddDate(1, 0, 0) = %v\n", oneYearLater)

}

func changeName(Person, string, string) {
	oakley := Person{
		FirstName: "Oakley",
		LastName:  "Niters",
	}
	fmt.Print(oakley)
}

func changeNames(Person, string, string) Person {
	w := Person{
		FirstName: "Oakley",
		LastName:  "Nieters",
	}
	return w
}

func NewPerson(string, string, int, time.Time) Person {
	p := Person{
		FirstName: "Oakley",
		LastName:  "Nieters",
		Age:       27,
		BirthDate: time.Date(1994, 03, 02, 0, 0, 0, 0, time.UTC),
	}
	return p
}

func main() {
	oakley := Person{
		FirstName: "oakley",
		LastName:  "nieters",
		Age:       27,
		BirthDate: time.Date(1994, 03, 02, 0, 0, 0, 0, time.UTC),
	}
	fmt.Print(oakley)
	fmt.Print(changeNames)
	
}

