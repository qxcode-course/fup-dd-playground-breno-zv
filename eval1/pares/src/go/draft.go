package main
import "fmt"
func main() {
    var a, b, soma int64

    fmt.Scan(&a, &b)
    if a>b{
        fmt.Println("invalido")
        return
    }
    for i:=a; i<=b; i++{
        if i%2==0{
            soma+=i
        }
    }
    fmt.Println(soma)
}