package main
import "fmt"

func main() {
    var panjang, tinggi int
    fmt.Print("Masukkan panjang alas= ")
    fmt.Scanln(&panjang)
	fmt.Print("Masukkan tinggi segitiga= ")
    fmt.Scanln(&tinggi)
    luas := panjang * tinggi / 2
    luasFloat := float64(luas) + 0.5
    fmt.Printf("Luas Segitiga = %.1f\n", luasFloat) 
}