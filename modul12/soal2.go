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
	ketua := 1
	wakil := 2
	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua {
			if suara[i] > suara[wakil] ||
				(suara[i] == suara[wakil] && i < wakil) {
				wakil = i
			}
		}
	}

	if suara[wakil] > suara[ketua] {
		ketua, wakil = wakil, ketua
	}
	fmt.Println("Suara masuk:", totalMasuk)
	fmt.Println("Suara sah:", totalSah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}