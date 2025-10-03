package main

import "fmt"

func main() {
	var Celsius, Fahrenheit, Reamur, Kelvin float64

	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&Celsius)

	Fahrenheit = (Celsius * 9 / 5) + 32
	Reamur = Celsius * 4 / 5
	Kelvin = Celsius + 273.15

	fmt.Printf("Suhu dalam Fahrenheit: %.2f\n", Fahrenheit)
	fmt.Printf("Suhu dalam Reamur: %.2f\n", Reamur)
	fmt.Printf("Suhu dalam Kelvin: %.2f\n", Kelvin)
}
