package main

import "fmt"

func main() {
	var year int
	var leap bool

	fmt.Scan(&year)

	leap = (year%4 == 0) || (year%400 == 0 && year%100 != 0)
	fmt.Print(leap)
}
