package scraper

import (
	//"fmt"
	"log"

	"github.com/gocolly/colly"
)

func urlth() ([]Post) {
	url, font, fontImg := envValidator("URLTH_ENV", "URLTH_ENV_FONT", "URLTH_ENV_FONT_IMG")

	var postsList []Post
	c := colly.NewCollector()

	c.OnHTML("div.col-xs-12 div.col-md-3 div.post", func(e *colly.HTMLElement) {
		post := Post{
			postHrefLink: e.ChildAttr("div.post a.post-img", "href"),
			postImgSrc: e.ChildAttr("div.post a.post-img img", "src"),
			postText: e.ChildText("div.post div.post-body h3.post-title a"),
			font: font,
			fontImgSrc: fontImg,
			section: "",
		}
		postsList = append(postsList, post)
	})

	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	return postsList
}