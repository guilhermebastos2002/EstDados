package metricas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"api/dados"
)

func Metricas(w http.ResponseWriter, r *http.Request) {
	metricasJSON, err := json.Marshal(dados.MetricasColetadas)
	if err != nil {
		http.Error(w, "Erro ao converter mensagens para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(metricasJSON)

	currentTime := time.Now()
	fmt.Printf("%s Métricas coletadas com sucesso\n", currentTime.Format("02/01/2006 15:04:05"))
}