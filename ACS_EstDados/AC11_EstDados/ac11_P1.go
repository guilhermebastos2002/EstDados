package main

import (
    "fmt"
)

type Car struct {
    name     string
    velocity float64
}

func main() {
    carX := Car{name: "Carro X", velocity: 60} // Carro X com velocidade 60 km/h
    carY := Car{name: "Carro Y", velocity: 90} // Carro Y com velocidade 90 km/h

    var distancia float64
    fmt.Scanln(&distancia)

    tempo := tempoParaAlcancarDistancia(distancia, carX.velocity, carY.velocity)
    fmt.Printf("%.0f minutos", tempo)
}


func tempoParaAlcancarDistancia(distancia, velocidadeX, velocidadeY float64) float64 {
    tempo := distancia / (velocidadeY - velocidadeX) * 60 // tempo em minutos
    return tempo
}