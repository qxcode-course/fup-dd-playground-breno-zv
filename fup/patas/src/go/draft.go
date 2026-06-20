package main
import "fmt"
func main() {
	var chico, cebolinha, n int
	var animal string
	totalPatas := 0

	fmt.Scan(&chico)
	fmt.Scan(&cebolinha)
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&animal)

		switch animal {
		case "v", "c":
			totalPatas += 4
		case "g":
			totalPatas += 2
		}
	}

	fmt.Println(totalPatas)

	difChico := chico - totalPatas
	if difChico < 0 {
		difChico = -difChico
	}

	difCebolinha := cebolinha - totalPatas
	if difCebolinha < 0 {
		difCebolinha = -difCebolinha
	}

	if difChico < difCebolinha {
		fmt.Println("Chico Bento")
	} else if difCebolinha < difChico {
		fmt.Println("Cebolinha")
	} else {
		fmt.Println("empate")
	}
}