package main

import "fmt"

func convertCelsiusToFahrenheit(c float64)float64{
	f := c * 9/5 + 32 
	return f
}

func convertFahrenheitToCelsius(f float64)float64{
	c := (f - 32) * 5 / 9 
	return c
}

func main() {
	fmt.Println(convertCelsiusToFahrenheit(32))
	fmt.Println(convertCelsiusToFahrenheit(40))
	fmt.Println(convertCelsiusToFahrenheit(50))
	fmt.Println(convertFahrenheitToCelsius(80))
	fmt.Println(convertFahrenheitToCelsius(60))
	fmt.Println(convertFahrenheitToCelsius(50))
}
