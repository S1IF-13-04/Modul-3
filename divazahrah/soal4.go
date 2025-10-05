package main
import "fmt"
func main(){
	var c, r, f, k float64

	fmt.Print("Masukan suhu dalam celcius:")
	fmt.Scan(&c)

	r= c * 4/5
	f= (c * 9/5) + 32
	k= (f + 459.67)* 5/9

	fmt.Printf("Derajat reamur: %.0f\n", r)
	fmt.Printf("Derajat fahrenheit: %.0f\n", f)
	fmt.Printf("Derajat kelvin: %.0f\n", k)
} 