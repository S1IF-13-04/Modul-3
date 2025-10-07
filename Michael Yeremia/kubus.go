package main

import "fmt"

func main() {
	var sisi, volume float64
	fmt.Print("Masukan sisinya : ")
	fmt.Scan(&sisi)
	volume = sisi * sisi * sisi
	fmt.Println("Volume nya adalah: ",volume)
}
