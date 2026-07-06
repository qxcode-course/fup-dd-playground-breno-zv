package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	soma := 0
	ases := 0

	for i := 0; i < n; i++ {
		var carta int
		fmt.Scan(&carta)

		switch carta {
		case 1: // Ás
			soma += 11
			ases++
		case 11, 12, 13: // J, Q, K
			soma += 10
		default:
			soma += carta
		}
	}

	// Ajusta os ases se estourar 21
	for soma > 21 && ases > 0 {
		soma -= 10
		ases--
	}

	fmt.Println(soma)
}