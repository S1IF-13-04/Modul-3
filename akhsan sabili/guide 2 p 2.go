package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	L := a * b / 2
	LFloat := float64(L)
	fmt.Print("Luas segitiga :", LFloat+0.5)
}
