package main
 
import (
    "fmt"
)
 
func main() {
    
    var A int
    var B int
    var C int
    
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)
    
    if (A >= B + C) {
        fmt.Println("NAO FORMA TRIANGULO")
    }
    
    if (A * A == B * B + C * C) {
        fmt.Println("TRIANGULO RETANGULO")
    }
    
    if (A * A > B * B + C * C) {
        fmt.Println("TRIANGULO OBTUSANGULO")
    }
    
    if (A * A < B * B + C * C) {
        fmt.Println("TRIANGULO ACUTANGULO")
    }
    
    if (A == B && B == C && A == C) {
        fmt.Println("TRIANGULO EQUILATERO")
    }
    
    if (A == B || B == C || A == C) {
        fmt.Println("TRIANGULO ISOSCELES")
    }
}
