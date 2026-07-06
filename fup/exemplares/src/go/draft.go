package main 
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	unicos := []int{}

	for i := 0; i < n; i++ {
		var x int
		fmt.Scan(&x)

		existe := false
		for j := 0; j < len(unicos); j++ {
			if unicos[j] == x {
				existe = true
				break
			}
		}

		if !existe {
			unicos = append(unicos, x)
		}
	}
	for i := 0; i < len(unicos); i++ {
		for j := 0; j < len(unicos)-1; j++ {
			if unicos[j] > unicos[j+1] {
				unicos[j], unicos[j+1] = unicos[j+1], unicos[j]
			}
		}
	}

	for i := 0; i < len(unicos); i++ {
		fmt.Print(unicos[i])
		if i < len(unicos)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}