package main
import (
	"fmt"
	"math"
)

func jarak(a, b, c, d int) int {
	var hasil = int(math.Sqrt(math.Pow(float64(a-c), 2) + math.Pow(float64(b-d), 2)))
	return hasil
}

func didalam(cx, cy, r, x, y int) bool {
	var hasil bool
	if r >= jarak(cx, cy, x, y) {
		hasil = true
	} else {
		hasil = false
	}
	return hasil
}

func main() {
	var x1, y1, r1, x2, y2, r2, x, y int
	fmt.Scanln(&x1, &y1, &r1)
	fmt.Scanln(&x2, &y2, &r2)
	fmt.Scanln(&x, &y)

	if didalam(x1, y1, r1, x, y) && didalam(x2, y2, r2, x, y) {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if didalam(x2, y2, r2, x, y) {
		fmt.Print("Titik di dalam lingkaran 2")
	} else if didalam(x1, y1, r1, x, y) {
		fmt.Print("Titik di dalam lingkaran 1")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}