package main

import "fmt"

func main() {
	var a, t, l float64
	fmt.Println("Input: ")
	fmt.Scanln(&a, &t)
	l = 0.5 * a * t
	fmt.Println(l + 0.5)
}
