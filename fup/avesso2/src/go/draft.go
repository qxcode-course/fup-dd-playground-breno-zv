package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	for k := 0; k < x; k++ {
		var grito int
		fmt.Scan(&grito)

		pos := -1
		for i := 0; i < n; i++ {
			if v[i] == grito || v[i] == -grito {
				pos = i
				break
			}
		}

		if pos != -1 {
			if pos > 0 {
				v[pos-1] = -v[pos-1]
			}
			if pos < n-1 {
				v[pos+1] = -v[pos+1]
			}
		}
	}

	fmt.Print("[")
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v[i])
	}
	fmt.Println("]")
}