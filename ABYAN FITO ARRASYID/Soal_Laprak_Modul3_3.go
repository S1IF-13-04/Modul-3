package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)
	fmt.Println("Tahun:", tahun)

	Kabisat := (tahun%4 == 0 && tahun%100 != 0) || (tahun%400 == 0)
	fmt.Println("Kabisat:", Kabisat)
}
