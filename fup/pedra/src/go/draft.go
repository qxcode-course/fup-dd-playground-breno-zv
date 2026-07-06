package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	vencedor := -1
	melhor := 101 // maior diferença possível é 99

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		// competidor válido
		if a >= 10 && b >= 10 {
			pontos := abs(a - b)

			if vencedor == -1 || pontos < melhor {
				melhor = pontos
				vencedor = i
			}
		}
	}

	if vencedor == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(vencedor)
	}
}