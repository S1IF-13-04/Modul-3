package main

import "fmt"

func main() {
	var r float64
	var pi float64
	var volume float64
	var luas float64

	fmt.Print("Jejari = ")
	fmt.Scan(&r)

	pi = 3.1415926535
	volume = (4.0 / 3.0) * pi * (r * r * r)
	luas = 4 * pi * (r * r)

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
