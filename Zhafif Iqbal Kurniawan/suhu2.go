package main

import (
	"fmt"
)

func main() {

	var celcius float64

	fmt.Print("Temperatur Celsius :")
	fmt.Scanln(&celcius)

	reamur := (4 / 5) * celcius
	fahrenheit := (celcius * 9 / 5) + 32
	kelvin := celcius + 273.15

	fmt.Println("Derajat Reamur:", reamur)
	fmt.Println("Derajat Fahrenheit:", fahrenheit)
	fmt.Println("Derajat Kelvin:", kelvin)

}
