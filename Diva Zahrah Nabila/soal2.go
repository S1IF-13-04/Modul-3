package main 
import "fmt"
func main(){
	var r, vb, lb float64
	const pi = 3.1415926535

	fmt.Print("Masukan jari-jari :")
	fmt.Scan(&r)
	vb = 4 * pi * r * r * r / 3
	lb = 4 * pi * r * r 
	fmt.Printf("Volume bola: %.4f\n", vb)
	fmt.Printf("Luas bola: %.4f\n", lb)
}