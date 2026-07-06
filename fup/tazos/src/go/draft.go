package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)

	var atual int
	fmt.Scan(&atual)

	cont := 1
	maior := 1
	resp := []int{}

	for i := 1; i < n; i++ {
		var x int
		fmt.Scan(&x)

		if x == atual {
			cont++
		} else {
			if cont > maior {
				maior = cont
				resp = []int{atual}
			} else if cont == maior {
				resp = append(resp, atual)
			}

			atual = x
			cont = 1
		}
	}

	if cont > maior {
		resp = []int{atual}
	} else if cont == maior {
		resp = append(resp, atual)
	}

	fmt.Print("[ ")
	for i := 0; i < len(resp); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(resp[i])
	}
	fmt.Println(" ]")
}