package main
import "fmt"
func main() {
 var alas, tinggi int
 fmt.Scan(&alas, &tinggi)
 luas := 0.5 * float64(alas) * float64(tinggi)
 fmt.Println(luas + 0.5)
}
