package main

import "fmt"

func main() {
	var uangIDR int
	const kurs = 15000

	fmt.Scan(&uangIDR)

	if uangIDR < 0 {
		fmt.Println("Rp.")
		return
	}

	uangUSD := uangIDR / kurs
	fmt.Printf("%d IDR %d USD\n", uangIDR, uangUSD)
}
