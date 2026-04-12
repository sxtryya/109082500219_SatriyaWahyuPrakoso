package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	if n%2 == 0 {
		for i :=n ; i >= 1 ; i-- {
			fmt.Println(i, " ")
		}
	} else {
		for i :=n ; i >= 1 ; i-- {
			fmt.Println(i, i, "-")
		}
	}
}