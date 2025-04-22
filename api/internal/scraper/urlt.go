package scraper

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func urlt() []Post {
	url, font, fontImg := envValidator("URLT_ENV", "URLT_ENV_FONT", "URLT_ENV_FONT_IMG")

	var postsList []Post
	c := colly.NewCollector()

	c.OnHTML("div._evg div._evt div.bastian-feed-item", func(e *colly.HTMLElement) {
		post := Post{
			PostHrefLink: e.ChildAttr("div.feed-post-body div.feed-post-body-title h2 a", "href"),
			PostImgSrc:   e.ChildAttr("div.feed-post-body div.feed-media-wrapper picture img", "src"),
			PostText:     e.ChildText("div.feed-post-body div.feed-post-body-title h2 a p"),
			Font:         font,
			FontImgSrc:   fontImg,
			Section:      "main",
		}
		postsList = append(postsList, post)
	})

	err := c.Visit(url)
	if err != nil {
		panic(fmt.Sprintf("Erro ao visitar a URL: %v", err))
	}

	log.Println("URLT", postsList)
	return postsList
}
