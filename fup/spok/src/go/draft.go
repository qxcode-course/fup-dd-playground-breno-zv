package main
import "fmt"

func inverter(n int) int {
	invertido := 0

	for n > 0 {
		digito := n % 10
		invertido = invertido*10 + digito
		n = n / 10
	}

	return invertido
}

func palindromo(n int) int {
	if n == inverter(n) {
		return 1
	}
	return 0
}

func main() {
	var id int

	fmt.Scan(&id)

	fmt.Println(palindromo(id))
}