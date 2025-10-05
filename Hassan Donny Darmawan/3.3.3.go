package main
import "fmt"
func main () {
	var rupiah, USD float64
fmt.Print("masukan nilai rupiah: ")
fmt.Scanln(&rupiah)
USD = rupiah/15000
fmt.Println("hasil dari konversi rupiah adalah", USD, "USD")
}