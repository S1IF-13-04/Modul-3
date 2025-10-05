package main
import "fmt"

func main() {
    var idr int
    fmt.Print("Masukkan mata uang IDR = ")
    fmt.Scanln(&idr)
	usd := idr / 15000
    fmt.Println("konversi dari IDR ke USD = ", usd) 
}