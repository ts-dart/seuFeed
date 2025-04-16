package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Post struct {
	PostHrefLink string `json:"post_href_link"`
	PostImgSrc string `json:"post_img_src"`
	PostText string `json:"post_text"`
	Font string `json:"font"`
	FontImgSrc string `json:"font_img_src"`
	Section string `json:"section"`
	//hora
}

var PostsList []Post = make([]Post, 0, 100)
func Scraper() {
	for {
		if len(PostsList) > 0 {
			PostsList = PostsList[:0]
		}

		handleScraperErrors(urlo)
		handleScraperErrors(urlt)

		writeJson(PostsList)
		fmt.Println(PostsList)

		time.Sleep(time.Duration(randomNumber()) * time.Minute)
	}
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))
func randomNumber() (int) {
	return rng.Intn(8) + 3
}

// Função para lidar com erros nos scrapers sem interromper o fluxo
func handleScraperErrors(scraperFunc func() []Post) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Erro no scraper: %v.", r) 
		}
	}()

	PostsList = append(PostsList, scraperFunc()...)
}

func writeJson(PostsList []Post) {
	// Converter a estrutura para JSON formatado
	jsonData, err := json.MarshalIndent(PostsList, "", "  ")
	if err != nil {
		log.Fatal("Erro ao converter para JSON:", err)
		return
	}

	// Criar um arquivo para salvar o JSON
	file, err := os.Create("data.json")
	if err != nil {
		log.Fatal("Erro ao criar arquivo:", err)
		return
	}
	defer file.Close()

	// Escrever os dados no arquivo
	_, err = file.Write(jsonData)
	if err != nil {
		log.Fatal("Erro ao escrever no arquivo:", err)
		return
	}
}