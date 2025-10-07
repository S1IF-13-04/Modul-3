package main
import "fmt"

func main(){
	var year int
	var hasil, satu, dua bool
	fmt.Print("Tahun: ")
	fmt.Scan(&year)
	satu = year % 400 == 0
	dua = year % 4 == 0 && year % 100 != 0
	hasil = satu || dua 
	fmt.Printf("Kabisat: %t",hasil)
}