package main

import "fmt"

func main() {
	var celcius, fahrenheit, reamur, kelvin float64
	fmt.Print("Temperatur celcius: ")
	fmt.Scan(&celcius)
	fahrenheit = (celcius * 9.0 / 5.0) + 32.0
	reamur = celcius * 4.0 / 5.0
	kelvin = celcius + 273.15
	fmt.Printf("Derajat fahrenheit: %0.f\n", fahrenheit)
	fmt.Printf("Derajat reamur: %0.f\n", reamur)
	fmt.Printf("Derajat kelvin: %0.f\n", kelvin)
}
