package main

import "fmt"

func main() {
	var x, nilai float64

	fmt.Print("Masukan X:")
	fmt.Scan(&x)

	nilai = 2/(x+5) + 5

	fmt.Println(nilai)

}
