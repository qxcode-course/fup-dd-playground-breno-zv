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

	trecho, _ := reader.ReadString('\n')
	trecho = strings.TrimRight(trecho, "\r\n")

	cont := 0

	for i := 0; i+len(trecho) <= len(texto); i++ {
		if texto[i:i+len(trecho)] == trecho {
			cont++
		}
	}

	fmt.Println(cont)
}