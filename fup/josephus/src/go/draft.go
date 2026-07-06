package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n)
	fmt.Scan(&x)

	// cria a lista de jogadores
	jogadores := make([]int, n)
	for i := 0; i < n; i++ {
		jogadores[i] = i + 1
	}

	// encontra a posição inicial do jogador X
	pos := 0
	for jogadores[pos] != x {
		pos++
	}

	// enquanto houver mais de um jogador
	for len(jogadores) > 1 {
		// índice da vítima
		morto := (pos + 1) % len(jogadores)

		// remove a vítima
		jogadores = append(jogadores[:morto], jogadores[morto+1:]...)

		// a espada passa para o próximo vivo
		if morto >= len(jogadores) {
			pos = 0
		} else {
			pos = morto
		}
	}

	fmt.Println(jogadores[0])
}