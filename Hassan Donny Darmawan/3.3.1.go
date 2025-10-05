package main
import "fmt"
func main () {
	var s, v int
	fmt.Print("masukan sisi kubus: ")
	fmt.Scanln(&s)
	v =  s * s * s
	fmt.Print ("jadi volume kubus adalah: ", float64(v)+0.5)
}