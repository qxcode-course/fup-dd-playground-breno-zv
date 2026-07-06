package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	leds := [10]int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)

		total := 0
		for _, c := range s {
			total += leds[c-'0']
		}

		fmt.Printf("%d leds\n", total)
	}
}