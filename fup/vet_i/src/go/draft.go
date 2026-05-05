package main
import "fmt"
func mostrar_vetor(arr []int, sep string) {
    fmt.Print("[")
    for i, valor := range arr {
        if i != 0 {
            fmt.Print(sep)
        }
        fmt.Printf("%v", valor)
    }
    fmt.Print("]\n")
}

func main() {
    var qtd int
    fmt.Scan(&qtd)
    var arr []int = make([]int, qtd)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    mostrar_vetor(arr, " ")
}