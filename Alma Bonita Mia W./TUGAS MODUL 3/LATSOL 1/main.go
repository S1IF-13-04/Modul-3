package main

import "fmt"

func main() {
    var fx, x float64

    fmt.Print("Masukkan nilai f(x): ")
    fmt.Scan(&fx)

    x = (27 - 5*fx) / (fx - 5)

    fmt.Printf("Nilai x adalah: %.0f\n", x)
}
