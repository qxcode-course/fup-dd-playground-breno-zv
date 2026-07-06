package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func substring(s string, inicio, qtd int) string {
	if inicio < 0 || inicio >= len(s) || qtd < 0 {
		return ""
	}

	fim := inicio + qtd
	if fim > len(s) {
		fim = len(s)
	}

	return s[inicio:fim]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// lê a linha inteira, incluindo espaços
	s, _ := reader.ReadString('\n')
	s = strings.TrimRight(s, "\r\n")

	var inicio, qtd int
	fmt.Fscan(reader, &inicio)
	fmt.Fscan(reader, &qtd)

	fmt.Println(substring(s, inicio, qtd))
}