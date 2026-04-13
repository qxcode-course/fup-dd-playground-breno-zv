package main
import "fmt"
func main() {
    var h, m, s, seg int64
    fmt.Scan(&seg)
    h = seg/3600
    m = (seg%3600)/60
    s = seg%60


    fmt.Printf("%d:%d:%d\n", h, m, s)
}