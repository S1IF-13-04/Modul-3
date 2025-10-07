package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukan tahun: ")
	fmt.Scan(&tahun)
	fmt.Println("Tahun", tahun)
	kabisat := (tahun%4 == 0 && tahun%100 != 0) || (tahun%4 == 0)
	fmt.Println("Kabisat", kabisat)

}
