package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	risco := 0

	for i := 0; i < n; i++ {
		if fila[i] == 0 {
			temMedico := false

			if i > 0 && fila[i-1] == 1 {
				temMedico = true
			}

			if i < n-1 && fila[i+1] == 1 {
				temMedico = true
			}

			if !temMedico {
				risco++
			}
		}
	}

	fmt.Println(risco)
}