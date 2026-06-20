package main
import "fmt"
func main() {
	var N, X, Y, S int
	var C string

	fmt.Scanln(&N)
	fmt.Scanln(&X)
	fmt.Scanln(&Y)
	fmt.Scanln(&C)
	fmt.Scanln(&S)

	switch C{
	case "R":
		X = (X+S)%N
	case "L":
		X = ((X-S)%N+N)%N
	case "D":
		Y = (Y+S)%N
	case "U":
		Y = ((Y-S)%N+N)%N
	}

	fmt.Println(X, Y)
}