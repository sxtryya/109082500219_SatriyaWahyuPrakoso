package main
import (
	"fmt"
	"math"
)
func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Print("Masukkan elemen: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("a. Semua elemen: ")
	for _, v := range arr {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("c. Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan x (kelipatan): ")
	fmt.Scan(&x)
	fmt.Print("   Indeks kelipatan ", x, ": ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var hapus int
	fmt.Print("e. Masukkan indeks yang dihapus: ")
	fmt.Scan(&hapus)
	arrHapus := append(arr[:hapus], arr[hapus+1:]...)
	fmt.Print("   Array setelah dihapus: ")
	for _, v := range arrHapus {
		fmt.Print(v, " ")
	}
	fmt.Println()

	sum := 0
	for _, v := range arr {
		sum += v
	}
	rata := float64(sum) / float64(n)
	fmt.Printf("f. Rata-rata: %.2f\n", rata)

	varSum := 0.0
	for _, v := range arr {
		varSum += math.Pow(float64(v)-rata, 2)
	}
	stdDev := math.Sqrt(varSum / float64(n))
	fmt.Printf("g. Standar deviasi: %.2f\n", stdDev)

	var cari int
	fmt.Print("h. Masukkan bilangan yang dicari frekuensinya: ")
	fmt.Scan(&cari)
	freq := 0
	for _, v := range arr {
		if v == cari {
			freq++
		}
	}
	fmt.Printf("   Frekuensi %d: %d\n", cari, freq)
}