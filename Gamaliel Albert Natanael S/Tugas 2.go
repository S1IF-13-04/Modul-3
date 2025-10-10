package main
import (
	"fmt"
	"math"
)
func main() {
	var jejari int
	fmt.Print("Jejari = ")
	fmt.Scan(&jejari)
	const pi = 3.1415926535
	var volume, luasKulit float64
	r := float64(jejari)
	volume = (4.0 / 3.0) * pi * math.Pow(r, 3)
	luasKulit = 4 * pi * math.Pow(r, 2)
	fmt.Printf("Bola dengan jejari %d memiliki volume %.4f dan luas kulit %.4f\n", jejari, volume, luasKulit)
}