package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	resultado := 0

	for i := 0; i < n; i++ {
		var digito int
		fmt.Scan(&digito)

		resultado = resultado*10 + digito
	}

	fmt.Println(resultado)
}