/*

- "crianca" se menor que 12 (não use o ç),
- "jovem" se menor que 18,
- "adulto" se menor que 65,
- "idoso" se menor que 1000,
- "mumia" caso contrario (não ponha o acento).

*/

package main
import "fmt"
func main() {
    var idade int
    var pessoa string

    fmt.Scan(&pessoa)
    fmt.Scan(&idade)

    if idade < 12{
        fmt.Printf("%s eh crianca\n", pessoa)
    }else if idade >= 12 && idade < 18{
        fmt.Printf("%s eh jovem\n", pessoa)
    }else if idade >= 18 && idade < 65{
        fmt.Printf("%s eh adulto\n", pessoa)
    }else if idade >= 65 && idade < 1000{
        fmt.Printf("%s eh idoso\n", pessoa)
    }else{
        fmt.Printf("%s eh mumia\n", pessoa)
    }
}
