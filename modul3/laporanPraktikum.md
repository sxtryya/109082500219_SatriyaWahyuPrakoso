# <h1 align="center">Laporan Praktikum Modul 3 - FUNGSI </h1>
<p align="center">Satriya Wahyu Prakoso - 109082500219</p>

## Unguided 

### 1. Soal1
#### soal1.go

```go
package main
import "fmt"
func factorial(n int) int {
	var hasil int = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutation(x, y int) int {
	var permutasi = factorial(x) / factorial(x-y)
	return permutasi
}

func combination(x, y int) int {
	var kombinasi = factorial(x) / (factorial(y) * factorial(x-y))
	return kombinasi
}

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
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
package main
import "fmt"
func fx(n int) int {
	var hasil int = n * n
	return hasil
}

func gx(n int) int {
	var hasil int = n - 2
	return hasil
}

func hx(n int) int {
	var hasil int = n + 1
	return hasil
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalC1.png)

##### Deskripsi Program
Program ini digunakan untuk
