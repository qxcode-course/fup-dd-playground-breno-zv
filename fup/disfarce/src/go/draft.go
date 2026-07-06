package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(in, &t)

	for ; t > 0; t-- {
		var ultron, pessoa string
		fmt.Fscan(in, &ultron)
		fmt.Fscan(in, &pessoa)

		ultron = strings.ToLower(ultron)
		pessoa = strings.ToLower(pessoa)

		// conjunto de letras do Ultron
		m := make(map[byte]bool)
		for i := 0; i < len(ultron); i++ {
			m[ultron[i]] = true
		}

		// conta quantas letras da pessoa aparecem no código do Ultron
		cont := 0
		for i := 0; i < len(pessoa); i++ {
			if m[pessoa[i]] {
				cont++
			}
		}

		if cont == len(pessoa) {
			fmt.Println("chefe")
		} else if cont*2 > len(pessoa) {
			fmt.Println("ultron")
		} else {
			fmt.Println("pessoa")
		}
	}
}