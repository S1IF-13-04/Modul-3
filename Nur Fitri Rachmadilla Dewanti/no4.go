package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Masukkan suhu Celsius: ")
	fmt.Scan(&c)

	f := (c * 9 / 5) + 32
	r := (c * 4 / 5)
	k := c + 273

	fmt.Println("Celsius =", c)
	fmt.Println("Reamur =", r)
	fmt.Println("Fahrenheit =", f)
	fmt.Println("Kelvin =", k)
}
