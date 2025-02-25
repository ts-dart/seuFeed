package scraper

import (
	"fmt"
	"time"
)

type Post struct {
	postHrefLink string
	postImgSrc   string
	postCategory string
	postTitle    string
	font string
	fontImgSrc string
	section string
}

func Scraper() {
	postsList := make([]Post, 0, 100)
	postsList = append(postsList, Urlo()...)

	for {
		fmt.Println(postsList)
		time.Sleep(3 * time.Minute)
	} 
}