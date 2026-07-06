package main

import "fmt"

func imprimir(v []int) {
	fmt.Print("[ ")

	for i := 0; i < len(v); i++ {
		fmt.Print(v[i])
		if i < len(v)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println(" ]")
}

func main() {
	var n, x int
	var alunos []int
	var servidores []int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		if x%2 != 0 {
			alunos = append(alunos, x)
		} else {
			servidores = append(servidores, x)
		}
	}

	imprimir(alunos)
	imprimir(servidores)
}