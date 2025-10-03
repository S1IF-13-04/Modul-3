package main

import (
	"fmt"
	"math"
)

func main() {
	var jari, luas, volume float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scanln(&jari)
	luas = 4 * math.Pi * math.Pow(jari, 2)
	volume = (4.0 / 3.0) * math.Pi * math.Pow(jari, 3)
	fmt.Printf("Bola dengan jejari %.f memiliki volume %.4f dan luas kulit %.4f", jari, volume, luas)
}
