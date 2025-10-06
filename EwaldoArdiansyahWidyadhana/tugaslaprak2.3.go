package main

import "fmt"

func main() {
	var tahun int
	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	// Tahun kabisat jika:
	// - Habis dibagi 400, atau
	// - Habis dibagi 4 tapi tidak habis dibagi 100
	kabisat := (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)

	fmt.Println("Kabisat:", kabisat)
}
