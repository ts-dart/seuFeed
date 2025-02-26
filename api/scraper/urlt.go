package scraper

import (
	//"fmt"
	"log"

	"github.com/gocolly/colly"
)

func urlt() ([]Post) {
	url, font, fontImg := envValidator("URLT_ENV", "URLT_ENV_FONT", "URLT_ENV_FONT_IMG")

	var postsList []Post
	c := colly.NewCollector()

	c.OnHTML("div._evt div.bastian-page div._evg div._evt", func(e *colly.HTMLElement) {
		post := Post{
			postHrefLink: e.ChildAttr("div div", "data-mrf-link"),
			postImgSrc: e.ChildAttr("div.feed-media-wrapper picture img", "src"),
			postText: e.ChildText("div.feed-post-body div.feed-post-body-title h2 a p"),
			font: font,
			fontImgSrc: fontImg,
			section: "main",
		}
		postsList = append(postsList, post)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return postsList
}
