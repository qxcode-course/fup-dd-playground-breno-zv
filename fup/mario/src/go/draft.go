package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)
	max := 0

	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
		if v[i] > max {
			max = v[i]
		}
	}

	for linha := max; linha >= 1; linha-- {
		for i := 0; i < n; i++ {
			if v[i] >= linha {
				fmt.Print("#")
			} else {
				fmt.Print("_")
			}
		}
		fmt.Println()
	}
}