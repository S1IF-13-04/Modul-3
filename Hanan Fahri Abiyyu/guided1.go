package main

import "fmt"

func main() {
	var s, V float64
	fmt.Scan(&s)
	V = s * s * s
	fmt.Println(V + 0.5)
}
