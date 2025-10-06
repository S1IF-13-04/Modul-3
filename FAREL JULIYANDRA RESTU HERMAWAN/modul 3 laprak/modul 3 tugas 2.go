package main

import ("fmt"
		"math")

func main (){
	var j float64
	var v,l float64
	fmt.Print("Masukan Jari-jari :")
	fmt.Scan(&j)
	v = 4.0 / 3.0 * math.Pi * j * j * j
	l = 4.0 * math.Pi * j * j
	fmt.Printf("Bola dengan jari-jari %.0f memiliki volume %.4f dan luas kulit %.4f\n", j, v, l)
}