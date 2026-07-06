package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var ultron string
	linha, _ := in.ReadString('\n')
	ultron = strings.TrimSpace(strings.ToLower(linha))

	linha, _ = in.ReadString('\n')
	linha = strings.TrimSpace(strings.ToLower(linha))

	// cria conjunto de letras do ultron
	m := make(map[rune]bool)
	for _, c := range ultron {
		m[c] = true
	}

	pessoas := strings.Fields(linha)
	resp := make([]string, 0, len(pessoas))

	for _, p := range pessoas {
		cont := 0
		for _, c := range p {
			if m[c] {
				cont++
			}
		}

		if cont == len(p) {
			resp = append(resp, "chefe")
		} else if cont*2 > len(p) {
			resp = append(resp, "ultron")
		} else {
			resp = append(resp, "pessoa")
		}
	}

	fmt.Println(strings.Join(resp, " "))
}