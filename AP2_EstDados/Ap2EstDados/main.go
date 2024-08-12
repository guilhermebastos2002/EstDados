package main

import (
	"fmt"
	"net/http"
	"time"

	"api/dados"
	"api/handlers/metricas"
	"api/handlers/pedidos"
	"api/handlers/produtos"
	"api/loja"
	"api/processamento"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// rotas
	r.HandleFunc("/produto", produtos.Adicionar).Methods("POST")
	r.HandleFunc("/produto", produtos.Buscar).Methods("GET")
	r.HandleFunc("/produto", produtos.Remover).Methods("DELETE")

	r.HandleFunc("/produtos", produtos.Listar).Methods("GET")

	r.HandleFunc("/pedido", pedidos.Adicionar).Methods("POST")
	r.HandleFunc("/pedidos", pedidos.Listar).Methods("GET")

	r.HandleFunc("/metricas", metricas.Metricas).Methods("GET")

	r.HandleFunc("/abrir", loja.Abrir).Methods("POST")
	r.HandleFunc("/fechar", loja.Fechar).Methods("POST")

	go processamento.ProcessaPedidos()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	currentTime := time.Now()
	fmt.Printf("%s Servidor iniciado em http://localhost:8080\n", currentTime.Format("02/01/2006 15:04:05"))

	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("%s Erro ao iniciar o servidor: %v\n", currentTime.Format("02/01/2006 15:04:05"), err)
	}
}
