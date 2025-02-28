package scraper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func envValidator(u string, f string, fi string) (string, string, string) {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o .env")
	}

	url := os.Getenv(u)
	if url == "" {
		log.Fatal("URLO_ENV não está definido")
	}

	font := os.Getenv(f)
	if font == "" {
		log.Fatal("URLO_ENV_FONT nao esta definido")
	}

	fontImg := os.Getenv(fi)
	if fontImg == "" {
		log.Fatal("URLO_ENV_FONT nao esta definido")
	}

	return url, font, fontImg
}
