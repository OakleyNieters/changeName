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

func (bd *Person) Birthday() {
	bd.BirthDate = time.Date(0,0,0,0,0,0,0, time.UTC)
	today := time.Date(2021, 10, 01,0,0,0,0,time.UTC)

	bd.Age = today.Year() - bd.BirthDate.Year()
	bd.Age++
}

func changeName(p *Person, f string, l string) {
	p.ChangeName(f, l)
}

func changeNames(p Person, fn string, ln string) Person {
	p.ChangeName(fn, ln)
	return Person{}
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
	fmt.Print(changeName)

}

