package main

import (
    "bufio"
    "fmt"
    "os"
    "unicode"
)

func main() {
    in := bufio.NewReader(os.Stdin)

    var n int
    fmt.Fscan(in, &n)
    in.ReadByte() // consome '\n'

    for ; n > 0; n-- {
        linha, _ := in.ReadString('\n')
        if len(linha) > 0 && linha[len(linha)-1] == '\n' {
            linha = linha[:len(linha)-1]
        }

        // descobre o case da primeira letra
        primeiraMaiuscula := false
        for _, c := range linha {
            if c != ' ' {
                primeiraMaiuscula = unicode.IsUpper(c)
                break
            }
        }

        pos := 0
        for _, c := range linha {
            if c == ' ' {
                fmt.Print(" ")
                continue
            }

            deveMaiuscula := (pos%2 == 0 && primeiraMaiuscula) ||
                             (pos%2 == 1 && !primeiraMaiuscula)

            if deveMaiuscula {
                fmt.Printf("%c", unicode.ToUpper(c))
            } else {
                fmt.Printf("%c", unicode.ToLower(c))
            }
            pos++
        }
        fmt.Println()
    }
}