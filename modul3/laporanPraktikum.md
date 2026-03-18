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
![Screenshot Output Unguided 1_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul3/output/output-soal1.png)

##### Penjelasan
Program ini digunakan untuk menghitung permutasi dan kombinasi dari 2 pasang bilangan yang dimasukkan. Program ditulis degan bahasa Go.

package main digunakan agar program bisa dijalankan. import "fmt" digunakan untuk mengimpor fungsi yang dipakai untuk membaca input dan menampilkan output. func main() digunakan sebagai tempat utama program dijalankan, karena tanpa fungsi ini program tidak bisa dijalankan.

Fungsi factorial(n int) digunakan untuk menghitung nilai faktorial dari suatu bilangan. Di dalam fungsi ini terdapat variabel hasil yang diinisialisasi dengan nilai 1. Kemudian digunakan perulangan for dari 1 sampai n, di mana setiap nilai akan dikalikan ke variabel hasil. Setelah perulangan selesai, nilai faktorial dikembalikan.

Fungsi permutation(x, y int) digunakan untuk menghitung permutasi. Di dalamnya terdapat variabel permutasi yang menyimpan hasil perhitungan dengan rumus x! / (x−y)!. Perhitungan ini menggunakan fungsi factorial, lalu hasilnya dikembalikan.

Fungsi combination(x, y int) digunakan untuk menghitung kombinasi. Di dalamnya terdapat variabel kombinasi yang menyimpan hasil perhitungan dengan rumus x! / (y! × (x−y)!). Fungsi ini juga menggunakan fungsi factorial.

Pada bagian main, dibuat variabel a, b, c, d bertipe integer untuk menyimpan input dari pengguna. Input dibaca menggunakan fmt.Scan(&a, &b, &c, &d).
Setelah itu program akan menghitung dan menampilkan permutasi dan kombinasi dari (a, c) lalu permutasi dan kombinasi dari (b, d)

Output ditampilkan menggunakan fmt.Println, di baris pertama berisi hasil dari (a, c) dan baris kedua berisi hasil dari (b, d).
### 2. Soal2
#### soal2.go

```go
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
![Screenshot Output Unguided 2_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul3/output/output-soal2.png)

##### Penjelasan
Program ini digunakan untuk menghitung hasil dari komposisi beberapa fungsi berdasarkan input yang diberikan oleh pengguna. Program ditulis dengan bahasa Go.

package main digunakan agar program bisa dijalankan. import "fmt" digunakan untuk mengimpor fungsi yang dipakai untuk membaca input dan menampilkan output. func main() digunakan sebagai tempat utama program dijalankan, tanpa fungsi ini program tidak bisa dijalankan.

Fungsi fx(n int) digunakan untuk menghitung nilai kuadrat dari suatu bilangan. Di dalam fungsi ini terdapat variabel hasil yang berisi n × n, lalu nilai tersebut dikembalikan.

Fungsi gx(n int) digunakan untuk mengurangi suatu bilangan dengan 2. Di dalamnya terdapat variabel hasil yang berisi n − 2, kemudian dikembalikan.

Fungsi hx(n int) digunakan untuk menambahkan 1 pada suatu bilangan. Di dalamnya terdapat variabel hasil yang berisi n + 1, lalu dikembalikan.

Pada bagian main, dibuat variabel a, b, c bertipe integer untuk menyimpan input dari pengguna. Input dibaca menggunakan fmt.Scan(&a, &b, &c).
Setelah itu program akan menghitung dan menampilkan hasil dari komposisi fungsi sebagai berikut:

Baris pertama: fx(gx(hx(a))), artinya nilai a akan diproses dulu oleh hx, kemudian hasilnya diproses oleh gx, lalu hasil akhirnya diproses oleh fx.
Baris kedua: gx(hx(fx(b))), artinya nilai b diproses oleh fx, lalu hx, kemudian gx.
Baris ketiga: hx(fx(gx(c))), artinya nilai c diproses oleh gx, lalu fx, kemudian hx.

Lalu 3 baris output dari setiap komposisi fungsi ditampilkan menggunakan fmt.Println.

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
![Screenshot Output Unguided 3_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul3/output/output-soal3.png)

##### Deskripsi Program
Program ini digunakan untuk menentukan apakah sebuah titik berada di dalam atau di luar dua lingkaran berdasarkan input yang diberikan oleh pengguna. Program ditulis dengan bahasa Go.

package main digunakan agar program bisa dijalankan. import "fmt" digunakan untuk input dan output, sedangkan import "math" digunakan untuk operasi matematika seperti akar dan pangkat. func main() digunakan sebagai tempat utama program dijalankan.

Fungsi jarak(a, b, c, d int) digunakan untuk menghitung jarak antara 2 titik, yaitu titik (a, b) dan (c, d). Di dalam fungsi ini terdapat variabel hasil yang menghitung jarak menggunakan rumus jarak Euclidean, yaitu akar dari ((a−c)² + (b−d)²). Perhitungan ini menggunakan fungsi math.Pow untuk pangkat dan math.Sqrt untuk akar, lalu hasilnya dikembalikan dalam bentuk integer.

Fungsi didalam(cx, cy, r, x, y int) digunakan untuk mengecek apakah titik (x, y) berada di dalam lingkaran yang berpusat di (cx, cy) dengan jari-jari r. Di dalam fungsi ini terdapat variabel hasil bertipe boolean. Jika jarak antara titik dan pusat lingkaran kurang dari atau sama dengan r, maka hasil bernilai true (berarti titik di dalam lingkaran). Jika tidak, maka bernilai false.

Pada bagian main, dibuat variabel x1, y1, r1 untuk lingkaran pertama, x2, y2, r2 untuk lingkaran kedua, serta x, y untuk titik yang akan dicek. Input dibaca menggunakan fmt.Scanln.
Setelah itu program akan melakukan pengecekan:

Jika titik berada di dalam kedua lingkaran maka akan menampilkan "Titik di dalam lingkaran 1 dan 2"

Jika hanya di dalam lingkaran 2 maka akan menampilkan "Titik di dalam lingkaran 2"

Jika hanya di dalam lingkaran 1 maka akan menampilkan "Titik di dalam lingkaran 1"

Jika tidak di dalam keduanya maka akan menampilkan "Titik di luar lingkaran 1 dan 2"

Hasil akhir ditampilkan menggunakan fmt.Print sesuai kondisi yang terpenuhi.
