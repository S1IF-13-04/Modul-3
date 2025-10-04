package main

	import "fmt"

	func main() {
		var idr int 
		
		fmt.Print("Rupiah: ")
		fmt.Scanln(&idr)

		dollar := idr/15000

		
		fmt.Print(idr," Rupiah adalah $",dollar)
	}