package main

import "fmt"

func main() {
	var c, r, f, k float64
	fmt.Print("Input Suhu Dalam Celcius:  ")
	fmt.Scan(&c)

	r = (c * 4) / 5
	f = (c * 9 / 5) + 32
	k = (f + 459.67) * 5 / 9

	fmt.Println("Temperatur Celsius	: ", c)
	fmt.Println("Derajat Reamur		: ", r)
	fmt.Println("Derajat Fahrenheit	: ", f)
	fmt.Printf("Derajat Kelvin		:  %.f\n", k)
}
