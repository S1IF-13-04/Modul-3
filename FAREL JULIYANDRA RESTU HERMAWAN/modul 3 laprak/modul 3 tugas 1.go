package main

import "fmt"

func main (){
	var x,fx float64
	fmt.Printf("Masukan f(x) :")
	fmt.Scan(&x)
	fx = 2 / (x + 5) + 5
	fmt.Println(fx)
}