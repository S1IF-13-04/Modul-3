package main

import "fmt"

func main() {
	var c, f, k, r float64
	fmt.Print("Temperatur Celcius: ")
	fmt.Scan(&c)

	f = (9.0/5.0)*c + 32
	k = c + 273.15
	r = (4.0 / 5.0) * c

	fmt.Println("Temeperatur Celacius : ", c)
	fmt.Printf("reamur : %.0f\n", r)
	fmt.Printf("Farenheit : %.0f\n", f)
	fmt.Printf("Kelvin : %.0f\n", k)
}
