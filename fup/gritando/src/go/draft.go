package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	texto, _ := reader.ReadString('\n')
	texto = strings.TrimRight(texto, "\r\n")

	for i := 0; i < len(texto); i++ {
		c := texto[i]

		if c >= 'a' && c <= 'z' {
			c = c - 'a' + 'A'
		} else if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
		}

		fmt.Printf("%c", c)
	}
	fmt.Println()
}