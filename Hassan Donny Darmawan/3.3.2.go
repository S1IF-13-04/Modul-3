package main

import "fmt"

func main() {
	var s, t, v float64
	fmt.Print("masukan nilai sisi dan tinggi: ")
	fmt.Scan(&s, &t)
	v = ((1.0 / 2.0) * s * t)
	fmt.Println("luas dari segitiga adalah: ",v)
}
