package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func vogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	linha, _ := reader.ReadString('\n')
	linha = strings.TrimSpace(linha)

	palavras := strings.Fields(linha)

	if len(palavras) == 0 {
		return
	}

	resp := palavras[0]

	for i := 1; i < len(palavras); i++ {
		p := palavras[i]

		if vogal(resp[len(resp)-1]) && vogal(p[0]) {
			// pula todas as vogais iniciais da próxima palavra
			j := 0
			for j < len(p) && vogal(p[j]) {
				j++
			}
			resp += p[j:]
		} else {
			resp += " " + p
		}
	}

	fmt.Println(resp)
}