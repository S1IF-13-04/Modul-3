package main

import "fmt"

func main() {
	var r float64
	fmt.Print("Masukkan jari-jari bola: ")
	fmt.Scan(&r)

	pi := 3.1415926535
	volume := (4.0 / 3.0) * pi * r * r * r
	luas := 4 * pi * r * r

	fmt.Println("Volume =", volume)
	fmt.Println("Luas Permukaan =", luas)
}
