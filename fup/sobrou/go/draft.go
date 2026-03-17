package main
import "fmt"
func main(){
var in1, in2, quociente, resto int64

fmt.Scan(&in1, &in2)

quociente = in1 / in2
resto = in1 % in2

fmt.Printf("%d %d\n", quociente, resto)

}