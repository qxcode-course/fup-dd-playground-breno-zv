package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		var n, x int
		fmt.Scan(&n, &x)

		v := make([]int, n)

		pos := -1
		for i := 0; i < n; i++ {
			fmt.Scan(&v[i])
			if v[i] == x || v[i] == -x {
				pos = i
			}
		}

		if pos != -1 {
			if pos-1 >= 0 {
				v[pos-1] = -v[pos-1]
			}
			if pos+1 < n {
				v[pos+1] = -v[pos+1]
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
}