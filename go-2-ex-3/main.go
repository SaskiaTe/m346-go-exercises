package main

import (
	"fmt"
)

func main() {
	modules := make(map[int]string)

	modules[114] = "Verschlüsslungsverfahren"
	modules[117] = "Netzinfrastruktur"
	modules[346] = "Cloud Lösungen konzipieren"
	

	fmt.Println("Modul 114:", modules[114])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	// delete one
	delete(modules, 117)
	
	// add one
	modules[320] = "OOP Objekt orientiertes Programmieren"

	// replace one
	modules[320] = "Einfaches oop"

	fmt.Println(modules)
}
