package main
import "fmt"

func main () {
	var x int
	var kabisat bool
	fmt.Print("masukan tahun: ")
	fmt.Scanln(&x)
	kabisat = x % 400 == 0 || (x % 4 == 0 && x % 100 != 0 )
	fmt.Print("kabisat: ", kabisat)
}