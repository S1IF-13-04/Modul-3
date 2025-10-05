package main
import "fmt"

func main() {
    var sisi int
    fmt.Print("Masukkan sisi kubus= ")
    fmt.Scanln(&sisi)
	volume := sisi * sisi * sisi
    hasil := float64(volume) + 0.5
    fmt.Printf("Volume kubus = %.1f\n", hasil) 
}