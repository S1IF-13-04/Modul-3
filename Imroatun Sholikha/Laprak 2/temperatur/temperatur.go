package main

import (
	"fmt"
)

func main() {
	var c, f, r, k float64

	fmt.Print("Celsius: ")
	fmt.Scanln(&c)

	f = (c * 9 / 5) + 32

	r = c * 4 / 5

	k = (f + 459.67) * 5 / 9

	fmt.Println("fahrenheit :", f)
	fmt.Println("reamur :", r)
	fmt.Printf("kelvin :%.0f\n", k)
}
