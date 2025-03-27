package api

import (
	"encoding/json"
	"net/http"
)

func Start() {
	http.HandleFunc("/", d)
	http.HandleFunc("/getPostsBySection", getPostsBySectionHandler)
	http.HandleFunc("/getAllPosts", getAllPostsHandler)
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
