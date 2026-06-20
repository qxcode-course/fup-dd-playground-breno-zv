package main
import "fmt"
func main() {
	var p, n, num int
	contador := 0

	fmt.Scan(&p, &n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num == p {
			contador++
		}
	}

	fmt.Println(contador)
}