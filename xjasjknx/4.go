package main
import "fmt"
func isKabisat(tahun int) bool {
	return tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0)
}

func main() {
	var tahunawal, n int
	fmt.Scan(&tahunawal, &n)
	for i:=0 ; i < n ; i++ {
		tahun := tahunawal + i
		if isKabisat(tahun) {
			fmt.Printf("%d, merupakan tahun kabisat\n", tahun)
		} else {
			fmt.Println(tahun)
		}
	}
}