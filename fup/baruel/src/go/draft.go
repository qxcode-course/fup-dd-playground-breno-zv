package main
import "fmt"
func main() {
	var total, qtd int
	fmt.Scan(&total)
	fmt.Scan(&qtd)

	possui := make([]bool, total+1)
	repetiu := false
	faltou := false

	ultimo := -1

	for i := 0; i < qtd; i++ {
		var fig int
		fmt.Scan(&fig)

		if fig == ultimo {
			if repetiu {
				fmt.Print(" ")
			}
			fmt.Print(fig)
			repetiu = true
		}

		possui[fig] = true
		ultimo = fig
	}

	fmt.Println()

	for i := 1; i <= total; i++ {
		if !possui[i] {
			if faltou {
				fmt.Print(" ")
			}
			fmt.Print(i)
			faltou = true
		}
	}

	fmt.Println()
}