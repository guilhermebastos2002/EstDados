package main

import "fmt"

type Produto struct {
	nome  string
	preco float64
}

func valorTotal(produtos []Produto, compras map[string]int) float64 {
	total := 0.0
	for nomeProduto, quantidade := range compras {
		for _, produto := range produtos {
			if produto.nome == nomeProduto {
				total += float64(quantidade) * produto.preco
				break
			}
		}
	}
	return total
}

func main() {
	var N int
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var M, P int
		fmt.Scanln(&M)

		produtos := make([]Produto, M)
		for j := 0; j < M; j++ {
			fmt.Scanln(&produtos[j].nome, &produtos[j].preco)
		}

		fmt.Scanln(&P)
		compras := make(map[string]int)
		for j := 0; j < P; j++ {
			var nomeProduto string
			var quantidade int
			fmt.Scanln(&nomeProduto, &quantidade)
			compras[nomeProduto] += quantidade
		}
		valor := valorTotal(produtos, compras)
		fmt.Printf("R$ %.2f\n", valor)
	}
}