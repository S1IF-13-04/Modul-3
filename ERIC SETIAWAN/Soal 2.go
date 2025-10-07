package main

import "fmt"

func main() {
	var r float64
	const phi = 3.1415926535
	fmt.Print("Input: ")
	fmt.Scan(&r)
	volume := (4 * phi * r * r * r) / 3
	luas := 4 * phi * r * r
	fmt.Printf("Volume bola: %.4f\n", volume)
	fmt.Printf("Luas: %.4f\n", luas)

}
