package main

import "fmt"

func main() {
	var n, m int

	fmt.Scan(&n)
	vetor1 := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor1[i])
	}

	fmt.Scan(&m)
	vetor2 := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&vetor2[i])
	}

	contido := true

	for i := 0; i < n; i++ {
		encontrou := false

		for j := 0; j < m; j++ {
			if vetor1[i] == vetor2[j] {
				encontrou = true
				break
			}
		}

		if !encontrou {
			contido = false
			break
		}
	}

	if contido {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}