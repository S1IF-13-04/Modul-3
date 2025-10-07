package main

import "fmt"

func main() {
	var year int
	var hasil, satu, dua bool
	fmt.Print("Masukan tahun : ")
	fmt.Scan(&year)
	satu = year % 400 == 0
	dua = year % 4 == 0 && year % 100 != 0
	hasil = satu || dua 
	fmt.Printf("Tahun %d merupakan tahun kabisat: %t", year, hasil)
}
