package main
import "fmt"
func main() {
	var A, B int

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	if A>B{
		fmt.Println("invalido")
		return
	}
	soma := 0
	for i := A; i <= B; i++ {
		if i%2 == 0{
            soma += i
		}
	}
	fmt.Println(soma)
}