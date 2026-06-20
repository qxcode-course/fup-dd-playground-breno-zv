package main
import "fmt"
func main() {
	var c, m int
	passageiros := 0

	fmt.Scan(&c)

	for {
		fmt.Scan(&m)
		passageiros += m

		if passageiros == 0 {
    fmt.Println("vazio")
} else if passageiros < c {
    fmt.Println("ainda cabe")
} else if passageiros < 2*c {
    fmt.Println("lotado")
} else {
    fmt.Println("hora de partir")
    break
}
	}
}