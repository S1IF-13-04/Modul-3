package main
import "fmt"
func main(){
	var x, f float64

	fmt.Print("Masukan nilai x:")
	fmt.Scan(&x)
	f= 2/(x +5) + 5
	fmt.Println(f)
}