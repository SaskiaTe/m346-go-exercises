package main

import "fmt"

type FullName struct {
	Firstname string
	LastName  string
}

type BirthDate struct {
	DateOfBirth string
}

type Profile struct {
	FullName
	BirthDate
	NumberOfSiblings byte
	ZodiacSign       rune
}

func main() {
	var me = Profile{
		NumberOfSiblings: 1,
		ZodiacSign:       '\u2652',
		BirthDate:        BirthDate{DateOfBirth: "16.02.2007"},
		FullName: FullName{Firstname: "Saskia",
			LastName: "Tellenbach"},
	}

	fmt.Println(me)

	fmt.Println("Siblings Before:", me.NumberOfSiblings)
	fmt.Println("Siblings After:", me.NumberOfSiblings+1)
}
