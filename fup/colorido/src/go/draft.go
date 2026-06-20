package main
import "fmt"
func main() {
	var n int
	var pe string

	fmt.Scan(&n)
	fmt.Scan(&pe)
	fmt.Print("[ ")

	primeiro := true
	for i := 0; i <= 10; i++ {

		if i == n {
			continue
		}
		if i == 10 {
			if n != 10 {
				if !primeiro {
					fmt.Print(" ")
				}
				fmt.Print("ceu")
			}
			break
		}
		if !primeiro {
			fmt.Print(" ")
		}
		fmt.Print(i)
		fmt.Print(pe)
		primeiro = false

		if pe == "d" {
			pe = "e"
		} else {
			pe = "d"
		}
	}
	fmt.Print(" ]\n")
}