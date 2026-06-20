package main
import "fmt"
func main(){
	var M, A, B int

	fmt.Scanln(&M)
	fmt.Scanln(&A)
	fmt.Scanln(&B)

	C := M - A - B

	maior := A

	if B > maior{
		maior = B
	}
	if C > maior {
		maior = C
	}
	fmt.Println(maior)
}