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
![Screenshot Output Unguided 1_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul4/output/output-soal1-1.png)

![Screenshot Output Unguided 1_2](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul4/output/output-soal1-2.png)

##### Penjelasan
Program ini digunakan untuk menghitung permutasi dan kombinasi dari 2 pasang bilangan yang dimasukkan. Program ditulis menggunakan bahasa Go.

package main digunakan agar program bisa dijalankan. import "fmt" digunakan untuk mengimpor fungsi input dan output. func main() adalah fungsi utama tempat program mulai dijalankan.

Fungsi factorial(n int, hasil *int) digunakan untuk menghitung nilai faktorial dari suatu bilangan. Fungsi ini menggunakan pointer (*int) untuk menyimpan hasil perhitungan. Di dalam fungsi, nilai *hasil diinisialisasi dengan 1. Kemudian dilakukan perulangan dari 1 sampai n, di mana setiap nilai dikalikan ke *hasil. Setelah selesai, hasil faktorial langsung tersimpan di variabel yang dikirim melalui pointer.

Fungsi permutation(n, r int, hasil *int) digunakan untuk menghitung permutasi dengan rumus n! / (n−r)!. Di dalam fungsi ini dibuat variabel factN dan factNMinusR. Fungsi factorial dipanggil untuk menghitung n! dan (n−r)!. Hasil akhirnya disimpan ke *hasil.

Fungsi combination(n, r int, hasil *int) digunakan untuk menghitung kombinasi dengan rumus n! / (r! × (n−r)!). Di dalamnya terdapat variabel factN, factR, dan factNMinusR. Ketiga nilai faktorial tersebut dihitung menggunakan fungsi factorial, lalu hasil kombinasi disimpan ke *hasil.

Pada bagian main, dibuat variabel a, b, c, d bertipe integer untuk menyimpan input. Lalu ada variabel permResult dan combResult untuk menyimpan hasil perhitungan permutasi dan kombinasi.

Input dibaca menggunakan fmt.Scan(&a, &b, &c, &d). Setelah itu program menghitung permutasi dan kombinasi dari pasangan (a, c) lalu menhitung permutasi dan kombinasi dari pasangan (b, d)

Hasil ditampilkan menggunakan fmt.Println. Baris pertama berisi hasil dari (a, c), dan baris kedua berisi hasil dari (b, d).

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
![Screenshot Output Unguided 2_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul4/output/output-soal2-1.png)

![Screenshot Output Unguided 2_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul4/output/output-soal2-2.png)

##### Penjelasan
Program ini digunakan untuk menentukan pemenang kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Program ditulis dengan bahasa Go.

package main digunakan agar program bisa dijalankan. import "fmt" digunakan untuk mengimpor fungsi yang dipakai untuk membaca input dan menampilkan output. func main() digunakan sebagai tempat utama program dijalankan, karena tanpa fungsi ini program tidak bisa dijalankan.

Fungsi hitungSkor(soal *int, skor *int) digunakan untuk menghitung jumlah soal yang berhasil diselesaikan dan total waktu pengerjaan. Di dalam fungsi ini terdapat variabel waktu untuk menyimpan waktu tiap soal. Variabel *soal dan *skor diinisialisasi dengan nilai 0. Kemudian digunakan perulangan for sebanyak 8 kali untuk membaca waktu dari setiap soal. Kalau waktu yang dimasukkan kurang dari atau sama dengan 300, maka soal dianggap selesai, sehingga *soal ditambah 1 dan *skor ditambah dengan waktu tersebut. Hasil perhitungan langsung disimpan melalui pointer.

Pada bagian main, dibuat variabel nama untuk menyimpan nama peserta dan pemenang untuk menyimpan nama pemenang. Selain itu terdapat variabel soal dan skor untuk menyimpan hasil perhitungan, serta maxSoal dan minSkor untuk menentukan pemenang. maxSoal diinisialisasi dengan -1 dan minSkor dengan nilai besar agar bisa dibandingkan.

Program menggunakan perulangan for tanpa kondisi untuk membaca data beberapa peserta. Nama peserta dibaca menggunakan fmt.Scan(&nama). Kalau nama yang dimasukkan adalah "Selesai", maka perulangan akan berhenti. Kalau tidak, program akan memanggil fungsi hitungSkor untuk menghitung jumlah soal dan skor peserta tersebut.

Setelah itu dilakukan pengecekan pemenang. Kalau jumlah soal lebih banyak dari maxSoal, maka peserta tersebut menjadi pemenang. Kalau jumlah soal sama, maka dipilih peserta dengan skor (waktu) lebih kecil. Kalau kondisi terpenuhi, maka nilai maxSoal, minSkor, dan pemenang akan diperbarui.

Output ditampilkan menggunakan fmt.Println, yang berisi nama pemenang, jumlah soal yang diselesaikan, dan total waktu pengerjaan.

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
![Screenshot Output Unguided 3_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul4/output/output-soal3.png)

##### Deskripsi Program
Program ini digunakan untuk menampilkan deret bilangan berdasarkan aturan tertentu dari sebuah angka yang dimasukkan. Program ditulis menggunakan bahasa Go.

package main digunakan agar program bisa dijalankan. import "fmt" digunakan untuk mengimpor fungsi input dan output. func main() adalah fungsi utama tempat program mulai dijalankan.

Fungsi cetakDeret(n int) digunakan untuk menampilkan deret bilangan. Di dalam fungsi ini digunakan perulangan for yang berjalan selama nilai n tidak sama dengan 1. Pada setiap perulangan, nilai n akan dicetak menggunakan fmt.Printf.

Kemudian dilakukan pengecekan kondisi. Kalau n adalah bilangan genap (n % 2 == 0), maka nilai n dibagi 2 dan Kalau n adalah bilangan ganjil, maka nilai n diubah menjadi 3*n + 1

Proses ini akan terus berulang sampai nilai n menjadi 1. Setelah perulangan selesai, nilai 1 juga dicetak menggunakan fmt.Println.

Di bagian main, dibuat variabel n bertipe integer untuk menyimpan input dari pengguna. Program menampilkan teks "Masukkan angka: " menggunakan fmt.Print, lalu membaca input menggunakan fmt.Scan(&n).

Setelah itu, fungsi cetakDeret(n) dipanggil untuk menampilkan deret bilangan berdasarkan aturan yang telah ditentukan.

Output yang dihasilkan adalah deret angka yang dimulai dari nilai input hingga mencapai 1.
