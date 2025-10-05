package main

import "fmt"

func main() {
	var alas, tinggi, luas float64
	fmt.Print("Masukan alas dan tingginya: ")
	fmt.Scan(&alas, &tinggi)
	luas = 0.5 * alas * tinggi
	fmt.Print("Luas= ", luas)
}
