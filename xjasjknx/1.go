package main
import "fmt"
func split(A int) int {
	split := A * 3
	return split
}

func main(){
	var x ,z int
	fmt.Scan(&x)
	z = 5
	for i:=0; i<=x; i++{
		y := z + i
		if i%4 == 0 {
			fmt.Printf("%d Zonk!\n",i)
		} else if i == 1 {
			fmt.Printf("%d stock gives %d dividents %d splits\n",i , y, split(y))
		} else {
			fmt.Printf("%d stocks give %d dividents %d splits\n",i , y, split(y))
		}
	}
}