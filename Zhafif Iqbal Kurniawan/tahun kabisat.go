package main

import "fmt"

func main() {
	var tahun int
	var kabisat bool

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)

	kabisat = (tahun%400 == 0) || (tahun%4 == 0)

	fmt.Println("Kabisat:", kabisat)
}
