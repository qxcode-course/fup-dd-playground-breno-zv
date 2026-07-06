package main

import "fmt"

func main() {
	var m int
	fmt.Scan(&m)

	h := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&h[i])
	}

	esqDir := 0
	for i := 1; i < m; i++ {
		if h[i] > h[i-1] {
			esqDir += h[i] - h[i-1]
		}
	}

	dirEsq := 0
	for i := m - 2; i >= 0; i-- {
		if h[i] > h[i+1] {
			dirEsq += h[i] - h[i+1]
		}
	}

	if esqDir < dirEsq {
		fmt.Println(esqDir)
	} else {
		fmt.Println(dirEsq)
	}
}