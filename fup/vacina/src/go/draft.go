package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	vacinas := make([]int, n)
	pacientes := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vacinas[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&pacientes[i])
	}

	// Bubble Sort nas vacinas
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if vacinas[j] > vacinas[j+1] {
				vacinas[j], vacinas[j+1] = vacinas[j+1], vacinas[j]
			}
		}
	}

	// Bubble Sort nos pacientes
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if pacientes[j] > pacientes[j+1] {
				pacientes[j], pacientes[j+1] = pacientes[j+1], pacientes[j]
			}
		}
	}

	for i := 0; i < n; i++ {
		if vacinas[i] <= pacientes[i] {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}