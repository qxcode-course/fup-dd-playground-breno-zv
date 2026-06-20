package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	vet := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i])
	}

	fmt.Print("[")

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(vet[i])
	}

	fmt.Println("]")
}