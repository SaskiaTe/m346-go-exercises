package main

import "fmt"

func main() {
	// TODO: Declare and initialize the variables being used in the output!
	var firstName string = "Saskia"
	var lastName string = "Tellenbach"
	var dayOfBirth int = 16
	var monthOfBirth int = 2
	var yearOfBirth int = 2007
	var numberOfSiblings int = 1
	var heightInMeters float32 = 1.73
	var zodiacSign rune = '\u2652'
	fmt.Printf("Vor- und Nachname: %s %s\n", firstName, lastName)
	fmt.Printf("Geburtsdatum: %d.%d.%d\n", dayOfBirth, monthOfBirth, yearOfBirth)
	fmt.Printf("Anzahl Geschwister: %d\n", numberOfSiblings)
	fmt.Printf("Grösse (in Metern): %.2f\n", heightInMeters)
	fmt.Printf("Sternzeichen: %c\n", zodiacSign)
}
