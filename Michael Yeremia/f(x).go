package main

import "fmt"

func main() {
	var x, hasil float64
	fmt.Print("Masukkan angkanya: ")
	fmt.Scanln(&x)

	hasil = 2 /(x + 5) + 5

	fmt.Printf("hasilnya: %g", hasil)
}