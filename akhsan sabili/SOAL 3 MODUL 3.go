package main 
import "fmt"
func main(){
	var x int
	fmt.Scan(&x)
	tahun := x%400 == 0 || x%4 == 0 && x%100 != 0
	fmt.Println("kabisat :", tahun) 
}