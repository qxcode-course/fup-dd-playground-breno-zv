package main
import "fmt"
func cabe(x, y, h, l int) bool {
	return (x <= h && y <= l) || (x <= l && y <= h)
}

func main() {
	var A, B, C int
	var H, L int

	fmt.Scanln(&A)
	fmt.Scanln(&B)
	fmt.Scanln(&C)
	fmt.Scanln(&H)
	fmt.Scanln(&L)

	if cabe(A, B, H, L) ||
		cabe(A, C, H, L) ||
		cabe(B, C, H, L) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}
