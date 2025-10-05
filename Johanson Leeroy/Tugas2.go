package main

import "fmt"

func main() {
	var jejari int
	var volumebola, luasbola float64
	var PI float64 = 3.1415926535
	fmt.Print("Masukan jari-jari bola: ")
	fmt.Scan(&jejari)
	var r = float64(jejari)
	volumebola = 4 * PI * r * r * r / 3
	luasbola = 4 * PI * r * r
	fmt.Println("Volume Bola=", volumebola)
	fmt.Println("Luas Bola=", luasbola)
}
