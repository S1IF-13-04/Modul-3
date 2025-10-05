package main
import "fmt"

func main () {
	var y float64
	fmt.Print("masukan nilai f(x): ")
	fmt.Scan(&y)
	x := (2/(y+5))+5
	fmt.Printf("hasil dari f(x): %g ", x)
}