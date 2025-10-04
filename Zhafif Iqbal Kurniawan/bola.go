package main

import "fmt"

func main() {
	const pi = 3.1415926535
	var jari int
	var volume, luas float64

	fmt.Scan(&jari)

	jari2 := float64(jari)

	volume = (4.0 / 3.0) * pi * (jari2 * jari2 * jari2)
	luas = 4 * pi * (jari2 * jari2)

	fmt.Print("Bola dengan jejari ", jari, " memiliki volume ", volume, " dan luas kulit ", luas)
}
