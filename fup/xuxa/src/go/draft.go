package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	frase, _ := reader.ReadString('\n')

	// remove o '\n' se existir
	if len(frase) > 0 && frase[len(frase)-1] == '\n' {
		frase = frase[:len(frase)-1]
	}

	for i := len(frase) - 1; i >= 0; i-- {
		fmt.Print(string(frase[i]))
	}
	fmt.Println()
}