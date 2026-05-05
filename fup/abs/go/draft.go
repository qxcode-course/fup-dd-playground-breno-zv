package main
import (
	"fmt"
)
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	resultado := abs(a - b)
	fmt.Println(resultado)
}