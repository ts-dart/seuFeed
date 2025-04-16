package main

import (
	"log"
	"net/http"

	"github.com/ts-dart/seuFeed/api/internal/api"
)

func main() {
  api.Start()

  log.Println("API running")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatalf("Erro ao iniciar servidor: %v", err)
  }
}
