package main

import "fmt"

func main() {
	const pi = 3.1415926535
	var r float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	// Rumus:
	// Volume bola = (4/3) * π * r^3
	// Luas bola   = 4 * π * r^2
	volume := (4.0 / 3.0) * pi * r * r * r
	luas := 4 * pi * r * r

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
