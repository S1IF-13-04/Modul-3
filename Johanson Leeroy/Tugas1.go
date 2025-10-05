package main

import "fmt"

func main() {
	var fungsix, x float64
	fmt.Print("Masukan nilai dari x: ")
	fmt.Scan(&x)
	fungsix = 2/(x+5) + 5
	fmt.Print("f(x)= ", fungsix)
}
