package main

import (
	"fmt"
)

func main() {
	var alas, tinggi int

	fmt.Print("Masukkan panjang alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	luas := (alas * tinggi) / 2

	fmt.Println("Luas segitiga adalah:", luas)
}
