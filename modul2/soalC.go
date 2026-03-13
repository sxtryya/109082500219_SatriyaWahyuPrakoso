package main
import "fmt"
func main() {
	var beratawal, kilo, sisa int64
	var totalkilo, totalsisa, totalakhir int64
	fmt.Scan(&beratawal)
	kilo = beratawal / 1000
	totalkilo = kilo * 10000
	sisa = beratawal - (kilo * 1000)
	if sisa >= 500 {
		totalsisa = sisa * 5
	} else {
		totalsisa = sisa * 15
	}
	if kilo > 10 {
		totalakhir = totalkilo
	} else {
		totalakhir = totalkilo + totalsisa
	}
	fmt.Println("Berat parsel (gram) : ", beratawal)
	fmt.Printf("Detail berat : %d kg + %d gr \n", kilo, sisa)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d \n", totalkilo, totalsisa)
	fmt.Print("Total biaya : Rp. ", totalakhir)
}