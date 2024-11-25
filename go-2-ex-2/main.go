package main

import "fmt"

func main() {
	var fibs = []int{1, 1, 0, 0, 0}

		fibs[2] = fibs[1] + fibs[0]
		fibs[3] = fibs[2] + fibs[1]
		fibs[4] = fibs[3] + fibs[2]
	
		fibs = append(fibs, fibs[4]+fibs[3]) 
		fibs = append(fibs, fibs[5]+fibs[4]) 
		fibs = append(fibs, fibs[6]+fibs[5])
		fibs = append(fibs, fibs[7]+fibs[6]) 

	fmt.Println(fibs)
}
