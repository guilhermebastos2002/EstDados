package main

import (
    "fmt"
)

func contarLEDs(valor string) int {
    totalLEDs := 0

    for _, char := range valor {
        switch char {
        case '0':
            totalLEDs += 6
        case '1':
            totalLEDs += 2
        case '2':
            totalLEDs += 5
        case '3':
            totalLEDs += 5
        case '4':
            totalLEDs += 4
        case '5':
            totalLEDs += 5
        case '6':
            totalLEDs += 6
        case '7':
            totalLEDs += 3
        case '8':
            totalLEDs += 7
        case '9':
            totalLEDs += 6
        }
    }

    return totalLEDs
}

func main() {
    var casos int
    fmt.Scan(&casos)

    for i := 0; i < casos; i++ {
        var valor string
        fmt.Scan(&valor)

        totalLEDs := contarLEDs(valor)
        fmt.Printf("%d leds\n", totalLEDs)
    }
}