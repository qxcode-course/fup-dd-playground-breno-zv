package main
import "fmt"
func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print("[")

	primeiro := true

	for i:=0; i<=10; i++{
		if i==N{
			continue
		}
		if primeiro{
			fmt.Print(" ")
			primeiro = false
		}else{
			fmt.Print(" ")
		}
		if i==10 {
			fmt.Print("ceu")
		}else{
			fmt.Print(i)
		}
	}

	fmt.Print(" ]\n")
}