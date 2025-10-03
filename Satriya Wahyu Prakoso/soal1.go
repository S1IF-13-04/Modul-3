package main
import "fmt"
func main() {
	var x float64
	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&x)
	x = 2/(x+5)+5
	fmt.Println("Hasil :",x)
}