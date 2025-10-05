package main

import (
	"fmt"
)

func main() {
	var x, fx float64
	fmt.Print("Masukkan nilai x :")
	fmt.Scanln(&x)

	fx = 2/(x+5) + 5
	fmt.Println("Hasil :", fx)
}
