package main
import "fmt"

func computeGrade(gotPoints int, maxPoints int) float64 {
	if gotPoints <= 0 {
		return 1 ;
	}
	grade := float64(gotPoints) / float64(maxPoints) * 5 + 1
	return grade;
}

func main() {
	fmt.Println(computeGrade(45, 100)) 
	fmt.Println(computeGrade(80, 100)) 
	fmt.Println(computeGrade(57, 60)) 
}
