package main

import "fmt"

func main() {
	type student struct{
		Firstname string
		LastName string
	}

	s1 := student{Firstname: "Saskia", LastName: "Tellenbach"}
	s2 := student{Firstname: "Lea", LastName: "Schürch"}
	s3 := student{Firstname: "Tobias", LastName: "Clausen"}
	s4 := student{Firstname: "Lenny", LastName: "Amstutz"}
	s5 := student{Firstname: "Nicola", LastName: "Ming"}
	s6 := student{Firstname: "Colin", LastName: "Wütschert"}
	s7 := student{Firstname: "Lia", LastName: "Ditli"}

	type class []student
	
	modules := map[string][]class{
		"M320": {
			{s1, s2, s3, s4},
			{s5, s6, s7},
		},
		"M346": {
			{s1, s2, s3, s4},
			{s5, s6, s7},
		},
		"M254": {
			{s5, s6, s7},
		},
	}

	fmt.Println(modules)

}
