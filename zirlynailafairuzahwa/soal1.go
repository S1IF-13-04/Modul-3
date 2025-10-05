package main
import "fmt"

func main() {
	var x, hasil float64
	fmt.Print("Masukkan nilai x = ")
	fmt.Scan(&x)
	hasil = (2.0 / (x + 5.0)) + 5.0
	fmt.Println("Nilai f(x)-nya adalah ", hasil)
}