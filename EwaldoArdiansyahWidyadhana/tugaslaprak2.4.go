package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&c)

	// Rumus konversi
	f := (c * 9 / 5) + 32
	r := c * 4 / 5
	k := c + 273

	fmt.Printf("Derajat Reamur: %.0f\n", r)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", f)
	fmt.Printf("Derajat Kelvin: %.0f\n", k)
}
