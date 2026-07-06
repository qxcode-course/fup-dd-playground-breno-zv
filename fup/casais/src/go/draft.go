package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	animais := make(map[int]int)
	casais := 0

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x > 0 { // macho
			if animais[-x] > 0 {
				casais++
				animais[-x]--
			} else {
				animais[x]++
			}
		} else { // fêmea
			if animais[-x] > 0 {
				casais++
				animais[-x]--
			} else {
				animais[x]++
			}
		}
	}

	fmt.Println(casais)
}