package main
import "fmt"
const NMAX int = 127
type tabel [NMAX]rune
func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	for {
		fmt.Scanf("%c", &ch)
		if ch == '.' {
			break
		}
		if ch == ' ' || ch == '\n' || ch == '\r' {
			continue
		}
		t[*n] = ch
		*n++
		if *n >= NMAX {
			break
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	fmt.Print("Teks        : ")
	isiArray(&tab, &m)
	original := tab
	fmt.Print("Reverse teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)
	fmt.Print("Palindrom   ? ")
	fmt.Println(palindrom(original, m))
}