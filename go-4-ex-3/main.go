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
	a1, b1, c1 := 3.0, 4.0, 1.0
	solutions1 := computeQuadraticFormula(a1, b1, c1)
	fmt.Printf("Lösungen für a=%.1f, b=%.1f, c=%.1f: %v\n", a1, b1, c1, solutions1)

	a2, b2, c2 := 2.0, 4.0, 2.0
	solutions2 := computeQuadraticFormula(a2, b2, c2)
	fmt.Printf("Lösungen für a=%.1f, b=%.1f, c=%.1f: %v\n", a2, b2, c2, solutions2)

	a3, b3, c3 := 3.0, 4.0, 2.0
	solutions3 := computeQuadraticFormula(a3, b3, c3)
	fmt.Printf("Lösungen für a=%.1f, b=%.1f, c=%.1f: %v\n", a3, b3, c3, solutions3)
}
