package main
import "fmt"
func fx(n int) int {
	var hasil int = n * n
	return hasil
}

func gx(n int) int {
	var hasil int = n - 2
	return hasil
}

func hx(n int) int {
	var hasil int = n + 1
	return hasil
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}