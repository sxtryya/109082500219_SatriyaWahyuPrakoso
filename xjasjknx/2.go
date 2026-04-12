package main
import "fmt"
func theorem(w, x int) int {
	area := w * x
	return area
}
func main() {
	height := 5
	var val int = height * theorem(2,4)
	fmt.Print(val)
}
