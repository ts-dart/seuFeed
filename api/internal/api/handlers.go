package api

import (
	"encoding/json"
	"net/http"
)

func Handlers() {
	http.HandleFunc("/postsBySection", postsBySectionHandler)
}

func postsBySectionHandler(w http.ResponseWriter, r *http.Request) {
	// Definindo o cabeçalho da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Criando a resposta
	ft := r.URL.Query().Get("ft")
	response := postsBySection(ft)

	// Convertendo para JSON e enviando a resposta
	json.NewEncoder(w).Encode(response)
}
