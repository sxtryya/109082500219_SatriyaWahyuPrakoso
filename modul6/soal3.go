package main
import "fmt"
func main() {
	var klubA, klubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)
	var pemenang []string
	i := 1
	for {
		var sA, sB int
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&sA, &sB)

		if sA < 0 || sB < 0 {
			break
		}

		fmt.Printf("Hasil %d : ", i)
		if sA > sB {
			fmt.Println(klubA)
			pemenang = append(pemenang, klubA)
		} else if sB > sA {
			fmt.Println(klubB)
			pemenang = append(pemenang, klubB)
		} else {
			fmt.Println("Draw")
		}
		i++
	}
	fmt.Println("Pertandingan selesai")
	fmt.Println("\nDaftar pemenang:")
	for j, p := range pemenang {
		fmt.Printf("Hasil %d : %s\n", j+1, p)
	}
}