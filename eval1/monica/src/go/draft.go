package main
import "fmt"
func main() {
    var m, a, b, c int64
    
    fmt.Scan(&m)
    fmt.Scan(&a)
    fmt.Scan(&b)

    c = m - (a + b)
        
     if a>b && a>c{
            fmt.Println(a)
        }else if b>a && b>c{
            fmt.Println(b)
        }else if c>a && c>b{
            fmt.Println(c)
        }
    }
