package main

import "fmt"

func main() {
	var x, fahrenheit, reamur, kelvin float64
	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scanln(&x)

	fahrenheit = (x * (9.0/5.0)) + 32
	reamur = x * (4.0/5.0)
	kelvin = x + 273.15

	fmt.Printf("Suhu dalam fahrenheit: %g\n", fahrenheit)
	fmt.Printf("Suhu dalam reamur: %g\n", reamur)
	fmt.Printf("Suhu dalam kelvin: %g", kelvin)
	
}