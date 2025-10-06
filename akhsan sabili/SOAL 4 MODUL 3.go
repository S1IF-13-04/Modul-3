package main

import "fmt"

func main(){
	var C float64
	fmt.Print("Masukkan suhu dalam C :")
	fmt.Scanln(&C)
	F := (C * 9/5) + 32 
	R := C * 4 / 5
	K := (F + 459.67) * 5/9
	Kint := int(K)
	fmt.Println("Derajat Reamur : ",R)
	fmt.Println("Derajat Fahrenheit : ",F)
	fmt.Println("Derajat Kelvin : ",Kint)
}