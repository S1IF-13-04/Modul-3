package main
import ("fmt"
		"math")
func main(){
	var r int
	fmt.Print("jari-jari :")
	fmt.Scanln(&r)
	var V float64 = 4.0 / 3.0 * math.Pi * float64(r) * float64(r) * float64(r)
	var L float64 = 4 * math.Pi * float64(r) * float64(r) 
	fmt.Printf("bola dengan jejari : %d memiliki volume : %.4f dan luas kulit : %.4f \n ", r, V, L)
}