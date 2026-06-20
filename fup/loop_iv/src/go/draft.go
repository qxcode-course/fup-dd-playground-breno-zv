package main
import "fmt"
func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print("[")

	if a < b {
		for i := a; i < b; i++ {
			fmt.Printf(" %d", i)
		}
	} else {
		for i := a; i > b; i-- {
			fmt.Printf(" %d", i)
		}
	}

	fmt.Println(" ]")
}