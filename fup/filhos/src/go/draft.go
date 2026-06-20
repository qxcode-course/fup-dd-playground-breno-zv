package main
import "fmt"
func main() {
	var idade, filhos int

	fmt.Scanln(&idade)
	fmt.Scanln(&filhos)

	for i := 0; i < filhos; i++ {
		fmt.Println(idade+2*i)
	}
}
