package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	linha, _ := in.ReadString('\n')
	linha = strings.TrimSpace(linha)

	palavras := strings.Fields(linha)

	for i, p := range palavras {
		tipo := "int"

		for _, c := range p {
			if unicode.IsLetter(c) {
				tipo = "str"
				break
			}
		}

		if tipo != "str" {
			for _, c := range p {
				if c == '.' {
					tipo = "float"
					break
				}
			}
		}

		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(tipo)
	}
	fmt.Println()
}