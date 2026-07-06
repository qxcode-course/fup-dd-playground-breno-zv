package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	freq := make(map[int]int)

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)
		freq[x]++
	}

	// quantidade de elementos diferentes
	fmt.Println(len(freq))

	// encontra a maior frequência
	maior := 0
	for _, v := range freq {
		if v > maior {
			maior = v
		}
	}

	// coleta todos os elementos com frequência máxima
	var resp []int
	for k, v := range freq {
		if v == maior {
			resp = append(resp, k)
		}
	}

	sort.Ints(resp)

	// imprime os elementos
	for i, v := range resp {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}