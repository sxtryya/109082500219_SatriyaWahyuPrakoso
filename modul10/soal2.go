package main
import (
	"fmt"
	"math"
)
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	var beratIkan [1000]float64
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := int(math.Ceil(float64(x) / float64(y)))
	totalBeratWadah := make([]float64, jumlahWadah)
	idxWadah := 0
	countIkan := 0
	for i := 0; i < x; i++ {
		totalBeratWadah[idxWadah] += beratIkan[i]
		countIkan++
		if countIkan == y {
			idxWadah++
			countIkan = 0
		}
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%g", totalBeratWadah[i])
		if i < jumlahWadah-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	var sumTotal float64
	for _, berat := range totalBeratWadah {
		sumTotal += berat
	}
	rataRata := sumTotal / float64(jumlahWadah)
	fmt.Printf("%g\n", rataRata)
}