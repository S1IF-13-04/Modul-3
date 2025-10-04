package main

import "fmt"

func main() {
	var c, f, k, r float64
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&c)

	f = (9.0/5.0)*c + 32
	k = c + 273.15
	r = (4.0 / 5.0) * c

	fmt.Printf("Derajat reamur: %.0f\n", r)
	fmt.Printf("Derajat Farenheit: %.0f\n", f)
	fmt.Printf("Derajat Kelvin: %.0f\n", k)
}
