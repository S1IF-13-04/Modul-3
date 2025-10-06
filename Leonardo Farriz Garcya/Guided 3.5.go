package main

import (
	"fmt"
	"math"
)

func main() {
	var r int
	fmt.Scan(&r)
	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(r), 3)
	luas := 4 * math.Pi * math.Pow(float64(r), 2)
	fmt.Printf("hasil volume %.4f dan hasil luas %.4f", volume, luas)
	
}