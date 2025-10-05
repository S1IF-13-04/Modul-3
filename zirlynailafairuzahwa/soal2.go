package main
import "fmt"

func main() {
	var jariJari int
	fmt.Print("Jari-jari = ")
	fmt.Scan(&jariJari)
	volume := 4.0 / 3.0 * 3.1415926535 * float64(jariJari) * float64(jariJari) * float64(jariJari)
	luas := 4.0 * 3.1415926535 * float64(jariJari) * float64(jariJari)
	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", jariJari, volume, luas)
}