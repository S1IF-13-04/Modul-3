package main
import "fmt"

func main(){
	var r float64
	const pi = 3.1415926535

	fmt.Print("Masukan Jari-Jarinya:")
	fmt.Scan(&r)
	vb := (4* pi * r * r * r)/3
	lb :=  4 * pi * r * r

	fmt.Printf("Volume Bola Adalah: %.4f\n", vb)
	fmt.Printf("Luas Bola Adalah: %.4f\n", lb)

}