package main

import "fmt"

func main() {
	var rupiah, dolar int
	fmt.Scan(&rupiah)
	dolar = rupiah / 15000
	fmt.Print(rupiah, " rupiah= ", dolar, " dolar")
}
