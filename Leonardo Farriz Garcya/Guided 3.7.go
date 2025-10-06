package main
import "fmt"

func main () {
	var c,f,r,k float64
	fmt.Print("masukan derajat dalam celcius: ")
	fmt.Scan(&c)
	f = (c * 9/5) + 32
	r = c * 4/5
	k = c + 273
	fmt.Println("derajat reamur: ", r)
	fmt.Println("derajat fahrenheit: ", f)
	fmt.Println("derajat kelvin: ", k)
}