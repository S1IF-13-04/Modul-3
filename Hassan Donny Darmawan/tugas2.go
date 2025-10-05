package main
import (
	"fmt"
"math")

func main () {
	var r int
	fmt.Print("masukan nilai jari-jari bola: ")
	fmt.Scanln(&r)
	volume := (4.0/3.0 * math.Pi * math.Pow(float64(r),3))
	luas := (4 * math.Pi * math.Pow(float64(r),2))
	fmt.Printf("volume bola adalah: %.4f\n", volume)
	fmt.Printf("luas bola adalah: %.4f", luas)
}