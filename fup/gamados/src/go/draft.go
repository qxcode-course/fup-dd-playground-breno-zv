package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	palavras := strings.Fields(frase)

	ordenado := true
	for i := 0; i < len(palavras)-1; i++ {
		if palavras[i] > palavras[i+1] {
			ordenado = false
			break
		}
	}

	if ordenado {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}