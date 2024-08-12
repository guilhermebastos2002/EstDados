package main

import (
	"fmt"
	"scan"
)

func encontrarPar(nums []int, alvo int) (int, int) {
	mapaIndice := make(map[int]int)

	for i, num := range nums {
		complemento := alvo - num

		if j, ok := mapaIndice[complemento]; ok {
			return j, i

		}

		mapaIndice[num] = i
	}

	return -1, -1
}

func main() {
	numeros := []int{5, 2, 9, 1, 7}
	fmt.Println("Digite o número alvo: ")
	fmt.Scan(&alvo)

	indice1, indice2 := encontrarPar(numeros, alvo)
    if indice1 != -1 && indice2 != -1 {
        fmt.Printf("Par encontrado: %d + %d = %d\n", numeros[indice1], numeros[indice2], alvo)
    } else {
        fmt.Println("Nenhum par encontrado.")
    }

}