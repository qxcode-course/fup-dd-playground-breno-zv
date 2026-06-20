package main
import "fmt"
func main() {
	var C, bananas, goiabas, mangas, total int

	fmt.Scanln(&C)
	fmt.Scanln(&bananas)
	fmt.Scanln(&goiabas)
	fmt.Scanln(&mangas)

	total = bananas+goiabas+mangas

	minutos := (total+C-1)/C
	fmt.Println(minutos)
}