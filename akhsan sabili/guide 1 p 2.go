package main

import "fmt"

func main() {
	var s int
	fmt.Print("masukkan sisi :")
	fmt.Scanln(&s)
	v := s * s * s
	volumeFloat := float64(v)
	fmt.Print(volumeFloat + 0.5)
}
