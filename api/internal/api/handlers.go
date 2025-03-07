package api

import (
	"encoding/json"
	"net/http"
)

func handlers() {
	s := postsBySection("")

	http.HandleFunc("/postsBySection", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Definindo o cabeçalho da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Criando a resposta
	response := Response{Message: "Olá, mundo!"}

	// Convertendo para JSON e enviando a resposta
	json.NewEncoder(w).Encode(response)
}
