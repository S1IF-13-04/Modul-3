package main

import "fmt"

func main() {
	var r, volume, luas float32
	const phi = 3.1415926535
	fmt.Scan(&r)

	volume = (4.0 / 3.0 * phi) * r * r * r
	luas = 4.0 * phi * r * r

	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}
