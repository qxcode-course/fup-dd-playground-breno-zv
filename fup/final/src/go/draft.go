package main
import "fmt"
func main() {
	var n1, n2, nf int

	fmt.Scan(&n1)
	fmt.Scan(&n2)
	fmt.Scan(&nf)

	media := float64(n1+n2) / 2.0

	if media > 7 {
		fmt.Println("aprovado")
	} else if media <= 4 {
		fmt.Println("reprovado")
	} else {
		novaMedia := (media + float64(nf)) / 2.0

		if novaMedia >= 5 {
			fmt.Println("aprovado na final")
		} else {
			fmt.Println("reprovado na final")
		}
	}
}