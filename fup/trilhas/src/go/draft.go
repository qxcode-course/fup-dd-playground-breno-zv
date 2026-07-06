package main

import "fmt"

func esforco(v []int) int {
	ida := 0
	for i := 1; i < len(v); i++ {
		if v[i] > v[i-1] {
			ida += v[i] - v[i-1]
		}
	}

	volta := 0
	for i := len(v) - 2; i >= 0; i-- {
		if v[i] > v[i+1] {
			volta += v[i] - v[i+1]
		}
	}

	if ida < volta {
		return ida
	}
	return volta
}

func main() {
	var n int
	fmt.Scan(&n)

	melhorTrilha := 1
	menorEsforco := -1

	for t := 1; t <= n; t++ {
		var m int
		fmt.Scan(&m)

		trilha := make([]int, m)
		for i := 0; i < m; i++ {
			fmt.Scan(&trilha[i])
		}

		e := esforco(trilha)

		if menorEsforco == -1 || e < menorEsforco {
			menorEsforco = e
			melhorTrilha = t
		}
	}

	fmt.Println(melhorTrilha)
}