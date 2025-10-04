package main

	import "fmt"

	func main() {
		var a, t int
		
		fmt.Print("Input alas: ")
		fmt.Scanln(&a)
		fmt.Print("Input tinggi: ")
		fmt.Scanln(&t)

		v :=  (a * t )/2
		var f = float64(v)
		fmt.Print(f + 0.5)
	}