package scraper

import (
	//"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

func Urlo() ([]Post) {
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
	c.OnHTML("section.block--manchetes section.block_style_second-level div.container", func(e *colly.HTMLElement) {
		post := Post{
			postHrefLink: e.ChildAttr("div.post a.post-img", "href"),
			postImgSrc:   e.ChildAttr("div.post a.post-img img", "src"),
			postCategory: e.ChildText("div.post div.post-body div.post-meta a.post-category"),
			postTitle:    e.ChildText("div.post div.post-body h3.post-title a"),
			font:
			fontImgSrc: 
			section: 
		}
		postsList = append(postsList, post)
	})
	/*
	c.OnHTML("div.col-xs-12 div.col-md-3 div.post", func(e *colly.HTMLElement) {
		post := Post{
			postHrefLink: e.ChildAttr("div.post a.post-img", "href"),
			postImgSrc:   e.ChildAttr("div.post a.post-img img", "src"),
			postCategory: e.ChildText("div.post div.post-body div.post-meta a.post-category"),
			postTitle:    e.ChildText("div.post div.post-body h3.post-title a"),
		}
		postsList = append(postsList, post)
	})
	*/

	// Faz a requisição
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return postsList
}
