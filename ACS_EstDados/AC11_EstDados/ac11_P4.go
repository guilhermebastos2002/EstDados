package main
 
import (
    "fmt"
)

func maxNumerosMarcados(sequencia []int) int {
	marcados := 1 
	for i := 1; i < len(sequencia); i++ {
		if sequencia[i] != sequencia[i-1] { 
		    marcados++
		}
	}
	return marcados
}
 
func main() {
    var N int

    fmt.Scanln(&N)
    sequencia := make([]int, N)

    for i := 0; i < N; i++ {
        fmt.Scanln(&sequencia[i])
    }
    
    fmt.Println(maxNumerosMarcados(sequencia))
}