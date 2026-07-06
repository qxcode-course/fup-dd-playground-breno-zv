package main

import "fmt"

type Item struct {
	valor  int
	indice int
}

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]Item, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i].valor)
		v[i].indice = i
	}

	// Bubble Sort pelo valor
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if v[j].valor > v[j+1].valor {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}

	fmt.Print("[ ")
	for i := 0; i < n; i++ {
		fmt.Print(v[i].indice)
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}