package main
import "fmt"
func main() {
	var suara [21]int
	var x int
	var totalMasuk, totalSah int
	for {
		fmt.Scan(&x)

		if x == 0 {
			break
		}
		totalMasuk++

		if x >= 1 && x <= 20 {
			suara[x]++
			totalSah++
		}
	}
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}