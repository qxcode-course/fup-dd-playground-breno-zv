package main

import "fmt"

func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func mmc(a, b int) int {
	return (a * b) / mdc(a, b)
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)

	fmt.Println(mmc(n, m))
}