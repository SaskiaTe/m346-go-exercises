package main

import(
	"math"
	"fmt"
)

func computeQuadraticFormula(a, b, c float64) []float64 {
	D := (b*b -4*a*c)

	if D > 0 {
		root1 := (-b + math.Sqrt(D)) / (2 * a)
		root2 := (-b - math.Sqrt(D)) / (2 * a)
		return []float64{root1, root2};
	}else if D == 0{
		root1 := (-b + math.Sqrt(D)) / (2 * a)
		return []float64{root1};
	}else {
		return []float64{};
	}
}

func main() {
	fmt.Println(computeQuadraticFormula(3, 4, 1))
	fmt.Println(computeQuadraticFormula(2, 4, 2))
	fmt.Println(computeQuadraticFormula(3, 4, 2))
}