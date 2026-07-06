package main

import "fmt"

func main() {
	var n, x int
	soldados := 0
	rebeldes := 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)

		if x%2 != 0 {
			soldados += x
		} else {
			rebeldes += x
		}
	}

	if soldados > rebeldes {
		fmt.Println("soldados")
	} else if rebeldes > soldados {
		fmt.Println("rebeldes")
	} else {
		fmt.Println("empate")
	}
}