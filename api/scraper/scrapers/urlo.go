package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func Urlo(postsList []Post) {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o .env")
	}

	url := os.Getenv("URLO_ENV")
	if url == "" {
		log.Fatal("URLO_ENV não está definido")
	}

	var postsList []Post

	c := colly.NewCollector()

	// Captura os elementos necessários
	c.OnHTML("div.col-xs-12 div.col-md-3 div.post", func(e *colly.HTMLElement) {
		post := Post{
			PostHrefLink: e.ChildAttr("div.post a.post-img", "href"),
			PostImgSrc:   e.ChildAttr("div.post a.post-img img", "src"),
			PostCategory: e.ChildText("div.post div.post-body div.post-meta a.post-category"),
			PostTitle:    e.ChildText("div.post div.post-body h3.post-title a"),
		}

		postsList = append(postsList, post)
	})

	// Faz a requisição
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	// Exibe os resultados
	for i := range postsList {
		//fmt.Printf("Título: %s\nCategoria: %s\nImagem: %s\nLink: %s\n\n",
		//	post.PostTitle, post.PostCategory, post.PostImgSrc, post.PostHrefLink)
	}
}
