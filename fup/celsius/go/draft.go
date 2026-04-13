package main
import "fmt"
func main() {
    var cel, fah float64
    fmt.Scan(&cel)
    fah = (cel*1.8)+32
    fmt.Printf("%.6f\n", fah)
}
