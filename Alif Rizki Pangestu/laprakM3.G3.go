package main

import (
	"fmt"
)

func main() {
	var rupiah int
	const kurs = 15000

	fmt.Print("Masukkan jumlah uang (dalam IDR): ")
	fmt.Scan(&rupiah)

	usd := rupiah / kurs

	fmt.Println("Hasil konversi ke USD:", usd)
}
