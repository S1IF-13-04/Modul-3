package main
import "fmt"
func main() {
 var sisi int
 fmt.Scan(&sisi)
 volume := float64(sisi) * float64(sisi) * float64(sisi)
 fmt.Println(volume + 0.5)
}
