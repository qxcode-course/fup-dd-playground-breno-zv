package main

import "fmt"

func main() {
	var a, b rune
	var op string

	fmt.Scanf("%c\n", &a)
	fmt.Scanln(&op)
	fmt.Scanf("%c", &b)

	x := int(a - 'a')
	y := int(b - 'a')

	if op == "+" {
		fmt.Printf("%c\n", rune((x+y)%26)+'a')
	} else {
		fmt.Printf("%c\n", rune((x-y+26)%26)+'a')
	}
}