package main
import "fmt"
func main() {
	var p, s, e int

	fmt.Scan(&p)
	fmt.Scan(&s)
	fmt.Scan(&e)

	pos := 0

	for {
		if pos+s >= p {
			fmt.Println(pos, "saiu")
			break
		}

		fmt.Println(pos, pos+s)

		pos = pos + s - e
	}
}