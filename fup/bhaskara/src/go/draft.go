package main
import (
	"fmt"
	"math"
)
func ajustaZero(x float64) float64 {
	if math.Abs(x) < 1e-9 {
		return 0
	}
	return x
}
func main() {
	var a, b, c float64

	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	delta := b*b - 4*a*c

	if delta < 0 {
		fmt.Println("nao ha raiz real")
		return
	}

	if math.Abs(delta) < 1e-9 {
		x := -b / (2 * a)
		x = ajustaZero(x)
		fmt.Printf("%.2f\n", x)
		return
	}

	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)

	x1 = ajustaZero(x1)
	x2 = ajustaZero(x2)

	fmt.Printf("%.2f\n", x1)
	fmt.Printf("%.2f\n", x2)
}