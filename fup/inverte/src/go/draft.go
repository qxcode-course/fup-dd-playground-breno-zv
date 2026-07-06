package main

import "fmt"

func main() {
	var c byte
	fmt.Scanf("%c", &c)

	if c >= 'a' && c <= 'z' {
		c = c - 'a' + 'A'
	} else if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}

	fmt.Printf("%c\n", c)
}