package main
import "fmt"
func main() {
    var a, b, c int

    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)

    if a>b && a<c || a<b && a>c{
        fmt.Println(a)
    } else if c>b && c<a || c<b && c>a{
        fmt.Println(c)
    }else{
        fmt.Println(b)
    }
}
