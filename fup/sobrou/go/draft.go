package main

import (
	"fmt"
)

func main() {
	var q1, q2, q3 int //Quantidade dos 3 produtos
	var p1, p2, p3 float64 //Valor dos 3 produtos
	var dinheiro float64 //Quantidade de dinheiro

	//leitura dos valores enteriormente declarados
	fmt.Scan(&q1)
	fmt.Scan(&q2)
	fmt.Scan(&q3)

	fmt.Scan(&p1)
	fmt.Scan(&p2)
	fmt.Scan(&p3)

	fmt.Scan(&dinheiro)
	//relação dos valores
	total := float64(q1)*p1 + float64(q2)*p2 + float64(q3)*p3
	troco := dinheiro - total

	fmt.Printf("%.2f\n", troco) //O valor do troco que a pessoa deve receber, com duas casas decimais.
}