package main
import "fmt"
func main() {
	var r float64
	fmt.Print("Masukkan jari-jari lingkaran : ")
	fmt.Scan(&r)
	pi := 3.1415926535
	Luas := 4.0 * pi * r * r
	Volume := (4.0/3.0)* pi * r * r * r
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f", r, Volume, Luas)
}