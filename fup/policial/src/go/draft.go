package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)

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

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v[i])
	}
	fmt.Println()
}