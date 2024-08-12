package main
 
import (
    "fmt"
)
 
func main() {
    var T int
    fmt.Scanln(&T)
    
    for i := 0; i <T; i++ {
        var PA, PB int
        var G1, G2 float64
        fmt.Scanln(&PA, &PB, &G1, &G2)
    
        anos := calcularAnos(PA, PB, G1, G2)
    
        if anos > 100 {
        fmt.Println("Mais de um seculo.")
        } else {
            fmt.Println("%d anos.", anos)
        }
    }
}

func calcularAnos(PA, PB int, G1, G2 float64) int {
    anos := 0
    for PA <= PB {
        PA += int(float64(PA) * (G1 / 100))
        PB += int(float64(PB) * (G2 / 100))
        anos++
        if anos > 100 {
            break
        }
    }
    return anos
}