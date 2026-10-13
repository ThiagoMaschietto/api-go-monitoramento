package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Estrutura de dados para a resposta JSON da nossa API
type HealthCheck struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthCheck{
		Status:    "UP",
		Timestamp: time.Now(),
		Message:   "API em Go rodando com sucesso para o portfólio de DevOps!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)

	port := ":8080"
	fmt.Printf("Servidor Go iniciado na porta %s...\n", port)
	
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}