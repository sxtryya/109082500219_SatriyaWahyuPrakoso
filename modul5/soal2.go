package main
import "fmt"
func Bintang(baris, n int) {
	if baris > n {
		return
	}
	for i := 0; i < baris; i++ {
		fmt.Print("*")
	}
	fmt.Println()
	Bintang(baris+1, n)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	Bintang(1, n)
}