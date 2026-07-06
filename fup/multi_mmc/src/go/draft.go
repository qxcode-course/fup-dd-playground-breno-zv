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

func mmcVetor(v []int) int {
	resp := v[0]
	for i := 1; i < len(v); i++ {
		resp = mmc(resp, v[i])
	}
	return resp
}

func main() {
	var n int
	fmt.Scan(&n)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}

	fmt.Println(mmcVetor(v))
}