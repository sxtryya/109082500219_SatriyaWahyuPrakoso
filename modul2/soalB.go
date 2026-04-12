package main
import "fmt"
func main() {
	var urutan, masuk1, masuk2, masuk3, masuk4 string
	var hasil bool
	hasil = true
	urutan = "merah kuning hijau ungu"
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scanln(&masuk1, &masuk2, &masuk3, &masuk4)
		input := masuk1 + " " + masuk2 + " " + masuk3 + " " + masuk4
		if urutan != input {
			hasil = false
		}
	}
	fmt.Print("Berhasil : ", hasil)
}