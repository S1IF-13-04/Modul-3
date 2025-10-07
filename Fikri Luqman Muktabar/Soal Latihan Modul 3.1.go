package main

import "fmt"

func main() {
	var x float64
	fmt.Scan(&x)
	
	x = (2 / (x + 5)) + 5

	fmt.Println(x)
}