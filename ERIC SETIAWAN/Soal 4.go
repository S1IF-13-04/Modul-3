package main

import "fmt"

func main() {
	var c, f, k, r float64
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&c)
	f = (c * 9 / 5) + 32
	r = (c * 4) / 5
	k = c + 273
	fmt.Println("Derajat Celcius: ", c)
	fmt.Println("Derajat Reamur: ", r)
	fmt.Println("Derajat Fahrenheit: ", f)
	fmt.Println("Derajat Kelvin: ", k)

}
