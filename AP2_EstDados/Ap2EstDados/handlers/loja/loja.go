package loja

import (
	"fmt"
	"net/http"
	"time"

	"api/dados"
	"api/processamento"
)

func Abrir(w http.ResponseWriter, r *http.Request) {
	intervalo := r.URL.Query().Get("intervalo")
	var duracao time.Duration

	if intervalo != "" {
		novoIntervalo, err := time.ParseDuration(intervalo)
		if err != nil {
			http.Error(w, "Intervalo inválido", http.StatusBadRequest)
			return
		}
		duracao = novoIntervalo
	} else {
		duracao = 30 * time.Second
	}

	err := dados.AbrirLoja(duracao)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Loja aberta"))

	currentTime := time.Now()
	fmt.Printf("%s Loja aberta\n", currentTime.Format("02/01/2006 15:04:05"))
}

func Fechar(w http.ResponseWriter, r *http.Request) {
	err := dados.FecharLoja()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Write([]byte("Loja fechada"))

	currentTime := time.Now()
	fmt.Printf("%s Loja fechada\n", currentTime.Format("02/01/2006 15:04:05"))
}
