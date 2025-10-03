package main
import "fmt"
func main() {
	var c float64
	fmt.Print("Masukkan suhu dalam Celcius : ")
	fmt.Scan(&c)
	f := c * 9 / 5 + 32
	r := c * (4.0 / 5.0)
	k := (f + 459.67) * (5.0 / 9.0)
	fmt.Println("Temperatur Celcius :", c)
	fmt.Printf("Derajat Reamur : %0.f\n", r)
	fmt.Printf("Derajat Fanrenheit : %0.f\n", f)
	fmt.Printf("Derajat Kelvin : %0.f", k)
}