package main
import (
    "fmt"
    "math"
)
func main() {
    var x1, x2, y1, y2, calc float64

    fmt.Scan(&x1, &y1)
    fmt.Scan(&x2, &y2)

    calc = math.Sqrt(((x1 - x2)*(x1 - x2)) + ((y1 - y2)*(y1 - y2)))

    fmt.Printf("%.2f\n", calc)
}
