package main
import "fmt"
func main() {
	var a, b int

	fmt.Scan(&a, &b)
	fmt.Print("[ ")
	primeiro := true
	inicio := a
	fim := b

	for inicio <= b {
		if !primeiro {
			fmt.Print(" ")
		}

		fmt.Print(inicio)
		fmt.Print(" ")
		fmt.Print(fim)

		primeiro = false
		inicio++
		fim--
	}
	fmt.Print(" ]\n")
}