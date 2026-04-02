# <h1 align="center">Laporan Praktikum Modul 4 - Prosedur </h1>
<p align="center">Satriya Wahyu Prakoso - 109082500219</p>

## Unguided 

### 1. Soal1
#### soal1.go

```go
package main
import "fmt"
func factorial(n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++ {
		*hasil = *hasil * i
	}
}

func permutation(n, r int, hasil *int) {
	var factN, factNMinusR int
	factorial(n, &factN)
	factorial(n-r, &factNMinusR)
	*hasil = factN / factNMinusR
}

func combination(n, r int, hasil *int) {
	var factN, factR, factNMinusR int
	factorial(n, &factN)
	factorial(r, &factR)
	factorial(n-r, &factNMinusR)
	*hasil = factN / (factR * factNMinusR)
}

func main() {
	var a, b, c, d int
	var permResult, combResult int

	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &permResult)
	combination(a, c, &combResult)
	fmt.Println(permResult, combResult)

	permutation(b, d, &permResult)
	combination(b, d, &combResult)
	fmt.Println(permResult, combResult)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalA.png)

##### Penjelasan
Program ini digunakan untuk 

### 2. Soal2
#### soal2.go

```go
package main
import (
	"fmt"
)
func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0
	for i := 0; i < 8; i++ {
		fmt.Scan(&waktu)
		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var nama string
	var pemenang string
	var soal, skor int
	var maxSoal, minSkor int
	maxSoal = -1
	minSkor = 999999
	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}
	fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalB.png)

##### Penjelasan
Program ini digunakan untuk 

### 3. Soal3
#### soal3.go

```go
package main
import (
	"fmt"
)
func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n)
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(n) 
}

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)
	cetakDeret(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalC1.png)

##### Deskripsi Program
Program ini digunakan untuk
