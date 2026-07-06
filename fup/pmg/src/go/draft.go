package main
import "fmt"
func media(vet []float64) float64 {
	soma := 0.0

	for _, v := range vet {
		soma += v
	}

	return soma / float64(len(vet))
}

func main() {
	var n int
	fmt.Scan(&n)

	alturas := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&alturas[i])
	}

	m := media(alturas)

	fmt.Printf("%.2f\n", m)

	for i, h := range alturas {
		if h < m {
			fmt.Print("P")
		} else if h > m {
			fmt.Print("G")
		} else {
			fmt.Print("M")
		}

		if i < n-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}