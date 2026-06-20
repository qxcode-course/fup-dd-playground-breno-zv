package main
import "fmt"
func main() {
	var num, menor int

	for i := 0; i < 5; i++ {
		fmt.Scan(&num)

		if i == 0 || num < menor {
			menor = num
		}
	}

	fmt.Println(menor)
}