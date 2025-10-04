package main

import "fmt"

func main() {
	var s int
	var v int
	
	fmt.Print("Input: ")
	fmt.Scanln(&s)

	v =  s * s * s 
	var f = float64(v) 
	fmt.Print(f + 0.5)
}
