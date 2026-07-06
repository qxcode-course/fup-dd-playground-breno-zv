package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("[]")
		return
	}

	fmt.Print("[")

	for i := 0; i < n; i++ {
		var carta int
		fmt.Scan(&carta)

		if i > 0 {
			fmt.Print(", ")
		}

		switch carta {
		case 1:
			fmt.Print("A")
		case 11:
			fmt.Print("J")
		case 12:
			fmt.Print("Q")
		case 13:
			fmt.Print("K")
		default:
			fmt.Print(carta)
		}
	}

	fmt.Println("]")
}