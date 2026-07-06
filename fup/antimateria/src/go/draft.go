package main

import "fmt"

func main() {
	var materia, antimat string
	fmt.Scan(&materia)
	fmt.Scan(&antimat)

	i := len(materia) - 1
	j := 0

	// aniquila enquanto os caracteres coincidirem
	for i >= 0 && j < len(antimat) && materia[i] == antimat[j] {
		i--
		j++
	}

	fmt.Println(materia[:i+1] + antimat[j:])
}