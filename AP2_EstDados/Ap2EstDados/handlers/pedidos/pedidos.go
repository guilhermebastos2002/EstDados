package pedidos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api/dados"
)

func Adicionar(w http.ResponseWriter, r *http.Request) {
	var novoPedido dados.Pedido

	err := json.NewDecoder(r.Body).Decode(&novoPedido)
	if err != nil {
		http.Error(w, "Erro ao decodificar dados", http.StatusBadRequest)
		return
	}

	err = dados.FilaPedidos.Adicionar(novoPedido.Delivery, novoPedido.IdProdutos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Sucesso"))

	currentTime := time.Now()
	fmt.Printf("%s Pedido adicionado com sucesso\n", currentTime.Format("02/01/2006 15:04:05"))
}

func Listar(w http.ResponseWriter, r *http.Request) {
	filaJSON, err := json.Marshal(dados.FilaPedidos.Listar())
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(filaJSON)

	currentTime := time.Now()
	fmt.Printf("%s Fila de pedidos obtida\n", currentTime.Format("02/01/2006 15:04:05"))
}
