package main

import "fmt"

func main() {
	var n, inferior, superior int
	fmt.Scan(&n, &inferior, &superior)

	cont := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x >= inferior && x <= superior {
			cont++
		}
	}

	fmt.Println(cont)
}