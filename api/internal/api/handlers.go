package api

import (
	"encoding/json"
	"net/http"
)

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Definindo os headers de CORS
		w.Header().Set("Access-Control-Allow-Origin", "*") // ou especifique o domínio
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Para requests OPTIONS (pré-flight), respondemos com status 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Start() {
	http.HandleFunc("/", d)
	http.Handle("/PostsBySection", enableCORS(http.HandlerFunc(getPostsBySectionHandler)))
	http.Handle("/AllPosts", enableCORS(http.HandlerFunc(getAllPostsHandler)))
	http.Handle("/ClimateData", enableCORS(http.HandlerFunc(getClimateDateHandler)))
}

func d(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Api responding"))
}

func getPostsBySectionHandler(w http.ResponseWriter, r *http.Request) {
	// Definindo o cabeçalho da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Criando a resposta
	ft := r.URL.Query().Get("ft")
	response := getPostsBySection(ft)
	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
		return
	}

	// Convertendo para JSON e enviando a resposta
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getAllPostsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := getAllPosts()

	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func getClimateDateHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	lat := r.URL.Query().Get("latitude")
	lon := r.URL.Query().Get("longitude")
	response := getClimateDate(lat, lon)

	jsonData, err := json.Marshal(response)

	if err != nil {
		http.Error(w, "Erro ao gerar JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}