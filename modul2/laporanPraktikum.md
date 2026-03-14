# <h1 align="center">Laporan Praktikum Modul 1 - REVIEW ALGORITMA DAN PEMROGRAMAN 1 </h1>
<p align="center">Satriya Wahyu Prakoso - 109082500219</p>

## Unguided 

### 1. [Soal]
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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)

### 2. [Soal]
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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)

### 3. [Soal]
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
![Screenshot Output Unguided 1_1](https://github.com/shellyneu/103112430114_Shellyn/blob/main/modul1/output/output-soal1.png)
