package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    in := bufio.NewReader(os.Stdin)

    texto, _ := in.ReadString('\n')
    texto = strings.TrimSpace(texto)

    comando, _ := in.ReadString('\n')
    comando = strings.TrimSpace(comando)

    switch comando {
    case "M":
        fmt.Println(strings.ToUpper(texto))

    case "m":
        fmt.Println(strings.ToLower(texto))
    }
}