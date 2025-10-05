package main

import "fmt"

func main() {
	var a, t, L float64
	fmt.Print("Masukkan nilai alas :")
	fmt.Scanln(&a)
	fmt.Print("Masukkan nilai tinggi :")
	fmt.Scanln(&t)

	L = 0.5 * a * t
	fmt.Println("Hasil perhitungan luas segitiga adalah", L+0.5)
}
