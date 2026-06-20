package main
import "fmt"
func main() {
	var dia, hora int

	fmt.Scan(&dia)
	fmt.Scan(&hora)

	trabalhando := false

	if dia >= 2 && dia <= 6 {
    if (hora >= 8 && hora <= 11) || (hora >= 14 && hora <= 17) {
        trabalhando = true
    }
}

if dia == 7 {
    if hora >= 8 && hora <= 11 {
        trabalhando = true
    }
}
	if trabalhando {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
	}
}