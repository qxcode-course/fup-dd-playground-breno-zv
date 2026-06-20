package main
import "fmt"
func main() {
	var B, T int

	fmt.Scanln(&B)
	fmt.Scanln(&T)

	if B+T > 160 {
		fmt.Println(1)
	}else if B+T < 160{
		fmt.Println(2)
	}else{
		fmt.Println(0)
	}
}