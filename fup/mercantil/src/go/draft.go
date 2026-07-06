package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	valores := make([]float64, n)
	chutes := make([]float64, n)
	escolhas := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&valores[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&chutes[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&escolhas[i])
	}

	primeiro := 0
	segundo := 0

	for i := 0; i < n; i++ {
		valor := valores[i]
		chute := chutes[i]
		escolha := escolhas[i]

		if valor == chute {
			primeiro++
		} else if valor > chute {
			if escolha == "M" {
				segundo++
			} else {
				primeiro++
			}
		} else { // valor < chute
			if escolha == "m" {
				segundo++
			} else {
				primeiro++
			}
		}
	}

	if primeiro > segundo {
		fmt.Println("primeiro")
	} else if segundo > primeiro {
		fmt.Println("segundo")
	} else {
		fmt.Println("empate")
	}
}