package main
import "fmt"

func main() {
	var a, b int64
	fmt.Scan(&a, &b)
	fmt.Print("[ ")

	for {
		if a >= b {
			break
		}
		if a%2 == 0 {
			a++
			continue
		}
		fmt.Print(a)
		if a+2 < b {
			fmt.Print(" ")
		}
		a++
	}
	fmt.Print(" ]\n")
}
