package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func urlo() []Post {
	url, font, fontImg := envValidator("URLO_ENV", "URLO_ENV_FONT", "URLO_ENV_FONT_IMG")

	var postsList []Post
	c := colly.NewCollector()

	// Captura os elementos necessários
	c.OnHTML("#blockList_1", func(e *colly.HTMLElement) {
		post := Post{
			PostHrefLink: e.ChildAttr("figure a", "href"),
			PostImgSrc:   e.ChildAttr("figure a img", "src"),
			PostText:     e.ChildText("figcaption a h3"),
			Font:         font,
			FontImgSrc:   fontImg,
			Section:      "main",
		}
		postsList = append(postsList, post)
	})

	// Faz a requisição
	err := c.Visit(url)
	if err != nil {
		panic(fmt.Sprintf("Erro ao visitar a URL: %v", err))
	}

	log.Println("URLO", postsList)
	return postsList
}
