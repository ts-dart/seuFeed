package main
 
import (
  "net/http"
  "github.com/ts-dart/seuFeed/api/internal/api"
)

func main() {
  api.Handlers()

  http.ListenAndServe(":8080", nil)
}
