package main
import "fmt"
func main() {
    var value, parc, porcJu float64

    fmt.Scanf("%f", &value)
    fmt.Scanf("%f", &parc)

    porcJu = ((parc-1)*5)/100

    fmt.Printf("%.2f\n", ((value*porcJu)+value)/parc)
    fmt.Printf("%.2f\n", value*porcJu+value)

}
