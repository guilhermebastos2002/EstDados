package main

import (
    "fmt"
)

func main() {
    var cod1, num1, cod2, num2 int
    var val1, val2, preco float64

    fmt.Scan(&cod1, &num1, &val1)
    fmt.Scan(&cod2, &num2, &val2)

    preco = calcularPreco(val1, num1, val2, num2)

    fmt.Printf("VALOR A PAGAR: R$ %.2f\n", preco)
}

func calcularPreco(val1 float64, num1 int, val2 float64, num2 int) float64 {
    return val1 * float64(num1) + val2 * float64(num2)
}