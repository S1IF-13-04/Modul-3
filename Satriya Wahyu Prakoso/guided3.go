package main
import "fmt"
func main() {
 var rupiah int
 fmt.Scan(&rupiah)
 dolar := float64(rupiah) / 15000
 fmt.Println(dolar + 0.5)
}