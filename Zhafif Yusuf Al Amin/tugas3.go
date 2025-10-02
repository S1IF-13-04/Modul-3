package main
import "fmt"

func main(){
	var tahun int
	var kabisat bool

	fmt.Print("Masukan Tahun:")
	fmt.Scan(&tahun)

	kabisat = tahun % 4 == 0

	fmt.Print("Kabisat: " ,kabisat)
	
}