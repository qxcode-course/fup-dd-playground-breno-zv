package main

import "fmt"

func main() {
	var c, a, viagens int

	fmt.Scanln(&c)
	fmt.Scanln(&a)

	viagens = (a+(c-2))/(c-1)
	fmt.Println(viagens)
}