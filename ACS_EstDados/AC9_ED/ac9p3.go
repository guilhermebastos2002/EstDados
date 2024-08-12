package main

import (
    "fmt"
)

func decifrarCifra(cifra string, deslocamento int) string {
    resultado := ""

    for _, char := range cifra {
        // Convertendo para valor ASCII
        ascii := int(char)

        // Deslocando de acordo com o número dado
        ascii -= deslocamento

        // Tratando o caso de voltar ao início do alfabeto
        if ascii < 65 {
            ascii += 26
        }

        // Adicionando o caractere decifrado ao resultado
        resultado += string(ascii)
    }

    return resultado
}

func main() {
    var casos int
    fmt.Scan(&casos)

    for i := 0; i < casos; i++ {
        var cifra string
        var deslocamento int
        fmt.Scan(&cifra, &deslocamento)

        textoDecifrado := decifrarCifra(cifra, deslocamento)
        fmt.Println(textoDecifrado)
    }
}