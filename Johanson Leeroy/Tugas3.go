package main

import "fmt"

func main() {
	var tahun, kabisat int
	var cek bool
	fmt.Print("Masukan tahun: ")
	fmt.Scan(&tahun)
	kabisat = tahun % 4
	cek = kabisat == 0 == true
	fmt.Print(cek)

}
