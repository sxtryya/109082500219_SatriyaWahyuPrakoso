# <h1 align="center">Laporan Praktikum Modul 1 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Satriya Wahyu Prakoso - 109082500219</p>

## Unguided 

### 1. SoalA
#### soalA.go

```go
package main
import "fmt"
func main() {
	var (satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalA.png)

##### Deskripsi Program
Program ini digunakan untuk menukar posisi tiga string yang dimasukkan oleh user dengan bantuan variabel sementara (temp). Jadi program ini menunjukkan perubahan urutan data. Program ditulis dengan bahasa Go.

package main digunakan agar program bisa dijalankan.
import “fmt” digunakan untuk import fungsi-fungsi yang digunakan untuk membaca
input dan mencetak/menampilkan output.
func main () digunakan untuk tempat program ditulis, semua program yang ingin
dijalankan. Ini wajib karena tanpa ini program tidak bisa dijalankan. 

Variabel "satu, dua, tiga, temp" digunakan untuk menyimpan input bertipe string.
Pertama-tama program akan meminta user untuk memasukkan input string sebanyak 3 kali satu per satu. Input pertama akan mengubah isi dari variabel "satu" menjadi kata yang di input, input kedua akan mengubah isi dari variabel "dua" menjadi kata yang di input, input ketiga akan mengubah isi dari variabel "tiga" menjadi kata yang di input.

pada saat ini di baris output fmt.Println("Output awal = " + satu + " " + dua + " " + tiga) akan menggunakan input-input yang baru saja dimasukkan dan akan menampilkan hasilnya, tetapi ditampilkan bersama baris output yang nanti. Setelah melewati ini variabel "temp" akan berubah menjadi sama dengan variabel "satu", lalu variabel "satu" akan berubah menjadi sama dengan variabel "dua", lalu variabel "dua" akan berubah menjadi sama dengan variabel "temp". 

jadi setelah ini semua variabel sudah terganti dan baris output fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga) akan menggunakannya. Disini programnya berakhir dan program akan menampilkan hasil output awal dan akhir.

### 2. SoalB
#### soalB.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalB.png)

##### Deskripsi Program
Program ini digunakan untuk memeriksa apakah input pengguna selama 5 percobaan sama dengan urutan warna yang sudah ditentukan. Program ditulas dengan bahasa Go.

package main digunakan agar program bisa dijalankan.
import “fmt” digunakan untuk import fungsi-fungsi yang digunakan untuk membaca
input dan mencetak/menampilkan output.
func main () digunakan untuk tempat program ditulis, semua program yang ingin
dijalankan. Ini wajib karena tanpa ini program tidak bisa dijalankan. 

Variabel "urutan, masuk1, masuk2, masuk3, masuk4, input" digunakan untuk menyimpan data bertipe string sedangkan variabel "hasil" digunakan untuk menyimpan data bertipe boolean. Pertama-tama keadaan variabel "hasil" di set menjadi true lalu isi variabel urutan diubah menjadi "merah kuning hijau ungu".

pada bagian for dibuat variabel "i" untuk penghitung jumlah perulangan, disini di set menjadi 1 dahulu, selama "i" kurang dari atau sama dengan 5 maka program akan tetap berjalan. Pada setiap perulangan program akan menampilkan baris percobaan, pada fmt.Printf("Percobaan %d: ", i) nilai i akan bertambah 1 setiap perulangan jadi pada perulangan pertama akan menampilkan "Percobaan 1" lalu perulangan kedua akan menampilkan "Percobaan 2" dan seterusnya. Pada setiap perulangan user harus memasukkan input, kata pertama akan disimpan di variabel "masuk1", kata kedua akan disimpan di variabel "masuk2" , kata ketiga akan disimpan di variabel "masuk3" , kata keempat akan disimpan di variabel "masuk4". Jadi setelah input dimasukkan variabel "input" akan menggabungkan kata-kata tadi menjadi 1 tapi ditambahkan spasi setiap kata. Setelah ini alam di cek apakah urutanya sama menggunakan "if urutan != input {hasil = false}" yang berarti jika isi variabel urutan tidak sama dengan isi variabel input maka variabel "hasil" akan berubah menjadi false.

Setelah perulangan selesai program akan menampilkan output "Berhasil : " disertai status variabel hasil, true jika semua urutannya benar dan false jika ada urutan yang salah.
}

### 3. SoalC
#### soalC.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3_1](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalC1.png)

![Screenshot Output Unguided 3_2](https://github.com/sxtryya/109082500219_SatriyaWahyuPrakoso/blob/main/modul2/output/output-soalC2.png)

##### Deskripsi Program
Program ini digunakan untuk menghitung biaya pengiriman parsel berdasarkan berat dalam gram. Program akan memisahkan berat menjadi kilogram dan sisa gram, lalu menghitung biaya sesuai aturan tertentu. Program ditulis dengan bahasa Go.

package main digunakan agar program bisa dijalankan.
import “fmt” digunakan untuk import fungsi-fungsi yang digunakan untuk membaca
input dan mencetak/menampilkan output.
func main () digunakan untuk tempat program ditulis, semua program yang ingin
dijalankan. Ini wajib karena tanpa ini program tidak bisa dijalankan. 

Variabel "beratawal, kilo, sisa, totalkilo, totalsisa, totalakhir" digunakan untuk menyimpan data bertipe integer. Pertama-tama user harus memasukkan input berupa berat awal dari parselnya. Setelah input dimasukkan variabel "kilo" akan berganti menjadi hasil berat awal dibagi 1000, lalu variabel "totalkilo" akan berganti menjadi hasil variabel "kilo" dikali 10000, lalu variabel "sisa" akan berganti menjadi hasil berat awal dikurangi (variabel kilo dikali 1000). Jika nilai variabel "sisa" lebih dari sama dengan 500 maka variabel "totalsisa" akan berganti menjadi hasil sisa dikali 5, kalau dibawah 500 maka variabel "totalsisa" akan berganti menjadi hasil sisa dikali 15, jadi kalau 500 atau diatasnya maka dihitung 5 rupiah per gram sedangkan kalau dibawah 500 maka dihitung 15 rupiah per gram.

Lalu jika varibel "kilo" lebih dari 10 maka nilai variabel "totalakhir" berubah menjadi sama dengan isi variabel "totalkilo", jika varibel "kilo" kurang dari 10 maka nilai variabel "totalakhir" berubah menjadi sama dengan isi variabel "totalkilo" ditambah isi variabel "totalsisa".

Setelah itu program akan menampilkan 4 baris output. Di baris pertama akan menampilkan "Berat parsel (gram) : " beserta isi dari variabel "beratawal", disini menggunakan println agar menampilkan output ke layar dan langsung pindah ke baris baru. Di baris kedua menampilkan "Detail berat : %d kg + %d gr \n", %d pertama di isi dengan isi variabel "kilo", %d kedua di isi dengan isi variabel "sisa", "\n" digunakan agar pindah ke baris baru. Pada baris ketiga menampilkan "Detail biaya : Rp. %d + Rp. %d \n", %d pertama di isi dengan isi variabel "totalkilo", %d kedua di isi dengan isi variabel "totalsisa",  "\n" digunakan agar pindah ke baris baru. Di baris keempat menampilkan "Total biaya : Rp. " beserta isi dari variabel "totalakhir". 
