package main

import "fmt"

func main() {
	var r, volume, luas float64
	const pi float64 = 3.1415926535
	fmt.Scan(&r)

	volume = 4.0 / 3.0 * pi * r * r * r
	luas = 4.0 * pi * r * r
	
	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}