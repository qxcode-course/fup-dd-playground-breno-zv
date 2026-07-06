package main
import "fmt"
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var L, N int
	fmt.Scan(&L, &N)

	v := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&v[i])
	}

	ganhador := -1
	melhor := 1 << 30

	perdedor := 0
	pior := abs(v[0])

	for i := 0; i < N; i++ {
		d := abs(v[i])

		if d <= L {
			if d <= melhor {
				melhor = d
				ganhador = i
			}
		}

		if d >= pior {
			pior = d
			perdedor = i
		}
	}

	if ganhador == -1 {
		fmt.Println("nenhum")
	} else {
		fmt.Println(ganhador)
	}

	fmt.Println(perdedor)
}