package main

import (
	"fmt"
)

func main() {
	var alas, tinggi int

	fmt.Print("Masukan panjang alas segitiga: ")
	fmt.Scan(&alas)
	fmt.Print("Masukan tinggi segitiga: ")
	fmt.Scan(&tinggi)

	luas := (alas * tinggi) / 2

	fmt.Println("Luas segitiga adalah:", luas)
}
