package main
import "fmt"
func main() {
	var n, calorias int
	soma := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&calorias)
        soma += calorias
	}

	media := float64(soma) / float64(n)

	fmt.Printf("%.1f\n", media)
}