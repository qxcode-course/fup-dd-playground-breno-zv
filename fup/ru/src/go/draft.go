package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ehVogal(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	frase, _ := reader.ReadString('\n')
	frase = strings.TrimRight(frase, "\r\n")

	vogais := ""
	consoantes := ""

	for i := 0; i < len(frase); i++ {
		c := frase[i]

		if c == ' ' {
			continue
		}

		if ehVogal(c) {
			vogais += string(c)
		} else {
			consoantes += string(c)
		}
	}

	fmt.Println(vogais)
	fmt.Println(consoantes)
}