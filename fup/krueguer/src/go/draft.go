package main

import (
	"bufio"
	"fmt"
	"os"
)

func ehVogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	in.ReadByte() // consome o '\n'

	for ; n > 0; n-- {
		var s string
		fmt.Fscanln(in, &s)

		melhor := ""
		atual := ""

		for i := 0; i < len(s); i++ {
			if ehVogal(s[i]) {
				atual += string(s[i])
			} else {
				if len(atual) > len(melhor) {
					melhor = atual
				}
				atual = ""
			}
		}

		// verifica a última sequência
		if len(atual) > len(melhor) {
			melhor = atual
		}

		fmt.Println(melhor)
	}
}