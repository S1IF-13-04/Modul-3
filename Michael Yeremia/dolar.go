package main

import "fmt"

func main() {
	var rupiah, dolar int
	fmt.Print("Masukan jumlah uang yang akan di tukar : ")
	fmt.Scan(&rupiah)
	dolar = rupiah / 15000
	fmt.Println("Jumlah dolar yg di dapat : ",dolar)
}
