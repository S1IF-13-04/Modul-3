package main

import "fmt"

func main() {
	var s, v float64
	fmt.Println("Input: ")
	fmt.Scanln(&s)
	v = s * s * s
	fmt.Println(v + 0.5)
}
