package main

import (
    "fmt"
    "strconv"
)

func calcularValorReal(d, n int) int {
    valorFinal := 0
    multiplicador := 1
    for n > 0 {
        digito := n % 10
        if digito != d {
            valorFinal += digito * multiplicador
            multiplicador *= 10
        }
        n /= 10
    }
    return valorFinal
}

func main() {
    for {
        var d, n int
        _, err := fmt.Scan(&d, &n)
        if err != nil {
            break
        }
        if d == 0 && n == 0 {
            break
        }
        valorReal := calcularValorReal(d, n)
        fmt.Println(valorReal)
    }
}