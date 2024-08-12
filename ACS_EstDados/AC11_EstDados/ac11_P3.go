package main
 
import (
    "fmt"
)
 
func main() {
    var C, P, F int
    fmt.Scanln(&C, &P, &F)
    
    suficiente(C, P, F)
}

func suficiente(C, P, F int) {
    totalFolhasNecessarias := C * F
    if P >= totalFolhasNecessarias {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}