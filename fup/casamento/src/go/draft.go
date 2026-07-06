package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	menor := x
	maior := x

	for i := 1; i < 5; i++ {
		fmt.Scan(&x)

		if x < menor {
			menor = x
		}

		if x > maior {
			maior = x
		}
	}

	fmt.Println(menor + maior)
}