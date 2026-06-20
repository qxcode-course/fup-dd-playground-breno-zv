package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	for i := 1; i <= N; i += 2 {
		fmt.Println(i)
	}

	for i := N - 1; i >= 0; i -= 2 {
		fmt.Println(i)
	}
}