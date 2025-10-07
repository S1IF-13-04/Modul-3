package main

import "fmt"

func main() {
	var x, volume, luas float64
	const pi float64 = 3.1415926535
	fmt.Print("Masukkan angkanya: ")
	fmt.Scanln(&x)

	volume = 4.0/3.0 * pi * x * x * x
	luas = 4 * pi * x * x

	fmt.Printf("Bola dengan jari jari %g memiliki volume %g dan luas kulit %g",x,volume, luas)
}