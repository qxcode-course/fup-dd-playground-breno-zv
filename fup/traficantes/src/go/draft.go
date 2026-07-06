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

	procura, _ := reader.ReadString('\n')
	procura = strings.TrimRight(procura, "\r\n")

	troca, _ := reader.ReadString('\n')
	troca = strings.TrimRight(troca, "\r\n")

	for i := 0; i < len(texto); {
		if i+len(procura) <= len(texto) &&
			texto[i:i+len(procura)] == procura {

			fmt.Print(troca)
			i += len(procura)
		} else {
			fmt.Printf("%c", texto[i])
			i++
		}
	}
	fmt.Println()
}