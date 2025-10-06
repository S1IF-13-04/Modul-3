package main

import "fmt"

func main (){
	var Tahun int
	fmt.Print("Masukkan Tahun :")
	fmt.Scan(&Tahun)
	Kabisat := (Tahun%400 == 0) || (Tahun%4 == 0 && Tahun%100 != 0)
	fmt.Println("Kabisat :", Kabisat)
}