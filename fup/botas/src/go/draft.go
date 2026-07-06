package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var esquerda [61]int
	var direita [61]int

	for i := 0; i < n; i++ {
		var tamanho int
		var lado string

		fmt.Scan(&tamanho, &lado)

		if lado == "E" {
			esquerda[tamanho]++
		} else {
			direita[tamanho]++
		}
	}

	pares := 0

	for i := 30; i <= 60; i++ {
		if esquerda[i] < direita[i] {
			pares += esquerda[i]
		} else {
			pares += direita[i]
		}
	}

	fmt.Println(pares)
}