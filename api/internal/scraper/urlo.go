package scraper

import (
	//"fmt"
	"log"
	"github.com/gocolly/colly"
)

func urlo() ([]Post) {
	url, font, fontImg := envValidator("URLO_ENV", "URLO_ENV_FONT", "URLO_ENV_FONT_IMG")

	var postsList []Post
	c := colly.NewCollector()

	// Captura os elementos necessários
	c.OnHTML("#block11073171 .has--thumb", func(e *colly.HTMLElement) {
		post := Post{
			PostHrefLink: e.ChildAttr("figure a", "href"),
			PostImgSrc: e.ChildAttr("figure a img", "src"),
			PostText: e.ChildText("figcaption a h3"),
			Font: font,
			FontImgSrc: fontImg,
			Section: "main",
		}
		postsList = append(postsList, post)
	})
	 
	// Faz a requisição
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return postsList
}
