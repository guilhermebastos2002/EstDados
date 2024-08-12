package main

import (
	"fmt"
)

type Geometria struct {
	B, H, L1, L2 float64
}

func (g Geometria) CalculaArea() float64 {
	return g.B * g.H
}

func (g Geometria) CalculaPerimetro() float64 {
	return 2*g.L1 + 2*g.L2
}

func () {
	var op int
	var g Geometria

	for {
		fmt.Println("Qual operação deseja fazer?")
		fmt.Println("Digite 1 para calcular a área, 2 para calcular o perímetro e 3 para sair")
		fmt.Scan(&op)

		switch op {
		case 1:
			fmt.Println("Digite a base do retângulo: ")
			fmt.Scan(&g.B)
			fmt.Println("Digite a altura do retângulo: ")
			fmt.Scan(&g.H)
			fmt.Println("A área desse retângulo é:", g.CalculaArea())

		case 2:
			fmt.Println("Digite o lado maior do retângulo: ")
			fmt.Scan(&g.L1)
			fmt.Println("Digite o lado menor do retângulo: ")
			fmt.Scan(&g.L2)
			fmt.Println("O perímetro desse retângulo é:", g.CalculaPerimetro())

		case 3:
			fmt.Println("Saindo do programa...")
			return

		default:
			fmt.Println("Digite um valor válido.")
		}
	}
}
