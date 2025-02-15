package scraper

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type News struct {
	NewsHrefLink string
	NewsImgSrc   string
	NewsCategory string
	NewsTitle    string
}

func main() {
	// Carrega variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o .env")
	}

	url := os.Getenv("URLO_ENV")
	if url == "" {
		log.Fatal("URLO_ENV não está definido")
	}

	var newsList []News

	c := colly.NewCollector()

	// Captura os elementos necessários
	c.OnHTML("div.col-xs-12 div.col-md-3 div.post", func(e *colly.HTMLElement) {
		news := News{
			NewsHrefLink: e.ChildAttr("div.post a.post-img", "href"),
			NewsImgSrc:   e.ChildAttr("div.post a.post-img img", "src"),
			NewsCategory: e.ChildText("div.post div.post-body div.post-meta a.post-category"),
			NewsTitle:    e.ChildText("div.post div.post-body h3.post-title a"),
		}

		newsList = append(newsList, news)
	})

	// Faz a requisição
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	// Exibe os resultados
	for _, news := range newsList {
		fmt.Printf("Título: %s\nCategoria: %s\nImagem: %s\nLink: %s\n\n",
			news.NewsTitle, news.NewsCategory, news.NewsImgSrc, news.NewsHrefLink)
	}
}
