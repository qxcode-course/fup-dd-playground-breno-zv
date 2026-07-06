package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	// Bubble Sort
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}

	var mediana float64

	if n%2 == 1 {
		mediana = v[n/2]
	} else {
		mediana = (v[n/2-1] + v[n/2]) / 2
	}

	fmt.Printf("%.1f\n", mediana)
}