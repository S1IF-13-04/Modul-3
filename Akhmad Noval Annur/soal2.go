package main
import "fmt"
func main (){
    var r,V,L float64 
    const pi float64 = 3.1415926535
    fmt.Print("Masukkan Jejarinya: ")
    fmt.Scan(&r)
    V = 4.0/3.0 * pi * r * r * r 
    L = 4 * pi * r * r
    fmt.Printf("Bola dengan jejari %g memiliki volume %.4f dan luas kulit %.4f",r,V,L)
}