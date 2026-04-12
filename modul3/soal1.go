package main
import "fmt"
func factorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(x, y int) int {
	var permutasi = factorial(x) / factorial(x-y)
	return permutasi
}

func combination(x, y int) int {
	var kombinasi = factorial(x) / (factorial(y) * factorial(x-y))
	return kombinasi
}
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}