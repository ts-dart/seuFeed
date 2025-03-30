package scraper

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func envValidator(u string, f string, fi string) (string, string, string) {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar o .env", err))
	}

	url := os.Getenv(u)
	if url == "" {
		panic(fmt.Sprintf("URLO_ENV não está definido", err))
	}

	font := os.Getenv(f)
	if font == "" {
		panic(fmt.Sprintf("URLO_ENV_FONT nao esta definido", err))
	}

	fontImg := os.Getenv(fi)
	if fontImg == "" {
		panic(fmt.Sprintf("URLO_ENV_FONT nao esta definido", err))
	}

	return url, font, fontImg
}
