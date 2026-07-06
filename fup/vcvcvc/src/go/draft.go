package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func ehVogal(c rune) bool {
	c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var s string
	fmt.Fscanln(in, &s)

	// Caso a frase tenha espaços
	if len(s) == 0 {
		s, _ = in.ReadString('\n')
	}

	// Lê a linha completa
	in = bufio.NewReader(os.Stdin)
	if s == "" {
		fmt.Fscanln(in, &s)
	}

	for _, c := range s {
		if c == ' ' {
			fmt.Print(" ")
		} else if ehVogal(c) {
			fmt.Print("v")
		} else {
			fmt.Print("c")
		}
	}
}