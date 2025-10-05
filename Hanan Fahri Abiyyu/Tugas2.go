package main

import (
	"fmt"
)

func main() {
	const pi = 3.1415926535
	var r int
	var V, L float64
	fmt.Print("Masukkan jari-jari :")
	fmt.Scanln(&r)

	r2 := float64(r)

	V = (4.0 / 3.0) * pi * r2 * r2 * r2
	L = 4 * pi * r2 * r2
	fmt.Println("Bola dengan jari-jari ", r, "memiliki volume ", V, "dan luas kulit ", L)
}
