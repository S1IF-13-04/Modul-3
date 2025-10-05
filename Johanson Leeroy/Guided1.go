package main

import "fmt"

func main() {
	var sisi, volume float64
	fmt.Print("Masukan panjang sisinya: ")
	fmt.Scan(&sisi)
	volume = sisi * sisi * sisi
	fmt.Print("Volume= ", volume)
}
