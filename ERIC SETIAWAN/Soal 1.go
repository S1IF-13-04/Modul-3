package main

import "fmt"

func main() {
	var x float64
	fmt.Println("Input: ")
	fmt.Scanln(&x)
	x = 2/(x+5) + 5
	fmt.Println(x)
}
