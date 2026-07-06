package main
import "fmt"
func main() {
	var T int
	fmt.Scan(&T)

	jedi := 0
	sith := 0

	for i := 0; i < T; i++ {
		var x int
		fmt.Scan(&x)

		if i < T/2 {
			jedi += x
		} else {
			sith += x
		}
	}

	if jedi > sith {
		fmt.Println("Jedi")
	} else if sith > jedi {
		fmt.Println("Sith")
	} else {
		fmt.Println("Empate")
	}
}