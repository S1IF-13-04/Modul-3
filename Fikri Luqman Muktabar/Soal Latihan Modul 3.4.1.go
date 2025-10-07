package main

import "fmt"

func main() {
	var celsius float64
	fmt.Print("Temperatur Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Println("Derajat Fahrenheit:", fahrenheit)
}