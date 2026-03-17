package main
import "fmt"
func main() {
    var in1, in2, media float32
    fmt.Scan(&in1, &in2)

    media = (in1+in2)/2

    fmt.Printf("%.1f\n", media)
}


