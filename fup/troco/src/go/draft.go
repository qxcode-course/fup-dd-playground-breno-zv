package main

import "fmt"

func main() {
	var valor float64
	fmt.Scan(&valor)

	// converte para centavos para evitar erros de ponto flutuante
	centavos := int(valor*100 + 0.5)

	opcoes := []int{
		10000, // 100.00
		5000,  // 50.00
		2000,  // 20.00
		1000,  // 10.00
		500,   // 5.00
		200,   // 2.00
		100,   // 1.00
		50,    // 0.50
		25,    // 0.25
		10,    // 0.10
		5,     // 0.05
	}

	for _, moeda := range opcoes {
		qtd := centavos / moeda

		if qtd > 0 {
			fmt.Printf("%d de %.2f\n", qtd, float64(moeda)/100)
			centavos %= moeda
		}
	}

	if centavos > 0 {
		fmt.Printf("Falta %.2f\n", float64(centavos)/100)
	}
}