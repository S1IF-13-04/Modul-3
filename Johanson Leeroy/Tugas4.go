package main

import "fmt"

func main() {
	var celcius, reamur, fahrenheit, kelvin int
	fmt.Print("Masukan suhu dalam bentuk Celcius: ")
	fmt.Scan(&celcius)
	reamur = celcius * 4 / 5
	fahrenheit = celcius*9/5 + 32
	kelvin = celcius + 273
	fmt.Println("Reamur= ", reamur)
	fmt.Println("Fahrenheit= ", fahrenheit)
	fmt.Println("Kelvin= ", kelvin)
}
