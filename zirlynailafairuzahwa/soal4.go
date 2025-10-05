package main
import "fmt"

func main() {
	var c, f, r, k int
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&c)
	f = (c * 9 / 5) + 32
	r = c * 4 / 5
	kelvin := (float64(f) + 459.67) * 5 / 9
	k = int(kelvin)
	fmt.Println("Derajat Reamur: ", r)
	fmt.Println("Derajat Fahrenheit: ", f)
	fmt.Println("Derajat Kelvin: ", k)
}