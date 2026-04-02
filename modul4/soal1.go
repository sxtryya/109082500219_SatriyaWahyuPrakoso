package main
import "fmt"
func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var factN, factNMinusR int
	factorial(n, &factN)
	factorial(n-r, &factNMinusR)
	*hasil = factN / factNMinusR
}

func combination(n, r int, hasil *int) {
	var factN, factR, factNMinusR int
	factorial(n, &factN)
	factorial(r, &factR)
	factorial(n-r, &factNMinusR)
	*hasil = factN / (factR * factNMinusR)
}

func main() {
	var a, b, c, d int
	var permResult, combResult int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &permResult)
	combination(a, c, &combResult)
	fmt.Println(permResult, combResult)

	permutation(b, d, &permResult)
	combination(b, d, &combResult)
	fmt.Println(permResult, combResult)
}