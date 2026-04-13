package main
import "fmt"
func main() {
    var a, b int64

    fmt.Scan(&a)
    fmt.Scan(&b)

    for i:=a; i<b; i++{
        fmt.Println(i)
    }

}
