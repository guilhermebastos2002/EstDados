package main 

import (
	"fmt"
	"math"
)

type Ponto struct {
	X, Y, float64
}

func (p Ponto) DistanciaOrigem() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
	ponto := Ponto{3, 4}
	fmt.Println("Distancia até a origem: ", ponto.DistanciaOrigem)
}
