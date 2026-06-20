package main
import "fmt"
func main() {
	var n int
	var sabor, turno string
	chocolate, limao := 0, 0
	manha, tarde := 0, 0

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&sabor, &turno)

		if sabor == "c" {
			chocolate++
		} else if sabor == "l" {
			limao++
		}

		if turno == "m" {
			manha++
		} else if turno == "t" {
			tarde++
		}
	}
	if chocolate > limao {
		fmt.Println("c")
	} else if limao > chocolate {
		fmt.Println("l")
	} else {
		fmt.Println("empate")
	}

	if manha < tarde {
		fmt.Println("m")
	} else if tarde < manha {
		fmt.Println("t")
	} else {
		fmt.Println("empate")
	}
}