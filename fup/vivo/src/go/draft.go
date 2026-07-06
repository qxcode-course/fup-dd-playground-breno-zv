package main

import "fmt"

func main() {
	teste := 1

	for {
		var P, R int
		fmt.Scan(&P, &R)

		if P == 0 && R == 0 {
			break
		}

		fila := make([]int, P)
		for i := 0; i < P; i++ {
			fmt.Scan(&fila[i])
		}

		for rodada := 0; rodada < R; rodada++ {
			var N, J int
			fmt.Scan(&N, &J)

			novaFila := []int{}

			for i := 0; i < N; i++ {
				var A int
				fmt.Scan(&A)

				if A == J {
					novaFila = append(novaFila, fila[i])
				}
			}

			fila = novaFila
		}

		if teste > 1 {
    fmt.Println()
}
fmt.Printf("Teste %d\n", teste)
	}
}