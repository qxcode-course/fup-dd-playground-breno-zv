package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	linha, _ := reader.ReadString('\n')
	linha = strings.TrimSpace(linha)

	valores := strings.Fields(linha)
	vetor := make([]int, len(valores))

	for i, v := range valores {
		vetor[i], _ = strconv.Atoi(v)
	}

	fmt.Print("[ ")
	for i := len(vetor) - 1; i >= 0; i-- {
		fmt.Print(vetor[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println(" ]")
}