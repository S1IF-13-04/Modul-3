package main

import (
	"fmt"
	"math"
)

func main() {
	var j int

	fmt.Print("Jari - jari = ")
	fmt.Scanln(&j)

	r := float64(j)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)

	luas := 4 * math.Pi * math.Pow(r, 2)

	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", j, volume, luas)
}
